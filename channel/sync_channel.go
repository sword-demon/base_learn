package main

import (
	"log"
	"time"
)

func main() {
	// 初始化一个类型为int的chan， 缓冲区为3
	data := make(chan int, 3)
	quit := make(chan bool) // 达成和最后的time.Sleep一样的效果。去阻塞主线程防止异步任务中有未处理完的就退出
	// 读取数据的任务
	go func() {
		for d := range data {
			// 2秒钟处理一个任务
			time.Sleep(2 * time.Second)
			log.Printf("接收到了数据，开始处理: %v", d)
		}
		log.Printf("data chan 关闭了， 但是我还有清理工作，等我5秒钟")
		log.Printf("data chan 关闭了， 我再读取几个值看看: %v", <-data)
		log.Printf("data chan 关闭了， 我再读取几个值看看: %v", <-data)
		log.Printf("data chan 关闭了， 我再读取几个值看看: %v", <-data)
		time.Sleep(5 * time.Second)
		// 本任务处理完了，告诉主线程可以退出了
		quit <- true
		//for {
		//	if r, ok := <-data; ok {
		//		log.Printf("接收到了数据，开始处理: %v", r)
		//	} else {
		//		log.Printf("chan 关闭了")
		//		break
		//	}
		//}
		//for r := range data {
		//	log.Printf("接收到了数据，开始处理: %v", r)
		//}
		//log.Printf("chan 关闭了")
		//for {
		//	r := <-data
		//	log.Printf("接收到了数据，开始处理: %v", r)
		//}
	}()

	// 写入数据
	data <- 1
	time.Sleep(2 * time.Second)
	data <- 2
	time.Sleep(2 * time.Second)
	data <- 3
	time.Sleep(2 * time.Second)
	data <- 4
	log.Printf("发送4")
	data <- 5
	log.Printf("发送5")
	data <- 6
	log.Printf("发送6")
	data <- 7
	log.Printf("发送7")
	data <- 8
	log.Printf("发送8")
	data <- 9
	log.Printf("发送9")
	close(data)
	//time.Sleep(1 * time.Second)
	<-quit // 等待异步里的发送信号任务处理完
	log.Printf("真正退出了")
}
