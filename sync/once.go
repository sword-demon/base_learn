package main

import (
	"fmt"
	"sync"
)

func main() {
	PrintOnce()
	PrintOnce()
	PrintOnce()
	PrintOnce()
}

// 如果使用局部变量的时候，就会没效果。每次进来都会新建一个once
var once sync.Once

func PrintOnce() {
	once.Do(func() {
		fmt.Println("只输出一次")
	})
}