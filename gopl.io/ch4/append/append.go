package main

import "fmt"

// appendInt为整形切片添加元素
func appendInt(x []int, y int) []int {
	var z []int
	zLen := len(x) + 1
	if zLen <= cap(x) {
		// slice仍然有增长空间，扩展slice内容
		z = x[:zLen]
	} else {
		// slice已经没有空间，为他分配一个新的底层数组
		// 为了达到分摊线性复杂性，容量扩展一倍
		zCap := zLen
		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}
		z = make([]int, zLen, zCap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
