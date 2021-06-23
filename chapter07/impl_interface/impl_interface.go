package main

import (
	"bytes"
	"fmt"
	"io"
)

type myWriter struct {
	scope string
}

func (m *myWriter) Write(p []byte) (n int, err error) {
	return
}

func (m *myWriter) Func() {
	return
}

func main() {
	emptyInterfaceIsAny()
}

func interfaceIsPromise() {
	var w io.Writer
	var rw io.ReadWriter

	my := &myWriter{"cn"}
	fmt.Println(my.scope)
	my.Func()
	w = my
	//fmt.Println(w.scope)
	//w.Func()
	w = rw
	//p := make([]byte, 2)
	//w.Write(p)
	//w = myWriter{}
	//rw = myWriter{}  // compile error:
	fmt.Printf("%T", w)
}

func emptyInterfaceIsAny() {
	var any interface{}
	any = true
	fmt.Println(any)
	any = 12.34
	fmt.Println(any)
	any = "hello"
	fmt.Println(any)
	any = map[string]int{"one": 1}
	fmt.Println(any)
	any = new(bytes.Buffer)
	fmt.Println(any)
	fmt.Println(1, "hello", []string{"world"})
}

func printlnArgsIsEmptyInterface() {
	fmt.Println(1, "hello", []string{"world"})
}
