package main

import "fmt"

// 类型别名

// 将NewInt定义为int类型
type NewInt int

// 将int取一个别名叫IntAlias
type IntAlias = int

func main() {
	var a1 NewInt
	fmt.Printf("a type: %T\n", a1)

	// 将a2声明为IntAlias类型
	var a2 IntAlias
	// 查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2)
}
