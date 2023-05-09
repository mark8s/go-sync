package main

import (
	"fmt"
	"sync"
	"time"
)

/***
为了说明 channel 的使用，我们写一个模拟打乒乓球的 go 程序。和之前打印字符串的程序不同，打乒乓球需要两个选手参与（goroutine），
而且他们要等待同一个乒乓球（同步数据），只有乒乓球被打到自己这边的时候才能接球，否则就要一直处于准备状态（等待）。

参考：https://cizixs.gitbook.io/go-concurrency-programming/theory/channel/04-use-channel-to-communicate
*/

/**
无缓冲channel说明：当往里面写数据的时候，如果没有人把它取走，那么这么操作会一直阻塞；反之亦然。 这样的 channel 要求两个 goroutine 必须在这个时间点同步，
就像两个人约好在某个地方碰面，一起去看电影。如果其中一个人先到，它 必须在这里等待，一直等到另外一个人来到，才能去看电影。

*/

// 1.最简单的情况
/*func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
}
*/

func play(ch chan int, name string) {
	for i := 0; i < 100; i++ {
		ball := <-ch
		fmt.Printf(name+" get ball: %d\n", ball)
		ball++
		time.Sleep(1 * time.Millisecond)
		// 最后一次执行后，chan 中塞了个值，如果chan是unbuffer 的，那么会导致死锁。
		// 两个goroutine，其中有一个会后面执行，前一个已经完全执行完了，就退出了，而后一个goroutine 又写了一个值到 channel中，但是这个channel 中的值已经没有goroutine去消费了，所以此写值的goroutine将阻塞。
		// 主 goroutine 无法处理这个goroutine的阻塞，导致死锁。
		// 关于 死锁的定义：死锁是当 Goroutine 被阻塞而无法解除阻塞时产生的一种状态。
		ch <- ball
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ball := 1

	ch := make(chan int, 0)

	// 如果放这里，会阻塞main goroutine 导致阻塞，直接死锁
	//ch <- 1

	go func() {
		play(ch, "mark")
		wg.Done()
	}()

	go func() {
		play(ch, "james")
		wg.Done()
	}()

	// 发球
	ch <- ball
	wg.Wait()
}
