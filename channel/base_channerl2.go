package main

import "fmt"

func printer(c chan int) {
	for {
		data := <-c
		fmt.Println(data)
		if data == 0 {
			break
		}
	}
}

func main() {
	c := make(chan int)
	go printer(c)

	for i := 1; i <= 10; i++ {
		// 将数据通过channel投送给printer
		c <- i
	}

	// 通知并发的printer结束循环
	c <- 0

}
