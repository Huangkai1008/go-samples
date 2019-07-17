// clockwall 可以与多个clock服务器通信，从多服务器中读取时间
package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	for _, v := range os.Args[1:] {
		splits := strings.Split(v, "=")
		go connTcp(splits[1])
	}
	for {
		time.Sleep(1 * time.Second)
	}
}

func connTcp(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
