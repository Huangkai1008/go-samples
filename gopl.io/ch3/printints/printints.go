package main

import (
	"bytes"
	"fmt"
	"log"
)

// intSliceToString和fmt.Sprint(values)类似，但是插入了括号
func intSliceToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		_, err := fmt.Fprintf(&buf, "%d", v)
		if err != nil {
			log.Fatal("Write String to Buffer Error")
		}
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intSliceToString([]int{1, 2, 3}))
}
