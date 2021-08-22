/*
 * @Author: your name
 * @Date: 2021-08-22 15:43:50
 * @LastEditTime: 2021-08-22 16:13:09
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /base_learn/切片/main.go
 */
package main

import "fmt"

func qiepian() {
	v1 := []int{11, 22, 33, 44, 55, 66}

	v2 := v1[1:3]
	fmt.Println(v1) // [11 22 33 44 55 66]
	fmt.Println(v2) // [22, 33]
}

func main() {

	// 自动扩容
	v1 := make([]int, 1, 3)       // 默认长度=1 容量=3 默认值为0
	fmt.Println(len(v1), cap(v1)) // 获取长度和获取容量

	v2 := append(v1, 99) // 返回一个新的切片
	fmt.Println(v1)      // [0]
	fmt.Println(v2)      // [0, 99]

	v1[0] = 66

	fmt.Println(v1) // [66]
	fmt.Println(v2) // [66, 99]

	v3 := make([]int, 1, 3)
	v3 = append(v3, 999)
	fmt.Println(v3) // [0, 999]

	v4 := []int{11, 22, 33}
	fmt.Println(len(v4), cap(v4)) // 3 3
	v5 := append(v4, 44)
	fmt.Println(len(v4), cap(v4)) // 3 3
	fmt.Println(len(v5), cap(v5)) // 4 6

	qiepian()
}
