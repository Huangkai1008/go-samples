// 包tempconv提供摄氏和华氏度互转
package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9.0/5.0 + 32.0)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.0)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	_, _ = fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C":
		f.Celsius = Celsius(value)
		return nil
	case "F":
		f.Celsius = FtoC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
