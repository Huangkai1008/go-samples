package main

import "fmt"

func main() {
	var house = "Malibu Point 10880, 90265"

	ptr := &house

	// 打印ptr的类型
	fmt.Printf("ptr type is %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)

	// 对指针进行取值操作
	value := *ptr

	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)
}
