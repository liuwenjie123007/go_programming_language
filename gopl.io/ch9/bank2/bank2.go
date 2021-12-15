package bank2

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
	<-sema // 释放令牌
	return b
}
