package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

const text = `
package io

// Wirter is a interface for write
type Writer interface {
	Write (p byte[]) (n int, err error)
}
`

func main() {
	// sc := new(scanner.Scanner)
	// sc.Init(strings.NewReader(text))
	// sc.Mode = scanner.ScanIdents | scanner.ScanChars | scanner.ScanFloats | scanner.SkipComments
	// var tok rune
	// for tok != scanner.EOF {
	// 	tok = sc.Scan()
	// 	fmt.Println("At position", sc.Pos(), ":", sc.TokenText())
	// }

	sc := new(scanner.Scanner)
	sc.Init(strings.NewReader("5 * sqrt2 + 3)"))
	sc.Mode = scanner.ScanIdents | scanner.ScanChars | scanner.ScanFloats | scanner.SkipComments
	var tok rune
	for tok != scanner.EOF {
		tok = sc.Scan()
		fmt.Println("At position", sc.Pos(), ":", sc.TokenText(), tok == scanner.Int)
	}
}
