// cross命令打印了当前目标GOOS和GOARCH的值
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
