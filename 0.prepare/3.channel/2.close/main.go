package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
  在上一节，我们介绍了 channel 的基本概念，并使用 channel 模拟了打乒乓球的过程，最后留下了一个问题：怎么结束乒乓球游戏？ 这一节，我们将解答这个问题。
不妨我们先思考一下，现实中一场乒乓球是怎么结束的？无非是两种情况：一个是没接住对方的球；另外是接到了球，但是没有打到对方桌面上。
我们可以把这两种情况简化为一种情况：某个选手接到球之后，直接失败了。因为不同选手的能力不同，所以失败的几率也不同， 根据选手的接球成功率，
我们可以用简单的算法判断它每次接球是成功还是失败。 一方选手接球失败，另外一方就能直接看到，接收到这个消息，准备下一个回合的比赛，而不是像之前那样傻等着对方继续发球过来。
*/

func init() {
	// 每次运行用当前时间重置随机数生成器，增加随机性
	rand.Seed(time.Now().UnixNano())
}

type player struct {
	name         string
	successRatio int // a number between [0, 100]
}

func play(ch chan int, p *player) {
	for {
		ball, ok := <-ch
		// channel 已经关闭了，只有对方失败了才会关闭 channel，
		// 也就是说，当前队员赢得了游戏。
		if !ok {
			fmt.Printf("%s win!!!\n", p.name)
			return
		}
		// 生成一个 100 以内的随机数
		// 如果这个值大于成功率，则判定接球失败

		r := rand.Intn(100)
		if r > p.successRatio {
			// 失败之后，关闭 channel，然后退出函数
			fmt.Printf("%s lose.\n", p.name)
			close(ch)
			return
		}
		fmt.Printf("%d %s\n", ball, p.name)
		time.Sleep(time.Millisecond * 200)
		ball++
		ch <- ball
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ball := 1

	ch := make(chan int)

	// 如果放这里，会阻塞main goroutine 导致阻塞，直接死锁
	//ch <- 1

	go func() {
		play(ch, &player{
			name:         "mark",
			successRatio: 90,
		})
		wg.Done()
	}()

	go func() {
		play(ch, &player{
			name:         "james",
			successRatio: 80,
		})
		wg.Done()
	}()

	// 发球
	ch <- ball
	wg.Wait()
	fmt.Println("Game Over")
}
