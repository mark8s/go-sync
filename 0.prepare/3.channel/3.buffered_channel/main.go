package main

import (
	"fmt"
	"sync"
	"time"
)

/**
提到并发模型，最经典的就是生产者和消费者问题。 这个问题的描述是这样的：有一个生产数据的实体，称为生产者；
另外有一个消费数据的实体，称为消费者，它们之间通过固定大小的队列作为缓存来通信。生产者把产生数据，并把数据放到队列中；
同时，消费者从队列中取出数据，执行任务。而且，生产者在队列满的时候不会继续 往里面放数据，消费者在队列空的时候不能从里面读数据。
*/

/*
**
以下是一个生产者、多个消费者的场景
*/
func producer(ch chan string) {
	for {
		ch <- "百岁山"
		fmt.Println("自来水厂生产了一箱水")
		time.Sleep(1000 * time.Millisecond)
	}
}

func consumer(ch chan string, i int) {
	for {
		fmt.Printf("买家:%d 买了一箱水%s\n", i, <-ch)
		time.Sleep(1500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(11)
	ch := make(chan string, 10)

	go func() {
		producer(ch)
		wg.Done()
	}()

	for i := 0; i < 10; i++ {
		go func(num int) {
			consumer(ch, num)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
