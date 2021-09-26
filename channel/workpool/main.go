package main

import (
	"fmt"
	"os"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job: %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	// 任务通道
	jobs := make(chan int, 100)
	// 结果通道
	results := make(chan int, 100)

	// 开启3个goroutine
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	// 发送5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	// 发送完就关闭
	close(jobs)

	// 输出结果
	for a := 1; a <= 5; a++ {
		<-results
	}

	fmt.Println("Commencing countdown")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // 读取单个字节
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown, Press return to abort")
	select {
	case <-time.After(10 * time.Second):
	// 不执行任何操作
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
}
