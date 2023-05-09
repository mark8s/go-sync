package main

import (
	"fmt"
	"sync"
)

func work(arg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(arg)
	}
}

// wg的进阶使用，对业务代码没侵入

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		work("hello")
		wg.Done()
	}()

	go func() {
		work("world")
		wg.Done()
	}()

	wg.Wait()
}
