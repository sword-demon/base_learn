package main

import "sync"

var mutex sync.Mutex
var rwMutex sync.RWMutex

func Mutex() {
	mutex.Lock()

	// 你的业务代码
	// 然后你这里 panic 掉了
	// 然后就不会解锁了
	//mutex.Unlock()


	defer mutex.Unlock() // 一般来说解锁会放在defer中
}
