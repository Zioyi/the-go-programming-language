package main

import (
	"fmt"
	"sync"
	"time"
)

func syncWaitCroupExample() {
	var wg sync.WaitGroup
	for _, language := range []string{"python", "go", "java"} {
		wg.Add(1)
		go func(lang string) {
			defer wg.Done()
			fmt.Println(lang)
		}(language)
	}

	wg.Wait()
}

func syncCondExample() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from Queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}
func syncOnceExample() {
	/*
		Do方法只会执行一次
		适用场景：某段代码在多个协程间只希望被执行一次，如：初始化
	*/
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	fmt.Printf("Count: %d\n", count)
}

func main() {
	syncOnceExample()
}
