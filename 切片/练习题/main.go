/*
 * @Author: wxvirus
 * @Date: 2021-08-25 22:10:25
 * @LastEditTime: 2021-09-01 19:05:31
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

	fmt.Println(fb(1))
	fmt.Println(fb(2))
	fmt.Println(fb(3))
	fmt.Println(fb(4))

}

// 切片扩容，不再共享内存地址

func Add(s []int, index int, value int) []int {
	// todo: 切片的增加
	// 重新构建一个切片，比传入的切片的容量大1
	// 将指定的索引前面的内容追加到新的切片里
	// 将需要加入的值，追加到切片
	// 将原切片指定索引后面的值追加到切片
	// 最后形成一个新的切片返回，就实现了加的操作
	// todo : 需要对传入的index，和value值进行校验
	result := make([]int, 0, len(s)+1)
	result = append(result, s[:index]...)
	result = append(result, value)
	result = append(result, s[index:]...)
	return result
}

func Delete(s []int, index int) []int {
	// todo 切片的删除
	// s:=[]int{11,22,33,44} index = 2
	// [11, 22]
	// [44]
	// [11, 22, 44]
	// s => [11, 22, 44, 44] 被覆盖
	return append(s[:index], s[index+1:]...)
}

// 递归实现斐波那契数列
func fb(x int) int {
	if x <= 0 {
		return 0
	}
	if x == 1 || x == 2 {
		return x
	}
	return fb(x-1) + fb(x-2)
}

// 优化斐波那契数列
func fb1(x int) int {
	if x <= 0 {
		return 0
	}
	a, b := 1, 1
	result := 0
	for i := 3; i <= x; i++ {
		result = a + b
		a, b = b, result
	}
	return result
}
