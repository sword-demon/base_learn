package main

import "fmt"

func f4() (s1 []int) {
	s1 = []int{1, 1}
	defer func() {
		s1[1] = 10
	}()

	return []int{3, 3}
}

// 1. 先给返回值赋值
// 2. 把defer改造成正常的函数
// 3. 空的return

func f44() (s1 []int) {
	s1 = []int{1, 1}
	s1 = []int{3, 3}
	func() {
		s1[1] = 10 // 闭包 s1 [3, 10]
	}()

	return
}

func main() {
	res := f4()
	fmt.Println(res) // [3, 10]

	fmt.Println(f44()) // [3, 10]
}
