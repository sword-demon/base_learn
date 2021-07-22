package main

import "fmt"

func main() {
	var name string = "wujie"

	// 1. 索引获取字节
	fmt.Println(name[0]) // 119

	// 2. 切片获取字节空间
	fmt.Println(name[2:3]) // j

	// 3. 手动循环获取所有字节
	for i := 0; i < len(name); i++ {
		fmt.Println(i, name[i])
	}

	/*
		0 119
		1 117
		2 106
		3 105
		4 101
	*/

	// 4. for range循环获取所有字符
	for index, item := range name {
		fmt.Println(index, item, string(item))
	}

	/**
	1 117
	2 106
	3 105
	4 101
	*/

	// 5. 转换成rune集合
	dataList := []rune(name)
	fmt.Println(dataList[0], string(dataList[0])) // 119 w
}
