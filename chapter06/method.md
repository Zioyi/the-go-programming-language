# 方法
在面向对象编程的编程思想里，类、对象、方法是基础。类比到Golang中
```go
// 类
type Point struct {X, Y int}
// 对象
p := Point{1, 2}
// 方法 即绑定在struct上的函数
// ...
```
## 方法声明
方法和函数类似，区别在于它在函数名前多了一个参数（接收器），用来将方法绑定在参数对应的类型上
```go
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

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(p.Distance(q))  // 5
}
```
**每个类型都有自己的命令空间，在同一个命名空间里不能有相同名称的方法和成员**
```go
type Line struct {
	Start  Point
	End    Point
	// Length float64
    // 如果取消上面这行的注释 编译报错：type Line has both field and method named Length
}

func (L Line) Length() float64 {
	return L.Start.Distance(L.End)
}

func main() {
    p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(p.Distance(q))  // 5
	line := Line{p, q}
	fmt.Println(line.Length())  // 5
}
```
**不同类型的命名空间是独立的，可以在不同类型中使用相同名字的方法**
```go
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
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())  // 12
}
```
## 指针接收者的方法
函数调用实参变量是以复制一份的方式传递的，如果我们想在函数中进行更改会很麻烦；如果一个实参太大，我们希望避免复制整个实参，我们可以通过指针的方式传递变量地址。这也同样使用与方法
```go
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
    p := Point{1, 2}
    p.ScaleBy(200)
    fmt.Printf("%+v", p) // {X:200 Y:400}
}
```
习惯上，如果`Point`上任何一个方法绑定指针接收者，那么所有的Point方法都应该使用指针接收者。方法的接收者只能是类型（*Point）或者类型指针（*Point）。

为了防止混淆，不允许本身是指针的类型进行方法声明：
```go
type p *int
func (p) f() {/*...*/}  // 编译错误：非法的接收者类型
```
以下几种写法都是合法的：
```go
// case1
r := &Point{1, 2}
r.ScaleBy(2)
fmt.Println(*r)  // {2, 4}

// case2
p1 := Point{1, 2}
pptr := &p1
pptr.ScaleBy(2)
fmt.Println(p1)  // {2, 4}

// case3
p2 := Point{1, 2}
(&p2).ScaleBy(2)
fmt.Println(p2)  // {2, 4}
```
注意，不能对一个不能取地址的Point接收者参数调用*Point方法，因为无法获得临时变量的地址。
```go
Piont{1,2}.ScaleBy(2)  // 编译错误
```
反过来，指针类型(*Point)变量，它是可以调用Point类型的方法
```go 
type Point struct{}

func (p *Point) PtrFunc() {}
func (p Point) Func()     {}

func main() {
	p := Point{}
	ptr := &Point{}
	ptr.PtrFunc()
	ptr.Func()

	Point{}.Func()
	Point{}.PtrFunc() // 编译错误：cannot call pointer method on Point literal

	p.Func()
	p.PtrFunc() // 编译器做了隐式转换
}
```
**疑惑：如果所有类型T方法的接收者是T自己（而非\*T），那么复制它的实例是安全的；调用方法的时候必须进行一次复制。但是任何方法的接收者是指针的情况下，应该避免复制T的实例，因为这么做可能会破坏原本的数据。**

## nil是一个合法的接收
方法的接收者可以是nil
```go
// *IntList的类型nil代表空列表
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

	fmt.Println(a3.Sum())  // 6

}
```
***当定义一个类型允许为nil作为接收者，应该在文档中显式地表明***

## 通过结构体内嵌组成类型
在一个结构体A中嵌套另一个结构体B，则结构体A可以调用结构体B的方法
```go
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
	Point
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)  // 1

	p := ColoredPoint{Point{1, 1}, color.RGBA{255, 0, 0, 255}}
	q := ColoredPoint{Point{5, 4}, color.RGBA{0, 0, 255, 255}}

	//fmt.Println(p.Distance(q)) // 编译错误：cannot use q (type ColoredPoint) as type Point in argument to p.Point.Distance
	fmt.Println(p.Distance(q.Point))  // 5
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))  // 10
}
```
`ColoredPoint`类型内嵌了`Point`类型，它可以调用`Point`的`Distance`和`ScaleBy`方法。也可以直接访问`Point`的成员变量。

实际上，内嵌字段会告诉编译生成额外的包装方法来调用 `Point`声明的方法：
```go
func (p ColoredPoint) Distance(q Point) float64 {
    return p.Point.Distance(q)
}

func (p *ColoredPoint) ScaleBy(factor float64) {
    p.Point.ScaleBy(factor)
}
```

匿名字段可以是指向命名类型的指针，字段和方法间接地来自于所指向的对象。**这可以让我们共享通用的结构以及使对象之间的关系更加动态、多样化。**
我们将`ColoredPoint`的匿名字段改成指针类型，在对比一下和上面非指针类型的区别：
```go
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

func main() {
	var cp ColoredPoint
	cp.Point = &Point{}  // 匿名指针类型的默认值是nil，必须对其进行初始化
	cp.Point.X = 1  // 如果没有上面的那一行，执行报错：panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Println(cp.Point.X) // 1

	p := ColoredPoint{&Point{1, 1}, color.RGBA{255, 0, 0, 255}}  // 初始化是Point传地址
	q := ColoredPoint{&Point{5, 4}, color.RGBA{0, 0, 255, 255}}

	//fmt.Println(p.Distance(q)) // 编译错误：cannot use q (type ColoredPoint) as type Point in argument to p.Point.Distance
	fmt.Println(p.Distance(*q.Point)) // 5  实参传递时，要转化为值
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(*q.Point)) // 10
}
```

结构体类型也可以由多个匿名字段
```go
type ColoredPoint struct {
    Point
    color.RGBA
}

p := ColoredPoint{Point{1, 1}, color.RGBA{255, 0, 0, 255}}
```
当调用`p.ScaleBy`方法时，它会先查找`ColoredPoint`有没有声明这个方法，如果没有，再从其内嵌对象`Point`和`color.RGBA`上查找，再从`Point`和`color.RGBA`的内嵌对象上查找。当同一个查找级别中有同名方式时，编译器报错；
```go
type A struct {}
func (a A) Func() {}
type B struct {}
func (b B) Func() {}
type C struct {
    A
    B
}

func main() {
    c := C{}
    c.Func()  // 编译错误：ambiguous selector c.Func
}
```

**方法只能在命名的类型（比如Point）和指向他们指针（*Point）中声明，但内嵌帮助我们能够在未命名的结构体类型中声明方法。**

## 方法变量与表达式
我们可以将方法赋予一个**方法变量**，方法变量是一个函数，本质上会绑定到接收者上，可以理解为方法的引用，方法变量只要传递实参就可以调用成功。
```go
a := Point{1, 2}
b := Point{4, 6}
distanceFromA := a.Distance  // 方法变量
fmt.Println(distanceFromA(b))  // 5
origin := Point{0, 0}
fmt.Println(distanceFromA(origin)) // 2.23606797749979

scaleA := a.ScaleBy  // 方法变量
scaleA(2)
fmt.Println(a)  // {2, 4}
```

**方法表达式**与方法变量相似，区别是方法变量是由将类型声明的变量的方法赋予的，而方法表达式是有类型的方法赋予的，有点绕，看一下例子：
```go
a := Point{1, 2}
b := Point{4, 6}
distanceFromA := a.Distance  // 方法变量 由a的方法赋予
distance := Point.Distance // 方法表达式 由Point类型的方法赋予
```
方法的接收者会替换成函数的第一个参数
```go
fmt.Println(distanceFromA(b))  // 5 方法变量
fmt.Println(distance(a, b))  // 5 方法表达式 
fmt.Printf("%T\n", distance) // func(Point, Point) float64

// scale := Point.ScaleBy // 编译报错：nvalid method expression Point.ScaleBy (needs pointer receiver: (*Point).ScaleBy
scale := (*Point).ScaleBy
scale(&a, 2)
fmt.Println(a)
fmt.Printf("%T\n", scale)  // func(*Point, float64)
```

## 封装
控制变量和方法不能通过对象访问（私有），即为封装。Go语言中通过控制命名的大小写来实现，首字母大写的标识符可以被导出，小写的就不可以。因此，可以通过结构体来是实现封装，向调用者隐藏重要的数据和实现细节，防止非法更改。
```go
type IntSet struct {
    words []uint64
}

type IntSet2 []uint64
```
对比两个类型，`IntSet`将实际存储数据的slice封装成了一个不可访问字段，`IntSet2`也将数据存储在slice，但它是可以被访问的，我们可以同*s在其他包中访问、更改。

> ### 思考：结构体里的字段一定都要封装起来，不让使用者看到吗？
>
> ### 封装的优点：
>- Go语言封装的单元是包而不是类型，包内的函数和方法对结构体的字段是可见的
>- 实现细节可以对包的使用方屏蔽，方便设计者灵活改变
>- 防止使用者非法更改结构体内的变量
> ### 封装的缺点：
>- 需要设计者编写很多的方法来实现对字段的读取和更新，因为调用者无法自助。

**封装并不总会需要的，要结合实际的适用场景区别对待。**