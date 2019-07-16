// XmlSelect 输出XML 文档中指定元素下的文本
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // 元素名的栈
	for {
		token, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch token := token.(type) {
		case xml.StartElement:
			stack = append(stack, token.Name.Local) // 入栈
		case xml.EndElement:
			stack = stack[:len(stack)-1] // 出栈
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), token)
			}
		}
	}

}

func containsAll(x, y []string) bool {
	for len(y) < len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
