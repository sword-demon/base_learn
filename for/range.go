package main

import (
	"fmt"
	"log"
)

func main() {
	a1 := make([]*int, 3)
	a2 := make([]int, 3)

	for k, v := range []int{1, 2, 3} {
		// v在每次遍历的时候都是一个新的变量
		v := v
		log.Printf("[v的值: %v][v的地址: %p]", v, &v)
		a1[k] = &v
		a2[k] = v
	}

	for i := range a1 {
		fmt.Println("指针切片的值为", *a1[i])
	}

	for i := range a2 {
		fmt.Println("铺路值类型的切片的值为: ", a2[i])
	}
}