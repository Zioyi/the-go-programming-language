package main

import (
	"fmt"
	"time"
)

func main() {
	var f = func() {
		e := recover()
		if e != nil {
			fmt.Printf("panic catch: %v\n", e)
		}
	}

	go func() {
		defer func() {
			e := recover()
			fmt.Printf("panic catch %v\n", e)
			fmt.Println("cat")
		}()
		panic("aaa")
	}()
	go func() {
		defer f()
		panic("bbb")
	}()
	time.Sleep(time.Second * 4)
}
