// countdown3模拟火箭发射倒计时
// Note: 如果启动中止，那么goroutine代码永远不会终止
// 存在goroutine泄露
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		_, _ = os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.  Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing
		case <-abort:
			fmt.Println("Launch aborted")
			return
		}
	}
	launch()

}

func launch() {
	fmt.Println("Lift off!")
}
