package main

import "fmt"

type IntList struct {
	Value int
	Next  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Next.Sum()
}

func main() {
	a1 := IntList{1, nil}
	a2 := IntList{2, &a1}
	a3 := IntList{3, &a2}

	fmt.Println(a3.Sum()) // 6

}
