// bank包实现了一个只有一个账户的并发安全银行
// 使用二进制信号量控制并发数量
package bank

var (
	sema    = make(chan struct{}, 1) // 用来保护balance的二进制信号量
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // 获取令牌
	balance = balance + amount
	<-sema // 释放令牌
}

func Balance() int {
	sema <- struct{}{} // 获取令牌
	b := balance
	<-sema // 获取令牌
	return b
}
