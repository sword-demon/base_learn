package main

import (
	"fmt"
	"time"
)

/*
构建一个通道
开启一个并发匿名函数
从3循环到0，并且发送3到0之间的数据，每次发送完值需要等到1秒钟
主程序main函数中遍历接收通道数据，然后打印出来
当遇到数据为0时，退出接收循环
*/

func main() {

	ch := make(chan int)

	go func() {
		for i := 3; i >= 0; i-- {
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for data := range ch {
		fmt.Println(data)
		if data == 0 {
			break
		}
	}
}
