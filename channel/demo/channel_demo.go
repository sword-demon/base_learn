package main

import "fmt"

// 只能往里面发
func sender(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i // 发送i
	}
	// 发送玩就关闭 ch
	close(ch)
}

// 从ch1取值，把结果发送个ch2
// ch1 只能取
// ch2 只能发
func receiver(ch1 <-chan int, ch2 chan<- int) {
	// 从channel中取值的方式1
	for {
		tmp, ok := <-ch1
		// 100个值取完了，ok => false 就代表取完了
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 200)

	go sender(ch1)
	go receiver(ch1, ch2)

	// 从channel中取值的方式2
	for ret := range ch2 {
		// 内部会判断取值遇到了false就会退出
		fmt.Println(ret)
	}
}
