package main

import (
	"fmt"
	"time"
)

func selectExample() {
	rand01 := func() chan int {
		ch := make(chan int)
		go func() {
			for {
				select {
				case ch <- 0:
				case ch <- 1:

				}
			}
		}()
		return ch
	}

	generator := rand01()
	for i := 0; i < 10; i++ {
		fmt.Println(<-generator)
	}
}

func timeAfterExample() {
	c := make(chan int)
	start := time.Now()
	go func() {
		x := <-c
		fmt.Printf("\nget val:%d from channel c\n", x)
	}()
	select {
	case c <- 2:
		fmt.Printf("\nwrite val to channel c\n")
	case <-time.After(3 * time.Second):
		fmt.Printf("Timeout ......, wait %v\n", time.Since(start))
	}
}

func main() {
	timeAfterExample()
}
