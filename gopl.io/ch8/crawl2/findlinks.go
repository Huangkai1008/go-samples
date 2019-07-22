// crawl2使用命令行参数抓取web链接
// 在这个版本中使用缓冲信道作为一个计数信号
// 限制并发调用links.Extract的数量
package main

import (
	"fmt"
	"log"
	"os"

	"go-starter/gopl.io/ch5/links"
)

// token(令牌)是一个计数信号用来进行最多20的强制并发请求
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // 请求令牌
	list, err := links.Extract(url)
	<-tokens // 释放令牌
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	workList := make(chan []string)
	var n int

	n++
	go func() { workList <- os.Args[1:] }()

	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-workList
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					workList <- crawl(link)
				}(link)
			}
		}
	}
}
