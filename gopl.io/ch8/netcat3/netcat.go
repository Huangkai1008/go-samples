// netcat3 实现一个只读的TCP 客户端程序，支持接收终端输入
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		mustCopy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} // 指示主goroutine
	}()
	defer conn.Close()
	mustCopy(conn, os.Stdin)
	<-done // 等待后台goroutine完成
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
