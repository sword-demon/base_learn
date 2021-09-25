package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello goroutine", i)
	wg.Done() // 告诉 main 函数 执行完了  通知 wg把计数器-1
}

func main() {

	//wg.Add(1) // 技数牌+1

	wg.Add(10000) // 一次全加满
	for i := 0; i < 10000; i++ {
		//wg.Add(1) // 或者每有一个goroutine加一个
		go hello(i)
	}

	fmt.Println("main goroutine done!")
	//time.Sleep(time.Second)

	// 等待别的goroutine干完活才结束
	wg.Wait() // 阻塞，等到计数器归零，就会结束
}
