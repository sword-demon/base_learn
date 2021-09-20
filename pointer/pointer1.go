package main

import "fmt"

func modify1(x int) {
	x = 100
}

func modify2(y *int) {
	*y = 100 // 取到指针地址的值进行 修改
}

func main() {
	a := 1
	modify1(a)
	fmt.Println(a) // 1

	modify2(&a)
	fmt.Println(a) // 100
}
