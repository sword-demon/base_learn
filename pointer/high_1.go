package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	dataList := [3]int8{11, 22, 33}
	fmt.Printf("数组的地址: %p;数组的第一个元素的地址: %p\n", &dataList, &dataList[0])
	// 数组的地址: 0x14000198000;数组的第一个元素的地址: 0x14000198000
	// 存的地址是一模一样的

	fmt.Println(reflect.TypeOf(&dataList))    // *[3]int8
	fmt.Println(reflect.TypeOf(&dataList[0])) // *int8

	// 指针的计算

	// 1. 获取数组第一个元素的地址(指针)
	var firstDataPtr *int8 = &dataList[0]

	// 2. 转换成Pointer类型
	ptr := unsafe.Pointer(firstDataPtr)

	// 3. 转换成uintptr类型，然后进行内存地址的计算，即地址加1个字节，意味着取第2个索引的位置的值
	targetAddress := uintptr(ptr) + 1

	// 4. 根据新地址， 重新转换成Pointer类型
	newPtr := unsafe.Pointer(targetAddress)

	// 5. Pointer对象转换为int8指针类型
	value := (*int8)(newPtr)

	// 6. 根据指针获取值
	fmt.Println("最终结果为: ", *value) // 22

}
