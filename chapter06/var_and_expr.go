package main

import (
	"fmt"
	"time"
)

type Rocket struct {
}

func (r *Rocket) Launch() {
	fmt.Println("Launch...") // 不会输出
}

func main() {
	r := new(Rocket)
	time.AfterFunc(time.Second*1, r.Launch)
}
