package main

import (
	"fmt"
	"time"
)

func main() {
	//channelWithCache()
	channelWithoutCache()
}

func channelWithCache()  {
	// 带1个缓冲的channel
	ch := make(chan string, 1)
	go func() {
		ch <- "Hello world, first msg from channel"
		time.Sleep(time.Second)
		ch <- "Hello second msg from channel"
	}()

	time.Sleep(2 * time.Second)
	msg := <- ch
	fmt.Println(time.Now().String() + msg)
	msg = <- ch
	fmt.Println(time.Now().String() + msg)
}

func channelWithoutCache() {
	// 不带缓冲
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second)
		ch <- "Hello. msg from channel"
	}()

	msg := <- ch
	fmt.Println(msg)
}
