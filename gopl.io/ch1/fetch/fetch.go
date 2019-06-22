// fetch 输出从URL获取的内容
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// 假如URL参数缺失协议前缀，添加一个http://前缀
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// 使用io.Copy方法节省内存
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		resp.Body.Close()
		// 添加状态码
		fmt.Println(resp.StatusCode)
	}
}
