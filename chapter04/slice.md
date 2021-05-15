# slice
slice 表示用于**相同类型**元素的**可变长度**的序列。

slice有三个属性：指针、长度和容量。
> - 指针：slice存储数据的内部结构是数组，指针指向的是数组的地址
> - 长度：保存slice中的元素数量
> - 容量：slice中可容纳的元素数量，在像slice插入元素时，如果超过容量，会对slice进行扩容


## slice的基本操作
### 1.初始化
```go
// 方式一 通过字面量自己初始化
s := []int{}
months := []string{1: "January", /* ... */,  12: "December"}
var z []int

// 方式二 通过切片语法从数组或slice中生成
a := [10]int{1, 2, 3}
b := a[1: 3]
c := b[1:]

// 注意这种奇怪的方式，可以指定具体某个下标的值
e := []int{1:42, 55, 66, 77, 7:88}
fmt.Println(e) // [0 42 55 66 77 0 0 88]
```

### 2.访问和更改
```go
// 访问和数组一致，通过下标访问
a := []int{1, 2, 3}
fmt.Println(a[1])  // 2

// 更改
a[1] = 100
fmt.Println(a[1])  // 100
b := a[:]
b[1] = 101  // b和a指向的底层数据一样的，所以b更改会影响a
fmt.Println(b[1], a[1])  // 101 101
```

### 3.增加
```go
// 如果想对slice增加一个元素，使用append函数
a := []int{1, 2, 3}
a = append(a, 4)
fmt.Println(a) // [1 2 3 4]
// 使用len函数获取当前slice中元素数量
// 使用cap函数获取当前slice所能支撑的容颜
for i := 5; i < 15; i++ {
    a = append(a, i)
    fmt.Printf("len: %d\t cap: %d\n", len(a), cap(a))
}
/*OUTPUT:
[1 2 3 4]
len: 5   cap: 6
len: 6   cap: 6
len: 7   cap: 12
len: 8   cap: 12
len: 9   cap: 12
len: 10  cap: 12
len: 11  cap: 12
len: 12  cap: 12
len: 13  cap: 24
len: 14  cap: 24
*/
```
`append`函数会处理当当前slice对象的容量不足时，自动扩容（重新申请一块空间并将原来的元素复制过来）。**我无法保证原始的slice和调用append后的结果slice执向同一个底层数组。也无法断定旧slice上对元素的操作会或者不会影响新的slice元素。*****所以通常我们将appednd的调用结果再次赋值给传入append函数的slice。***
### 4.删除
```go
// 通过copy实现删除
func remove(slice []int, i int) []int {
    copy(slice[i:], slice[i+1:])
    return slice[:len(slice)-1]
}
a := []{1, 2, 3}

a = remove(a, 1)
fmt.Println(a) // [1 3]
```

### 5.翻转slice
```go
func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

a := []int{1, 2, 3}
reverse(a)
fmt.Println(a) // [3 2 1]
```

### 6.比较
slice无法使用`==`进行比较，只允许和`nil`进行比较。如果想检查slice是否为空，使用`len(s) == 0`。
```go
var s []int     // len(s) == 0, s == nil
s = nil         // len(s) == 0, s == nil
s = []int(nil)  // len(s) == 0, s == nil
s = []int{}     // len(s) == 0, s != nil
```


## slice的用途
### 1.实现栈
栈的特点是先进后出
```go
stack = []int{}
// 入栈
stack = append(stack, 1)
// 出栈
top := stack[len(stack)-1]
stack = stack[:len(stack)-1]

```