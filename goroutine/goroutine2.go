package main

import (
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup

func main() {
	wg2.Add(10000) // 一次全加满
	for i := 0; i < 10000; i++ {
		// wg2.Add(1) // 或者每有一个goroutine加一个
		go func(i int) {
			// 换成匿名函数(闭包) 包含了一个外部函数的一个变量的引用
			fmt.Println("hello", i)
			wg2.Done() // 都执行完了，通知结束
		}(i) // 此时的i是每次for循环的i传进来的 副本
	}

	fmt.Println("main goroutine done!")

	// 等待别的goroutine干完活才结束
	wg2.Wait() // 阻塞，等到计数器归零，就会结束
}
