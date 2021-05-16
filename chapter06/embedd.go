package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColoredPoint struct {
	*Point
	Color color.RGBA
}

type A struct{}

func (a A) Func() {}

type B struct{}

func (b B) Func() {}

type C struct {
	A
	B
}

func main() {
	var cp ColoredPoint
	cp.Point = &Point{}
	cp.Point.X = 1
	fmt.Println(cp.Point.X) // 1

	p := ColoredPoint{&Point{1, 1}, color.RGBA{255, 0, 0, 255}}
	q := ColoredPoint{&Point{5, 4}, color.RGBA{0, 0, 255, 255}}

	//fmt.Println(p.Distance(q)) // 编译错误：cannot use q (type ColoredPoint) as type Point in argument to p.Point.Distance
	fmt.Println(p.Distance(*q.Point)) // 5
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(*q.Point)) // 10

	c := C{}
	c.Func()
}
