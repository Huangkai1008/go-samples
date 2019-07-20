// pipeline2 实现三级管道传输输出平方数，添加关闭通道操作
package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//	counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
			//time.Sleep(1 * time.Second)
		}
		close(naturals)
	}()

	// squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// printer
	for x := range squares {
		fmt.Println(x)
	}
}
