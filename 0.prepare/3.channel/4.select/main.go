package main

import (
	"fmt"
	"time"
)

/**
另外一个常见的需求是：从多个 channel 中读取数据，只要任意一个 channel 有数据就执行对应的逻辑。 我们不能一次去循环这些 channel，
因为第一个执行的逻辑只有运行完成才会继续运行后面的逻辑，而不是 预期的从多个 channel 中选择。
*/

/**
select使用场景：从多个chanel中读取数据

select 的特点：

1. select 可以跟多个 case 语句，以及可选的 default 语句
2. 如果没有 default 语句，并且case 语言都没有匹配，那么 select 会阻塞
3. 如果有 default 语句，并且case 语言都没有匹配，那么会执行default语句
4. 如果有多个case都可以执行，那么go 会随机选择一个执行
*/

func worker(done chan struct{}) {
	time.Sleep(7 * time.Second)
	done <- struct{}{}
}

/**
执行以下的代码会有如下结果：

ticker1
ticker2
ticker2
ticker1
ticker1
ticker2
ticker2
ticker1
ticker1
ticker2
ticker2
ticker1
ticker1

work done.

*/

func main() {

	done := make(chan struct{})
	ticker := time.NewTicker(1 * time.Second)
	ticker2 := time.NewTicker(1 * time.Second)

	go worker(done)

	for {
		select {
		// 执行一次case 其实已经执行了一次chanel的读取操作，已经取出了一次值
		case <-ticker.C:
			fmt.Println("ticker1")
		case <-ticker2.C:
			fmt.Println("ticker2")
		case <-done:
			fmt.Printf("\nwork done.\n")
			return
		}
	}
}
