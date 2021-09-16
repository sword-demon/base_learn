package main

import "fmt"

func foo(i int, ch chan int) {
	ch <- i
	fmt.Println("send", i)
}

func main() {
	c := make(chan int)

	go foo(2, c)

	res := <-c

	fmt.Println("receive", res)
}
