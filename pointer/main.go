package main

import "fmt"

func main() {
	// 声明 一个  字符串类型 的变量，值为 wujie
	var name string = "wujie"

	// 声明一个 字符串的指针类型 的变量，值为name对应的内存地址
	var pointer *string = &name // 0x14000010240

	fmt.Println(pointer)

	v1 := "wujie"

	v2 := &v1

	fmt.Println(v1, v2, *v2) // wujie

	v1 = "666"               // 仅仅是把值改变，内存地址并没有变，所以指针也没有变
	fmt.Println(v1, v2, *v2) // 666 0x14000104230 666

	// 声明一个指针类型的变量pl，内部存储name的内存地址
	var pl *string = &name

	// 声明一个指针的指针类型变量p2,内部存储指针pl的内存地址
	var p2 **string = &pl

	// 声明一个指针的指针的指针类型变量p3,内部存储的是指针p2的内存地址
	var p3 ***string = &p2

	fmt.Println(pl, p2, p3)
	fmt.Println(*pl, *p2, *p3)
	fmt.Println(&pl, &p2, &p3)
}
