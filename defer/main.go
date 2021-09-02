package main

import "fmt"

// return之前panic之后执行
// 类似finally
// 类似一个栈的概念 先进后出 先defer的后执行

// 不要再循环里使用

func main() {
	defer func() {
		fmt.Println("aaa")
	}()

	defer func() {
		fmt.Println("bbb")
	}()

	defer func() {
		fmt.Println("ccc")
	}()
}
