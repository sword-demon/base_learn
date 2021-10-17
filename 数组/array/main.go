package main

import "fmt"

func main() {
	/*
		数组
		数组是长度固定的数据类型,元素的类型都是一样的
		变量或者数组在定义时如果没有赋值，会有默认值，为零值
		整型： 默认值0
		字符串：默认值 空字符串
		bool： 默认值false
	*/

	var arrayVariables [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrayVariables)

	var length = len(arrayVariables)

	for i := 0; i < length; i++ {
		fmt.Printf("arrayVariables[%d]=%d,地址=%p\n", i, arrayVariables[i], &arrayVariables[i])
	}

	/**
	数组在内存中是连续的存储空间
	元素 下标			元素值		元素的地址
	0				1			0x1400001e0c0
	1				2			0x1400001e0c8
	2				3			0x1400001e0d0
	3				4			0x1400001e0d8
	4				5			0x1400001e0e0
	*/

	// 数组定义方式2
	// [10]int [5]int 被认为是不同的类型
	arrayVariables3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arrayVariables3, len(arrayVariables3))

	// 数组的定义方式3
	// 指定了数组下标 以最大的下标为数组的长度
	arrayVariables4 := [...]int{100: 200, 300: 500}
	fmt.Println(len(arrayVariables4))
}
