package wait

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// WaitForServer 尝试连接Url对应的服务器
// 在一分钟内使用指数退避策略进行尝试
// 所有的尝试失败后返回错误
func waitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // 成功
		}
		log.Printf("server not responding (%s); retrying ...", err)
		time.Sleep(time.Second << uint(tries)) // 指数退避策略
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
