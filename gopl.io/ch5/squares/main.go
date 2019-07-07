package main

import "fmt"

// squares 函数返回一个函数，后者包含下次要用到的平方数
func squares() func() int {
	var x int
	x++
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
