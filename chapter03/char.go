package main

import "fmt"

type Trie struct {
	children [26]*Trie
	count    int
}

func main() {
	a := 'a'
	fmt.Println(a - 'b')
	t := Trie{}
	fmt.Println(t.children)
}
