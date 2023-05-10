package main

import (
	"sync"
	"testing"
)

/*
*
具体结果会因为机器的配置以及每次运行时系统的负载不太相同，但是可以从上面两个简单的结果看出，
加锁之后每次操作的的时间从 0.37ns 变成了 75.9 ns，
如此简单的例子就能带来这么大的性能差距，
在实际上更负责的代码中，锁机制带来的性能损失可能会更严重，
是我们必须要考虑中的事情。在下一节中，我们将介绍如何使用读写锁来减少某些情况下锁机制带来的性能损耗。
*/
func BenchmarkAccountRead(b *testing.B) {
	a := Account{name: "cizixs", amount: 0}
	for i := 0; i < b.N; i++ {
		a.Balance()
	}
}

func BenchmarkAccountReadWithLock(b *testing.B) {
	a := Account{name: "cizixs", amount: 0}
	var mu sync.Mutex
	for i := 0; i < b.N; i++ {
		mu.Lock()
		a.Balance()
		mu.Unlock()
	}
}
