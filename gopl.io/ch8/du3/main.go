// du3命令计算文件目录中的磁盘使用情况
// 采用了并发限制计数信号
// 避免同时打开太多的文件
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 遍历在平行文件树的每一个根
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	// 定期输出结果
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nFiles, nBytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes关闭
			}
			nFiles++
			nBytes += size
		case <-tick:
			printDiskUsage(nFiles, nBytes)
		}
	}
	printDiskUsage(nFiles, nBytes)
}

func printDiskUsage(nFiles, nBytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nFiles, float64(nBytes)/1e9)
}

// walkDir递归遍历以dir为根目录的整个文件树
// 并在fileSizes上发送每个已找到的文件的大小
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirEntries(dir) {
		if entry.IsDir() {
			n.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walkDir(subDir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema用于限制并发数量的一个计数信号
var sema = make(chan struct{}, 20)

// dirEntries返回dir目录中的条目
func dirEntries(dir string) []os.FileInfo {
	sema <- struct{}{}        // 获取token
	defer func() { <-sema }() // 释放token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
