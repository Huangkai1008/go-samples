// strconv用来演示互相转换字符串和数字类型
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串转换为整数
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	// 按进制格式化数字
	fmt.Println(strconv.FormatInt(int64(x), 2))

	// 整数转换为字符串
	m, _ := strconv.Atoi("123")
	n, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(m, n)
}
