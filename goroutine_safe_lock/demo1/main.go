package main

import (
	"fmt"
	"sync"
)

// 多个goroutine并发操作全局变量x

var (
	x  int64
	wg sync.WaitGroup
	lock sync.Mutex
)

func add() {
	for i := 0; i < 5000; i++ {
		// 1. 拿到全局变量x
		// 2. 给这个值+1
		// 3. 加1后在赋值给全局变量x
		lock.Lock() // 加锁
		x++ // 对全局变量进行每次+1
		lock.Unlock() // 释放锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x) // 7180 5285 各种值
}
