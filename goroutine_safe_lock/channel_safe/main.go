package main

import "sync"

var (
	//sema    = make(chan struct{}, 1) // 用来保护  balance的二进制信号量
	mu sync.Mutex
	balance int
)

func Deposit(amount int) {
	//sema <- struct{}{} // 获取令牌
	mu.Lock()
	balance = balance + amount
	//<-sema // 释放令牌
	mu.Unlock()
}

func Balance() int {
	//sema <- struct{}{} // 获取令牌
	mu.Lock()
	b := balance
	//<-sema // 释放令牌
	mu.Unlock()
	return b
}
