/*
如果类型S内嵌了T，如果T定义函数receiver都是*T类型，那么S定义的函数也应该以*S为receiver为佳
*/
package main

import (
	"fmt"
	"sync"
)

type ConcurrentLocker struct {
	sync.Map
}

type LeaveFunc func()

func (cl *ConcurrentLocker) Enter(key string) (bool, LeaveFunc) {
	if _, occupied := cl.LoadOrStore(key, struct{}{}); !occupied {
		return true, func() {
			cl.Delete(key)
		}
	}
	return false, nil
}

func main3() {
	cl := ConcurrentLocker{}
	key := "some-key"
	fmt.Println(cl.Enter(key))
	fmt.Println(cl.Enter(key))
}
