package main

import "fmt"

func modifyMap(m map[int]int) {
	m[1] = 1
}

func main() {
	m1 := map[int]int{}
	m1[1] = 2
	m1[2] = 2
	m1[3] = 3
	fmt.Printf("before: %v\n", m1)
	modifyMap(m1)
	fmt.Printf("after: %v\n", m1)
}
