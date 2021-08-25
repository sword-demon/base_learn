/*
 * @Author: wxvirus
 * @Date: 2021-08-25 22:10:25
 * @LastEditTime: 2021-08-25 23:00:25
 * @LastEditors: Please set LastEditors
 * @Description: 切片练习题
 * @FilePath: /base_learn/切片/练习题/main.go
 */
package main

import "fmt"

func main() {
	v1 := make([]int, 2, 5)
	v2 := append(v1, 123)

	fmt.Println(len(v1), cap(v1)) // 2 5
	fmt.Println(len(v2), cap(v2)) // 3 5

	// 切片进行扩容的时候，会重新开辟一个内存，进行存储，此时数据就和原来的变量不一样

	v3 := [][]int{{11, 22, 33, 44}, {44, 55}}
	v3[0][2] = 99
	fmt.Println(v3)

	v4 := [][]int{{11, 22, 33, 44}, {44, 55}}
	v5 := append(v4[0], 99)
	v5[2] = 69
	fmt.Println(v4)
	fmt.Println(v5)

	v6 := [][]int{make([]int, 2, 5), make([]int, 2, 3)}
	v7 := append(v6[0], 99)
	v7[0] = 69
	fmt.Println(v6, len(v6), cap(v6))
	fmt.Println(v7, len(v7), cap(v7))

}
