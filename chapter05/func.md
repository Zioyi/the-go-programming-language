# 函数

## 函数声明
函数由5部分组成：函数名、形参列表、返回列表、和函数体。`func`为定义函数的关键字
```go
func name(parameters-list) (result-list) {
    body
}
```
形参列表的格式是参数名称+参数类型，相同类型的参数可以写在一起
```go
// 这两种个写法等价的
func f(x, y float64) float64 {}
func f(x float64, y float64) float64 {}
```
返回列表的格式是`(返回值名称+返回值类型...)`
返回值名称可以省，`当函数存在返回列表时，必须显示地已return语句结束`
```go
// 方式一：返回值名称+类型
func sum(x, y int) (z int) { 
	z = x + y  // 变量z已经被声明
	return // 必须显示地以return结束，可以不用指出返回的变量，因为函数第一行已写
}
// 方式二：只有返回类型
func sum2(x, y int) int {
    z := x + y  // 注意 这里z要初始化
    return z
}

// 如果有多个返回值，需要用括号包起来
func foo() (x, y int) {
	x, y = 1, 1
    return 
}
func foo2() (int, int) {
	x, y := 1, 1
	return x, y
}
```
函数的类型称作`函数签名`，由函数的形参列表和返回列表确定，形参和返回值名称不会形象函数类型
```go
func add(x, y int) int { return x + y }
func sub(a int, b int) (c int) { c = a - b; return }

fmt.Printf("%T\n", add)  // func(int, int) int
fmt.Printf("%T\n", sub)  // func(int, int) int
```
我们可以只定义函数签名，函数实现放在其他地方或其他语言
```go
func Sin(x float64) float64
```

## 函数的形参
形参变量是函数的局部变量。通常情况下，调用函数时实参是`按值`传递的，因此函数内修改变量不会改变实参的值。
```go
func incr(x int) {
	x++
}
func main() {
	a := 1
	incr(a)
	fmt.Println(a)  // 1
}
```
但是，如果实参是引用类型，比如：指针、slice、map、函数或者通道，那么就`有可能`改到实参的值
```go
func updateSlice(s []int, index int, val int) {
	s[index] = val
}

func updateMap(m map[int]int, k int, v int) {
	m[k] = v
}
func main() {
	s1 := []int{1, 1, 3, 4, 5}
	updateSlice(s1, 1, 2)
	fmt.Println(s1) // [1 2 3 4 5]

	m1 := map[int]int{0: 1, 1: 1, 2: 3}
	updateMap(m1, 0, 100)
	fmt.Println(m1) // basic_type[0:100 1:1 2:3]
}
```

## 函数的递归
递归实现斐波那契数列
```go
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fib(5)  // 5
}
```
递归的实现使用栈结构来保存当前上下文信息。Go语言的实现使用了可变长的栈，栈的长度可以随着使用增加。

## 函数多指返回
Go语言支持函数的返回值不止一个，一般情况是一个期望计算得到的结果和一个错误值或者一个表示函数调用是否正确的布尔值。
```go
func calculate(expr string) (result float64, err error) {
	...
}

func main() {
	res, err := calculate("30*50")
}
```

## 函数变量
Go语言中，可以声明函数类型的变量，即函数变量。函数变量之间不能比较，只能和`nil`比较。
```go
var sum func(int, int) int
fmt.Printf("%T\n", sum)  // func(int, int) int
sum = func(a, b int) int {
	return a + b
}
fmt.Println(sum(1, 1))  // 2

var f func(int, int) int
if f != sum {  // 编译错误  f != sum (func can only be compared to nil)
	f = sum
}
```
函数变量可以作为参数传递
```go
func add1(r rune) rune { return r + 1 }  // 将字符的Unicode值加1
fmt.Println(strings.Map(add1, "HAL-9000"))  // IBM.:111
```
可以在函数内部声明递归函数
```go
func main() {
	var fib func(int) int
	fib := func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(10))  // 5
}
/*下面的写法是错误的
func main() {
	func fib(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(10))  // 5
}
*/
```
## 变长函数
函数可以支持可变的参数数量，比如：`fmt.Printf`就是支持可变的数量。在参数列表最后的类型名称前使用`...`表示声明一个边长函数，下面我们来实现一个简易的`Sprintf`：
```go
func Sprintf(format string, params ...interface{}) string {
	i, j := 0, 0
	s := ""
	for i < len(format)-1 {
		if format[i] == '%' && format[i+1] == 'd' {
			s = s + strconv.Itoa(params[j])
			j++
			i++
		} else {
			s = s + string(format[i])
		}
		i++
	}
	return s
}

func main() {
	var s string
	s = Sprintf("%d+%d=%d", 1, 2, 3)
	fmt.Println(s)  // 1+2=3
	s = Sprintf("%d+%d+%d=%d", 1, 2, 3, 6)
	fmt.Println(s)  // 1+2+3=6
	fmt.Printf("%T", Sprintf) // func(string, ...int) string
}
```
可变长度参数只能声明在最后，并且只能有一个，这样就限制了可变参数只能是一种类型，但是`fmt.Printf`可以这样写：
```go
fmt.Printf("%d %s", 1, "abc")  // 1 abc
```
这是因为它将可变长度参数的类型声明生成了`interface{}`，将会在后面的章节研究。

## 延迟函数
在一个函数调用或者方法调用前加上`defer`关键字，就声明了这个函数（方法）延迟执行
- 延迟到return语句后执行
- 延迟到函数执行完毕后执行
- 延迟到发生宕机时执行
在一个函数作用域内，可以有声明多次延迟函数，执行的时候是以调用`defer`语句顺序的倒序进行。
延迟函数一般用于声明函数正常或异常结束后释放资源。
```go
conn, err := Client.GetConn()
defer coon.Close()
...
```
此外，还可以结合`闭包`实现对一个函数执行时的监控
```go
func clock(msg string) func() {
	start := time.Now()
	fmt.Printf("enter %s\n", msg)
	// 因为匿名函数可以得到其外层函数作用域内的变量（包括命名的结果)
	return func() { fmt.Printf("exit %s (%s)\n", msg, time.Since(start)) }
}

func SlowFunc() {
	defer clock("SlowFunc")()
	time.Sleep(3 * time.Second)
}
```

## 函数的宕机和恢复
宕机发生在程序的运行时出现了严重的异常情况，比如：错误的输入、配置或者I/O失败等。此时程序执行会终止，`goroutine`中的所有延迟函数会执行，然后程序会异常退出。
一些标准库会对`不可能发生`的情况做宕机处理，我们自己也可以同`宕机函数 panic`来实现：
```go
switch isRight {
	case true: //...
	case false: // ...
	default:
		panic("invalid")
}
```

有些情况下，当程序发生宕机，我们也不期望程序退出，比如，当Web服务器遇到处理用户请求时遇到宕机情况，不能直接退出，而是要给用户返回当前遇到的错误：
- 如果是用户查询的记录不存在，应该返回404
- 如果是用户输入的参数有问题，应该返回400
...

我们可以通过在函数的延迟函数中调用`recover`函数来终止当前的宕机状态并做一些逻辑处理
```go
func RequestHandler(c *Context) (res Response) {
	defer func() {
		switch p := recover(); p {
		case notFound{}:
			res = NotFoundRes{}
		case invalidParam{}:
			res = InvalidParamRes{}
		default:
			res = InternalErrorRes{}
		}
	}
	//具体处理逻辑
	//...
	return 
}
```

需要注意的是，要合理评估当前情况是否需要对宕机进行恢复，恢复会有一定风险，比如导致资源泄露或使失败的处理函数处于未定义的状态从而导致其他问题。