package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func main() {
	c1 := make(chan int, 0)
	c2 := make(chan int, 0)
	c3 := c2
	println(c1 == c2)
	println(c2 == c3)

	var cr <-chan int
	fmt.Printf("%T\n", cr)

	a := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c}
	c := bytes2str(a)
	fmt.Printf("%s\n", c)

	d := str2runes(a)
	fmt.Printf("%v\n", d)

	str1 := "Hello, 世界"
	str1Bytes := str2bytes(str1)
	fmt.Printf("%x\n", str1Bytes)

	strd := runes2string(d)
	fmt.Printf("%s\n", strd)
}

// []byte(s)转换模拟
func str2bytes(s string) []byte {
	p := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		p[i] = c
	}
	return p
}

// string(bytes)转换模拟
func bytes2str(s []byte) (p string) {
	// 必须要拷贝一份，不然如果s对内容进行了修改，会影响到p
	data := make([]byte, len(s))
	for i, c := range s {
		data[i] = c
	}
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&p))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(s)
	return p
}

// []rune(s)转换模拟
func str2runes(s []byte) []rune {
	var p []int32
	for len(s) > 0 {
		r, size := utf8.DecodeRune(s)
		p = append(p, int32(r))
		s = s[size:]
	}
	return []rune(p)
}

// rune == 3个byte
// string(runes)转换模拟
func runes2string(s []int32) string {
	var p []byte
	buf := make([]byte, 3)
	for _, r := range s {
		n := utf8.EncodeRune(buf, r)
		p = append(p, buf[:n]...)
	}
	return string(p)
}
