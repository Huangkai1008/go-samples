// echo4 输出命令的名字
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
}
