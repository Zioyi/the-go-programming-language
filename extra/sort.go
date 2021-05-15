package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 1, 2, 3444, 23, 0}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	fmt.Println(a)

	var prev int
	prev = nil
	fmt.Println(prev)
}
