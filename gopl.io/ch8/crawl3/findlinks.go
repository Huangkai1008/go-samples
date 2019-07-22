// crawl3使用命令行参数抓取web链接
// 在这个版本中限制了goroutine的数量来限制并发数量
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
	workList := make(chan []string)  // url的列表，可能含有重复的对象
	unseenLinks := make(chan string) // 去重的网址

	go func() { workList <- os.Args[1:] }()

	// 创建20个并发的爬虫goroutine去爬取每一个未爬取的链接
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { workList <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
