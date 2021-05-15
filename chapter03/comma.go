package main

import (
	"fmt"
	"strings"
)

func comma(s string) string {
	ss := strings.Split(s, ".")
	var decimal string
	if len(ss) == 2 {
		s, decimal = ss[0], ss[1]
	}

	n := len(s)
	if n <= 3 {
		return s
	}

	s = comma(s[:n-3]) + "," + s[n-3:]
	if decimal != "" {
		s = s + "." + decimal
	}
	return s
}

// 非递归写法
func commaIter(s string) string {
	ss := strings.Split(s, ".")
	var decimal string
	if len(ss) == 2 {
		s, decimal = ss[0], ss[1]
	}
	n := len(s)
	offset := n % 3
	newBuf := []rune{}
	for i, val := range s {
		if i != 0 && i%3 == offset {
			newBuf = append(newBuf, ',')
		}
		newBuf = append(newBuf, val)
	}
	// 支持浮点数处理
	if decimal != "" {
		newBuf = append(newBuf, '.')
		newBuf = append(newBuf, []rune(decimal)...)

	}
	return string(newBuf)
}

func main() {
	number := "123456789011111"
	fmt.Println(comma(number))
	fmt.Println(commaIter(number))

	// 正确处理浮点数
	n := "123213415.12354"
	fmt.Println(comma(n))
	fmt.Println(commaIter(n))
}
