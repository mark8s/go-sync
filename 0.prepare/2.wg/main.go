package main

import (
	"fmt"
	"sync"
)

func sayHelloWorld(wg *sync.WaitGroup, arg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(arg)
	}
	wg.Done()
}

/*func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// 一定要传指针，不然会死锁
	go sayHelloWorld(&wg, "hello")
	go sayHelloWorld(&wg, "world")

	wg.Wait()
}
*/
