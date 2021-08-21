package main

type foo struct {
}

//func getFooPointer() *foo {
//	var result foo
//	return &result
//}

func getFooValue() foo {
	var result foo
	return result
}

func f(a *foo) {}

func main() {
	var p foo
	f(&p)
}
