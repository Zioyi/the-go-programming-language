package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()

	fmt.Fprintf(conn, "CET / HTTP/1.0\r\n\r\n")

	var sb strings.Builder
	buf := make([]byte, 256)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
				return
			}
			break
		}

		sb.Write(buf[:n])
	}

	fmt.Println("response:", sb.String())
	fmt.Println("total response size:", sb.Len())
}
