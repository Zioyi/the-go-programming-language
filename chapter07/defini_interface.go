package main

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"golang.org/x/net/html"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
}

type ReadWriterCloser interface {
	Reader
	Writer
	Closer
}

type ReadWriter2 interface {
	Read(p []byte) (n int, err error)
	Writer
}

// homework

type HtmlReader struct {
	rawDoc string
}

func (h HtmlReader) Read(p []byte) (n int, err error) {
	b := []byte(h.rawDoc)
	n = copy(p, b)
	return n, io.EOF

}

type LimitedReader struct {
	r io.Reader
	N int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}

func (l LimitedReader) Read(p []byte) (n int, err error) {
	return
}

func main() {
	var w Writer
	fmt.Printf("%T %v %+v", w, w, w)
	fmt.Println(reflect.TypeOf(w))

	raw := "<html></html>"
	hr := HtmlReader{rawDoc: "<html></html>"}
	r := strings.NewReader(raw)
	doc11, _ := html.Parse(r)
	fmt.Printf("%v\n", doc11)
	doc, _ := html.Parse(hr)
	fmt.Printf("aaa: %v\n", doc)
}
