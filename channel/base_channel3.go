package main

import "fmt"

func main() {
	// 构建一个通道
	ch := make(chan int)

	go func() {
		fmt.Println("start goroutine")
		ch <- 0
		fmt.Println("exit groutine")
	}()

	fmt.Println("wait groutine")

	// 等待匿名goroutine
	// channel的数据值如果没有接收的情况下，是阻塞的
	<-ch
	fmt.Println("all done")
}
