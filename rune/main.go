/*
 * @Author: your name
 * @Date: 2021-08-31 23:00:32
 * @LastEditTime: 2021-09-01 00:18:01
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /base_learn/rune/main.go
 */
package main

import "fmt"

func main() {
	// rune 直观理解就是字符
	// rune 本质是int32 一个rune四个字节
	// rune 不是byte，golang没有char类型
	// 非面试情况下可以理解为  rune = 字符
	// 接近一般语言的char和character
	a := make([]int, 0, 3)
	fmt.Println(a) // []
}
