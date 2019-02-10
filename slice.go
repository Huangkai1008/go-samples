package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 3}
	fmt.Println(a, a[1:2])

	// 切片扩容
	var numbers []int

	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}

	for _, v := range numbers {
		fmt.Println(v)
	}

	// 批量增加元素
	team := []int{1, 3, 7}

	numbers = append(numbers, team...)

}
