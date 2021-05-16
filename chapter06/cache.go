package main

import (
	"fmt"
	"strconv"
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func Set(key, vaule string) {
	cache.Lock()
	cache.mapping[key] = vaule
	cache.Unlock()
}

func testFunc(wg *sync.WaitGroup, key string, value string) {
	defer wg.Done()
	Set(key, value)
	// time.Sleep(time.Second * 3)  取消注释，会发生数据竞争
	fmt.Printf("expect :%s, actual: %s\n", value, Lookup(key))

}

func main() {
	wg := sync.WaitGroup{}
	nums := 5

	for i := 0; i < nums; i++ {
		wg.Add(1)
		go testFunc(&wg, "sbat", strconv.Itoa(i))
	}
	wg.Wait()

}
