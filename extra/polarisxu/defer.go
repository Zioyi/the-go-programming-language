package main

import "fmt"

func hello(i *int) int {
	defer func() {
		*i = 19
	}()
	return *i
}

func hello2(i *int) (j int) {
	defer func() {
		*i = 19
		j = *i
	}()
	return *i
}

func main() {
	i := 10
	j := hello(&i)
	fmt.Printf("i: %d, j: %d\n", i, j)

	i = 10
	j = hello2(&i)
	fmt.Printf("i: %d, j: %d\n", i, j)
}

/*
返回值有命名和无命名的区别是
有命名类型的值会被defer函数修改
*/
