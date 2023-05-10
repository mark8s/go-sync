package main

import (
	"fmt"
	"sync"
)

/***
解决数据竟态问题：

1. 不要在多个 goroutine 中读写变量
2. 运行多个 goroutine 读写变量，但每次保证操作的时候只有一个 goroutine 能操作。其实也就是互斥锁。

*/

// Account 代表某个人的银行账户，有用户名和余额两个字段
type Account struct {
	name   string
	amount uint32
}

// Deposit 往账户里面存特定数量的钱
func (a *Account) Deposit(amount uint32) {
	// 会有数据竟态问题，多个goroutine 操作。导致不能计算处理的值不正确
	a.amount = a.amount + amount
}

// Balance 返回账户里还有多少余额
func (a *Account) Balance() uint32 {
	return a.amount
}

func worker(a *Account) {
	for i := 0; i < 100000; i++ {
		a.Deposit(1)
	}
}

func main() {
	a := &Account{name: "Alice", amount: 0}
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		worker(a)
		wg.Done()
	}()

	go func() {
		worker(a)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(a.Balance())
}
