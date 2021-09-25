package main

import "fmt"

func main() {
	defer_func()
}

func defer_func() {
	defer func() {
		fmt.Println("1")
	}()

	defer func() {
		fmt.Println("2")
	}()

	defer func() {
		fmt.Println("3")
	}()

	panic("我是panic")

	// panic 之后的就没有机会

	//defer func() {
	//	fmt.Println("4")
	//}()

	/**
	3
	2
	1
	panic: 我是panic
	*/
}
