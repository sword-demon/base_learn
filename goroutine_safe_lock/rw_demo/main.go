package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x    int64
	wg   sync.WaitGroup
	rwLock sync.RWMutex
)

func read() {
	rwLock.RLock() // 读锁
	time.Sleep(time.Millisecond)
	rwLock.RUnlock() // 是否读锁
	wg.Done()
}

func write() {
	rwLock.Lock()
	x++
	time.Sleep(time.Millisecond * 10)
	rwLock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
