/*
 * @Author: your name
 * @Date: 2021-09-15 23:06:22
 * @LastEditTime: 2021-09-15 23:35:16
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /base_learn/functions/close_function.go
 */

// 闭包的累加器
package main

import (
	"fmt"
	"log"
	"time"
)

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

func f2() (res int) {
	defer func() {
		res++
	}()

	// return xxx 这条语句不是一个原子操作
	// 先给返回值赋值
	// 然后调用defer语句
	// 最后返回调用函数中
	return 0
}

func f4() (t int) {
	defer func() {
		t = t * 10
	}()

	return 1
}

func main() {
	// callCF()

	// for i := 0; i < 10; i++ {
	// 	log.Printf("[%d=%d]", i, fb(i))
	// }

	start := time.Now()
	log.Printf("开始时间为: %v\n", start)
	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // 后进先出
	}
	// defer log.Printf("时间差: %v\n", time.Since(start)) // defer会立刻拷贝函数中引用的外部的参数，包括start和now

	defer func() {
		log.Printf("开始调用defer")
		log.Printf("时间差: %v\n", time.Since(start))
		log.Printf("结束调用defer")
	}()
	time.Sleep(3 * time.Second)
	log.Printf("函数结束")

	result := f4()
	fmt.Println(result)

	fmt.Println(f2())
}
