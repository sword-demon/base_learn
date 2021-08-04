package main

import (
	"fmt"
)

func main() {
	nums := [3]int32{11, 22, 33}
	fmt.Printf("数组的内存地址%p\n", &nums)
	fmt.Printf("数组的第1个元素的地址%p\n", &nums[0])
	fmt.Printf("数组的第2个元素的地址%p\n", &nums[1])

	names := [2]string{"wujie", "virus"}
	fmt.Printf("内存地址%p\n", &names)
	fmt.Printf("内存地址%p\n", &names[0])
	fmt.Printf("内存地址%p\n", &names[1])

	// 数组的切片
	data := nums[0:2] // 获取下标大于等于0小于2
	fmt.Println(data)

	// 数组的循环
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// for range 循环
	for key, item := range nums { // 如果只写一个，获取的只能是索引
		fmt.Println(key, item)
	}

	for _, item := range nums {
		fmt.Println(item)
	}

}
