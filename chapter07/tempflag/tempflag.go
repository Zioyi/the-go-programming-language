package main

import (
	"flag"
	"fmt"

	"github.com/Zioyi/the-go-programming-language/chapter07/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
