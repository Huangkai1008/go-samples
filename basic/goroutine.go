package main

import "fmt"

// 协程

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
func main() {
	f("direct")

	go f("goroutine")

	// 匿名函数调用协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	_, _ = fmt.Scanln()
	fmt.Println("done")
}
