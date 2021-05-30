package main

import (
	"fmt"
	"time"
)

func forSelectExample() {
	chan1 := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		close(chan1)
	}()
	start := time.Now()
	for {
		fmt.Printf("-------\n")
		select {
		case x, ok := <-chan1:
			if !ok {
				return
			}
			fmt.Printf("got x: %d\n", x)
		case <-time.After(3 * time.Second):
			fmt.Printf("Timeout..., wait %v", time.Since(start))

		}
	}
}

func goroutineLeakExample() {
	// 会协程泄露不
	//doWork := func(strings <-chan string) <-chan interface{} {
	//	complete := make(chan interface{})
	//	go func() {
	//		defer fmt.Println("doWork exited.")
	//		defer close(complete)
	//		for s := range strings {
	//			fmt.Println(s)
	//		}
	//	}()
	//	return complete
	//}
	//
	//doWork(nil)
	//fmt.Println("Done")

	return
}

func main() {
	forSelectExample()
}
