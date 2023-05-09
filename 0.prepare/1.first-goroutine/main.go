package main

import (
	"fmt"
	"time"
)

func sayHelloWorld(arg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(arg)
	}

}

func main() {
	go sayHelloWorld("hello")
	go sayHelloWorld("world")

	time.Sleep(2 * time.Second)
}
