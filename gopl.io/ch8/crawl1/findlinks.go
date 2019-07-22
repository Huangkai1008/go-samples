// Crawl1使用命令行参数抓取web链接
package main

import (
	"fmt"
	"log"
	"os"

	"go-starter/gopl.io/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	workList := make(chan []string)

	go func() { workList <- os.Args[1:] }()

	// 并发获取网页
	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
			}
			go func(link string) {
				workList <- crawl(link)
			}(link)
		}
	}
}
