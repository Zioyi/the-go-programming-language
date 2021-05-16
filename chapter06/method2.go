package main

type Point struct{}

func (p *Point) PtrFunc() {}
func (p Point) Func()     {}

func main() {
	p := Point{}
	ptr := &Point{}
	ptr.PtrFunc()
	ptr.Func()

	Point{}.Func()
	// Point{}.PtrFunc() // 编译错误：cannot call pointer method on Point literal

	p.Func()
	p.PtrFunc() // 编译器做了隐式转换
}
