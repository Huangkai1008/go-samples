// 包bank提供单一账户的并发安全的银行
package bank

import "sync"

var (
	mu      sync.Mutex // 保护balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance = balance + amount
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}
