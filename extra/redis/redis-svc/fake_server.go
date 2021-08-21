package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func parseCmd(cmdString string) {
	subs := strings.Split(cmdString, "\r\n")
	validSubs := make([]string, 0)
	for i := 2; i < len(subs); i += 2 {
		// validation
		validSubs = append(validSubs, subs[i])
	}

	cmd := validSubs[0]
	switch cmd {
	case "set":
		fmt.Printf("handle set, %v\n", validSubs)
	case "get":
		fmt.Printf("handle get, %v\n", validSubs)

	}

}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}

		recv := string(buf[:n])
		parseCmd(recv)

		fmt.Printf("收到的数据：\n-----start-----\n%v\n-----end-----\n", recv)

		_, err = conn.Write([]byte("+OK\r\n"))
		if err != nil {
			fmt.Printf("write from conn falied, err:%v\n", err)
			break
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Printf("listen failed, err: %v\n", err)
		return
	}
	fmt.Printf("Fake Redis Sever Running\n")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			return
		}
		process(conn)
	}
}
