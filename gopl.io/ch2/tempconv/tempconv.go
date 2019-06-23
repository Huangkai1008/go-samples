// 包 tempconv 维护华氏温度和摄氏温度的基本信息
package tempconv

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreeingC      Celsius = 0
	BoilingC      Celsius = 100
)
