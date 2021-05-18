package main

import "fmt"

func main() {
	var name string
	var age int
	var isOk bool

	// 空字符串 0 false
	fmt.Println(name, age, isOk)

	// 声明变量的同时指定初始值
	var name1 string = "无解"
	var age1 int = 21
	fmt.Println(name1, age1)

	var name2, age2 = "virus", 19
	fmt.Println(name2, age2)
}
