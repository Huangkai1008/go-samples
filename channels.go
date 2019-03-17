package main

import "fmt"

func main() {
	messages := make(chan string) // 构造管道
	//使用 channel <- 语法 发送(send) 一个新的值到通道中。
	go func() { messages <- "ping" }()
	//使用 <-channel 语法从通道中 接收(receives) 一个值
	msg := <-messages
	fmt.Println(msg)
}
