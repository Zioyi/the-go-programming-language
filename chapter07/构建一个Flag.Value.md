## flag是什么
`flag`是Go的一个标准库，用于解析命令行

## flag.Value
flag.Value是一个接口类型，实现了它的类型值可以作为新的符号标记

```go
package flag

// Value is the interface to the value stored in a flag.
type Value interface {
	String() string
	Set(string) error
}
```