package main

import (
	"fmt"
	"time"
)

func main() {
	GoRoutine()
}

func GoRoutine() {
	// 用户级线程，又称之为 协程
	go func() {
		time.Sleep(10 * time.Second)
	}()

	// 这里直接输出，不会等待10秒
	fmt.Println("I am here")
}
