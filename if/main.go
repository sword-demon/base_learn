package main

import "fmt"

// if 判断
func main() {
	//var score = 65
	//
	//// 基本写法
	//if score >= 90 {
	//	fmt.Println("A")
	//} else if score > 75 {
	//	fmt.Println("B")
	//} else {
	//	fmt.Println("C")
	//}

	// 2. 特殊写法 只在if判断里生效
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
