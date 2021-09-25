package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg3 sync.WaitGroup

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("a", i)
	}
	wg3.Done()
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("b", i)
	}
	wg3.Done()
}

func c() {
	for i := 0; i < 10; i++ {
		fmt.Println("c", i)
	}
	wg3.Done()
}

func main() {
	runtime.GOMAXPROCS(6) // 只占用一个CPU核心

	wg3.Add(2)
	go a()
	go b()

	wg3.Wait()
	//time.Sleep(time.Second)
}
