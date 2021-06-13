package tempconv

import "fmt"

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度
type Kelvin float64     // 开氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点文档
	BollingC      Celsius = 100     //沸水温度
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }
