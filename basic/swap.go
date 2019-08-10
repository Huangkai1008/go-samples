package main

import "fmt"

// 交换函数
func swap(a, b *int) {
	//// 取a指针的值，赋给临时变量t
	//t := *a
	//
	//// 取b指针的值，赋给a指针指向的变量
	//*a = *b
	//
	//// 将a指针的值赋给b指针指向的变量
	//*b = t
	*a, *b = *b, *a
}

func main() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
}
