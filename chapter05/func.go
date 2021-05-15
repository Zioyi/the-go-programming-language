package main

import (
	"fmt"
	"time"
)

var f func(x, y float64) float64
var g func(x float64, y float64) (a float64)

func sum(x, y int) (z int) {
	z = x + y
	return
}

func foo() (x, y int) {
	x, y = 1, 1
	return
}

func foo2() (int, int) {
	x, y := 1, 1
	return x, y
}

func add(x, y int) int         { return x + y }
func sub(a int, b int) (c int) { c = a - b; return }

var f1 func(x float64) float64

func main2() {
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)

	f = func(x, y float64) float64 {
		return x + y
	}
	fmt.Println(sum(1, 3))
	fmt.Println(foo2())

	fmt.Printf("%T\n", add) //
	fmt.Printf("%T\n", sub)
}

func incr(x int) {
	x++
}

func updateSlice(s []int, index int, val int) {
	s[index] = val
}

func updateMap(m map[int]int, k int, v int) {
	m[k] = v
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
func add1(r rune) rune { return r + 1 }

func Sprintf(format string, params ...interface{}) string {
	i, j := 0, 0
	s := ""
	for i < len(format)-1 {
		if format[i] == '%' && format[i+1] == 'd' {
			s = s + fmt.Sprintf("%v", params[j])
			j++
			i++
		} else {
			s = s + string(format[i])
		}
		i++
	}
	return s
}

func double(x int) int {
	return x + x
}

func clock(msg string) func() {
	start := time.Now()
	fmt.Printf("enter %s\n", msg)
	return func() { fmt.Printf("exit %s (%s)\n", msg, time.Since(start)) }
}

func SlowFunc() {
	defer clock("SlowFunc")()
	time.Sleep(3 * time.Second)
}

func myPanic() (a int) {
	defer func() {
		if p := recover(); p != nil {
			a = p.(int)
		}
	}()
	panic(222222)
}

func main() {
	var bill float64
	bill = 117.5*4 + 288
	bill /= 2
	fmt.Println(bill)
}
