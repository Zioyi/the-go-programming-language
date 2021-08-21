package main

import (
	"fmt"
	"sort"
)

func FindX(sortedA []int, x int) int {
	i := sort.Search(len(sortedA), func(i int) bool { return sortedA[i] >= x })
	if i < len(sortedA) && sortedA[i] == x {
		return i
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	// i := sort.Search(len(a), func(i int) bool {
	// 	defer func() {
	// 		a[0] = 301
	// 		a[1] = 302
	// 		a[2] = 303
	// 		a[3] = 304
	// 		a[4] = 305
	// 	}()
	// 	println(i, a[i], a[i] >= 305)

	// 	return a[i] >= 305
	// })
	// fmt.Println(i, a)

	fmt.Println(FindX(a, 4))
	fmt.Println(FindX(a, 100))
}
