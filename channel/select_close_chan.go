package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var quitC = make(chan struct{})

func signalWork() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// 当c中读取到值得时候，说明有人发了信号
	sig := <-c

	// 通知所有读取quitC的任务
	close(quitC)
	log.Printf("接收到了停止的信号接收, 信号是: %v, pid=%d, 要退出了", sig, os.Getppid())
}

func worker01() {
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-ticker.C:
			log.Printf("[我是worker01][5秒周期到了，干活]")
		case <-quitC:
			log.Printf("[我是worker01][接收到主进程退出的信号]，进行清理操作")
			return
		}
	}
}

func worker02() {
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-ticker.C:
			log.Printf("[我是worker02][5秒周期到了，干活]")
		case <-quitC:
			log.Printf("[我是worker02][接收到主进程退出的信号]，进行清理操作")
		}
	}
}

func worker03() {
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-ticker.C:
			log.Printf("[我是worker03][5秒周期到了，干活]")
		case <-quitC:
			log.Printf("[我是worker03][接收到主进程退出的信号]，进行清理操作")
		}
	}
}

func main() {
	go worker01()
	go worker02()
	go worker03()

	signalWork()
}
