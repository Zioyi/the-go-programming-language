package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var w io.Writer
	b := new(bytes.Buffer)
	fmt.Println(b.String())
	w = b
	w.Write([]byte("hello"))
	fmt.Println(b.String())
}
