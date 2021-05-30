package main

import (
	"encoding/json"
	"expvar"
	"fmt"
	"net"
	"net/http"
)

var (
	counts = expvar.NewMap("counters")
)

func init() {
	counts.Add("a", 10)
	counts.Add("b", 10)
}

func main() {
	a := make(map[string]string)
	a["name"] = "zyf"
	b, _ := json.Marshal(a)
	fmt.Println(string(b))

	sock, err := net.Listen("tcp", "localhost:9123")
	if err != nil {
		panic("sock error")
	}
	go func() {
		fmt.Println("HTTP now available at port 9123")
		http.Serve(sock, nil)
	}()
	fmt.Println("htllo")
	select {}
}
