package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	appendI(a, 4)
	fmt.Println(a)
}

func appendI(s []int, i int) {
	s[0] = 0
	s = append(s, i)
	fmt.Println(s)
}
