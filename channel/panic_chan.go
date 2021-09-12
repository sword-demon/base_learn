package main

func main() {
	// 声明
	// panic: close of nil channel
	var c1 chan int
	close(c1)

	// panic: close of closed channel
	c2 := make(chan int)
	close(c2)
	close(c2)

	// panic: send on closed channel
	c3 := make(chan int)
	close(c3)
	c3 <- 1
}