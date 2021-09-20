package main

import "fmt"

// 自定义类型和类型别名

// MyInt 1. 自定义类型
type MyInt int // 具有和 int 的一些相同的特性

// NewInt 类型别名
type NewInt = int

func main() {
	var i MyInt

	// type: main.MyInt, value: 0
	fmt.Printf("type: %T, value: %d\n", i, i)

	var x NewInt
	fmt.Printf("type: %T, value: %d\n", x, x)
}
