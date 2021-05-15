# 结构体
**结构体**是将零个或多个任意类型的命名变量组合在一起的聚合数据类型。每个变量都叫做结构体的成员。

```go
type Employee struct {
    ID      int
    Name    string
    age     int
}
```
`Employee`就是一个结构体。

## 定义结构体时要注意 
1.如果一个成员变量的首字母大写，则它是可导出的。

2.成员变量的顺序对于结构体的同一性很重要，顺序不同表示不同的结构体。

3.结构体中的成员变量类型不能是结构体本身，但可以是该结构体的指针类型
```go
type Node struct {
    value   int
    next    *Node
    //这样是非法的
    //next  Node
}
```
4.结构体的零值由结构体成员的零值构成。

## 结构体变量的声明
```go
type Point struct {
    X   int
    Y   int
}

// 方式一
p1 := Point{1, 2} // X=1 Y=2
// 方式二
p2 := Point{X: 3, Y: 4} // X=3 Y=4

// 如果某个成员变量没有指定，那么该成员变量的值为该类型零值
p3 := Point{Y:6} // X=0 Y=6
```
方式一和方式二不可以混用。并且不能通过方式一绕过`不可导出变量无法在其他包中使用的规则`。
```go
// p.go
package p
type T struct { a, b int }

// q.go
package q
import "p"
var _ = p.T{1, 2}  // 编译错误
var _ = p.T{a: 1, b: 2} // 编译错误
```

如果在函数中需要修改结构体变量的内容时，需要传递指针
```go
type Point struct {
	X, Y int
}

func update(p Point, x int) { // 函数接收到的是实惨的一个副本
	p.X = x
}

func update2(p *Point, x int) {
	p.X = 2
}

func main() {
    p := Point{}
	update(p, 1)
	fmt.Println(p)  // {0 0}
	update2(&p, 2)
	fmt.Println(p)  // {2 2}
}
```


## 结构体的比较
必须所有的成员变量是可比较的，那这个结构体是可变比较
```go
type Point struct {
	X, Y int
}

type Book struct {
	name     string
	chapters []string
}

func main() {
	p1 := Point{X: 1, Y: 2}
	p2 := Point{X: 1, Y: 2}
	p3 := Point{X: 1, Y: 3}
	fmt.Println(p1 == p2, p2 == p3)  // true false

	b1 := Book{name: "abc", chapters: []string{"1", "2"}}
	b2 := Book{name: "abc", chapters: []string{"1", "2"}}
	fmt.Print(b1 == b2) // 编译错误，无法比较，因为chapters是slice类型 不能比较 
}
```

## 结构体嵌套和匿名成员
结构体嵌套的目的是为了复用
Go允许我们定义不带名称的结构体成员（即匿名成员），这样可以快捷访问到匿名成员里的成员变量
```go
type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8         // 等价于w.Circle.Point.X
	w.Y = 8
	w.Radius = 5    // 等价于w.Circle.Radius
	w.Spokes = 20

	fmt.Printf("%#v", w) // main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}
}

```
外围的结构体`Wheel`不仅获得匿名成员`Circle`的内部成员变量，也会获得`Circle`的方法
```go
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


type Wheel struct {
	Circle
	Spokes int
}

// func (w Wheel) Print() {
// 	fmt.Printf("%#v", w)
// }

func main() {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	fmt.Printf("%#v\n", w) // main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}
	w.Print() // main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}
}
```
如果在`Wheel`上再加一个匿名命名成员`Band`, 它也有`Print`方法
```go
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


func main() {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	fmt.Printf("%#v\n", w) // main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Band:main.Band{Name:""}, Spokes:20}
	w.Print() // 编译报错 
    w.Circle.Print()  // main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}
    w.Band.Print()  // main.Band{Name:""}
}
```
如果外围结构体`Wheel`有`Print`方法，会直接调用该方法；如果没有，但其内部的多个匿名成员都有该方法，需要**显示**指定调用哪个匿名结构体的`Print`方法
