package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a rune = '你'
	var b uint32 = 20320
	fmt.Printf("%c %d, 占用字节：%d\n", a, a, unsafe.Sizeof(a))
	fmt.Printf("%c %d\n", b, b)

	var china []rune = []rune{'中', '国'}
	fmt.Printf("%c %d", china, unsafe.Sizeof(china))
}
