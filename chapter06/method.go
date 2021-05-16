package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 指针方式的接收器
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type Line struct {
	Start Point
	End   Point
	// Length float64
}

func (L Line) Length() float64 {
	return L.Start.Distance(L.End)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	p := Point{1, 2}

	q := Point{4, 6}
	fmt.Println(p.Distance(q))
	line := Line{p, q}
	fmt.Println(line.Length())

	// 计算三角形周长
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())

	p.ScaleBy(200)
	fmt.Println(p)

	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	p = Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p) // {2, 4}

	p2 := Point{1, 2}
	(&p2).ScaleBy(2)
	fmt.Println(p2) // {2, 4}

	// 方法变量
	a := Point{1, 2}
	b := Point{4, 6}
	distanceFromA := a.Distance
	fmt.Println(distanceFromA(b)) // 5
	origin := Point{0, 0}
	fmt.Println(distanceFromA(origin)) // 2.23606797749979

	scaleA := a.ScaleBy
	scaleA(2)
	fmt.Println(a)

	distance := Point.Distance // 方法表达式 由Point类型的方法赋予
	fmt.Println(distance(a, b))
	fmt.Printf("%T\n", distance)

	// scale := Point.ScaleBy // 编译报错：nvalid method expression Point.ScaleBy (needs pointer receiver: (*Point).ScaleBy
	scale := (*Point).ScaleBy
	scale(&a, 2)
	fmt.Println(a)
	fmt.Printf("%T\n", scale)
}
