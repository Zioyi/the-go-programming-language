package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if cap(x) >= zlen {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z

}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main1() {
	a := []int{1, 2, 3}
	fmt.Println(a, a[1])
	b := a[:]
	b[1] = 100
	fmt.Println(a, b)
	a = append(a, 101)
	fmt.Println(a)

	a = remove(a, 0)
	fmt.Println(a)

	a = appendInt(a, 1111111)
	fmt.Println(a)

}

// 练习4.3: 重写reverse，使用数组指针作为参数而不是slice。
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse_array(s *[3]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main2() {
	a := []int{1, 2, 3}
	reverse(a)
	fmt.Println(a) // [3 2 1]

	b := [3]int{1, 333, 3}
	reverse_array(&b)
	fmt.Println(b)
}

// 练习4.4: 重写一个函数rotate，实现一次遍历就可以完成yuan

func main() {
	e := []int{1: 42, 55, 66, 77, 7: 88}
	fmt.Println(e)

	var a []int
	println(len(a), a == nil)
}
