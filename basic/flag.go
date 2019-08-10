package main

import (
	"flag"
	"fmt"
)

/**
指针变量获取命令行的输入信息
*/

// 定义命令行参数
var mode = flag.String("mode", "", "process mode")

func main() {
	// 解析命令行参数
	flag.Parse()

	fmt.Println(*mode)
}
