package main

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) Print() {
	fmt.Printf("%#v", p)
}

type Circle struct {
	Point
	Radius int
}

func (c Circle) Print() {
	fmt.Printf("%#v", c)
}

type Band struct {
	Name string
}

func (b Band) Print() {
	fmt.Printf("%#v", b)
}

type Wheel struct {
	Circle
	Band
	Spokes int
}

func (w Wheel) Print() {
	fmt.Printf("%#v", w)
}

func main() {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	fmt.Printf("%#v\n", w) // main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}
	w.Print()              // ambiguous selector w.Print
	w.Circle.Print()
	w.Band.Print()
}
