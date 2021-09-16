package main

import (
	"fmt"
	"time"
)

func running() {
	var times int
	for {
		times++
		fmt.Println("tick", times)
		// 延时1秒
		time.Sleep(time.Second)
	}
}

func main() {
	go running()

	// 所有的goroutine在main函数结束时会一同结束
}