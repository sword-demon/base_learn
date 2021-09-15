/*
 * @Author: your name
 * @Date: 2021-09-15 23:06:22
 * @LastEditTime: 2021-09-15 23:14:08
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /base_learn/functions/close_function.go
 */

// 闭包的累加器
package main

import "log"

func add(sum int) func(int) int {
	return func(i int) int {
		sum += i
		return sum
	}
}

func callCF() {
	f := add(10)
	sum := 0
	for i := 0; i < 10; i++ {
		sum = f(i)
		log.Printf("[闭包的累加器]:%d", sum)
	}
}

// 递归函数
func fb(i int) int {
	if i <= 0 || i == 1 {
		return i
	}
	return fb(i-1) + fb(i-2)

}

func main() {
	// callCF()

	for i := 0; i < 10; i++ {
		log.Printf("[%d=%d]", i, fb(i))
	}
}
