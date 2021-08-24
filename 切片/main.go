/*
 * @Author: your name
 * @Date: 2021-08-22 15:43:50
 * @LastEditTime: 2021-08-24 21:25:29
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /base_learn/切片/main.go
 */
package main

import (
	"fmt"
)

func qiepian() {
	v1 := []int{11, 22, 33, 44, 55, 66}

	v2 := v1[1:3]
	fmt.Println(v1) // [11 22 33 44 55 66]
	fmt.Println(v2) // [22, 33]
}

func DeleteMap() {
	v1 := []int{11, 22, 33, 44, 55, 66}
	deleteIndex := 2
	result := append(v1[:deleteIndex], v1[deleteIndex+1:]...)
	fmt.Println(result)
}

func InsertMap() {
	v1 := []int{11, 22, 33, 44, 55, 66}
	insertIndex := 3 // 在索引3的位置插入99

	/*
		result := make([]int, 0, len(v1)+1)
		result = append(result, v1[:insertIndex]...)
		result = append(result, 99)
		result = append(result, v1[insertIndex:]...)
		fmt.Println(result)
	*/

	// 错误思维
	result := append(v1[:insertIndex], 99)       // 相当于覆盖了v1里的44变成了99
	result = append(result, v1[insertIndex:]...) // [11, 22, 33, 99, 99 ,55, 66]
	fmt.Println(result)
}

func ForMap() {
	v1 := []int{11, 22, 33, 44, 55, 66}
	for i := 0; i < len(v1); i++ {
		fmt.Println(i, v1[i])
	}

	for index, item := range v1 {
		fmt.Println(index, item)
	}
}

func ParameterEqual() {
	v1 := []int{6, 9}
	v2 := v1 // 变量赋值

	v1[0] = 1111
	fmt.Println(v1, v2) // 内存地址不相同，但是内部指向数据的地方是相同的
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
