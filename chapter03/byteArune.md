## byte
`byte`，占用1个字节，和`uint8`等价。

## rune
`rune`，占用4个字节，和`uint32`等价。用于存储Unicode字符

```go
func main() {
	var a rune = '你'
	var b uint32 = 20320
	fmt.Printf("%c %d\n", a, a) // 你 20320, 占用字节：4
	fmt.Printf("%c %d\n", b, b) // 你 20320
}
```