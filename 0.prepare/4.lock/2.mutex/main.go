package main

import (
	"fmt"
	"sync"
)

/**
互斥锁

锁机制是传统的编程语言（比如 Java、C 等）对于并发程序的解决方案，线程读写数据区之前先进行加锁操作，
只能加锁成功才能执行读写逻辑，执行完成后需要释放锁， 以供其他线程能够使用。当某个线程执行加锁动作，
其他想要执行相同加锁操作的线程只能等待，等到锁解开后才能继续运行。

注意：加锁和解锁的动作必须是成对出现的，如果某个线程只执行加锁操作，但是忘记执行解锁操作，
那么所有要读写关键区变量的进程都会 一直处于阻塞的状态，这被称为死锁，在后面的章节我们会介绍死锁的检测。
*/

// 改善上一个例子，这次我们使用 sync.mutex ，它是go中的互斥锁
// Account 代表某个人的银行账户，有用户名和余额两个字段
type Account struct {
	name   string
	amount uint32
	mu     sync.Mutex
}

// Deposit 往账户里面存特定数量的钱
func (a *Account) Deposit(amount uint32) {
	a.mu.Lock()
	a.amount = a.amount + amount
	a.mu.Unlock()
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
