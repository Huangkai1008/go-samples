// clock3 是一个定期报告时间的TCP服务器，支持多个客户端并发调用，支持传入时区作为当地时钟
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.String("port", "8000", "server-port")

func main() {
	flag.Parse()
	fmt.Println(*port)
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // 例如，连接终止
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
