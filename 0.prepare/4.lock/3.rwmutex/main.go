package main

import (
	"fmt"
	"sync"
)

/**
互斥锁的核心思想在于每次只有一个并发实体能够访问共享的变量，不管是执行读操作还是写操作。
虽然能达到并发安全的需求，但是却给性能带来很大的影响，在上节的最后我们对此进行过测试。
只要使用锁就避免不了性能的损耗，除了控制锁的粒度尽可能小之外，还有一种办法可以减缓这种问题， 那就是这节要讲的读写锁。

在竞争条件那部分我们讲到数据竞争的条件是多个 goroutine 并发操作共享变量，并且至少一个操作为写。
后面这个条件非常关键，因为并发地读取操作并不会出现数据不一致的问题。可以利用这个特性把读操作 的锁和写操作的锁分开，
从而提升整个系统的性能。这就是读写锁的思想。

*/

type Account struct {
	name   string
	amount uint32
	mu     sync.RWMutex // 读写锁
}

func (a *Account) Deposit(amount uint32) {
	// 写锁
	a.mu.Lock()
	defer a.mu.Unlock()
	a.amount = a.amount + amount
}

func (a *Account) Balance() uint32 {
	// 读锁
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.amount
}

func main() {
	a := Account{name: "cizixs", amount: 0}
	var wg sync.WaitGroup

	wg.Add(10)
	// 启动 10 个 goroutine，并发往账户里存钱
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100000; j++ {
				a.Deposit(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(a.Balance())
}
