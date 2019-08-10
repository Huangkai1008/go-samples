package main

import "fmt"

// 闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	for i := 0; i < 3; i++ {
		fmt.Println(nextInt())
	}

	newInt2 := intSeq()
	fmt.Println(newInt2())
}
