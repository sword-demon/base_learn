package main

import (
	"log"
	"sync"
	"time"
)

// HcMutex 是一个互斥锁
var HcMutex sync.Mutex

func runMutex(id int) {
	log.Printf("[任务id：%d][尝试获取锁]", id)
	HcMutex.Lock()
	log.Printf("[任务id：%d][获取到了锁]", id)
	time.Sleep(20 * time.Second) // 睡眠10秒钟
	HcMutex.Unlock()
	log.Printf("[任务id：%d][干完活了，释放锁]", id)
}

func runHcLock() {
	go runMutex(1)
	go runMutex(2)
	go runMutex(3)
}

func main() {
	runHcLock()

	time.Sleep(6 * time.Minute)

	/*
		2021/09/02 22:15:46 [任务id：3][尝试获取锁]
		2021/09/02 22:15:46 [任务id：3][获取到了锁]
		2021/09/02 22:15:46 [任务id：1][尝试获取锁]
		2021/09/02 22:15:46 [任务id：2][尝试获取锁]
		2021/09/02 22:15:56 [任务id：1][获取到了锁]
		2021/09/02 22:15:56 [任务id：3][干完活了，释放锁]
		2021/09/02 22:16:06 [任务id：1][干完活了，释放锁]
		2021/09/02 22:16:06 [任务id：2][获取到了锁]
		2021/09/02 22:16:16 [任务id：2][干完活了，释放锁]
	*/
}
