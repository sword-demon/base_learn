/*
 * @Author: your name
 * @Date: 2021-09-12 23:18:44
 * @LastEditTime: 2021-09-13 00:21:52
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /base_learn/切片/slice.go
 */

package main

import "fmt"

func main() {
	var nums []int

	fmt.Printf("%T\n", nums)
	fmt.Println(nums == nil)

	// 字面量
	nums = []int{1, 2, 3}
	fmt.Printf("%#v\n", nums)
	nums = []int{1, 2, 3, 4}
	fmt.Printf("%#v\n", nums)

	// make函数
	a := make([]int, 3, 10)
	fmt.Printf("%#v %d %d\n", a, len(a), cap(a))

	n := a[1:3:10]
	// n_cap - start 10 - 1
	fmt.Printf("%T %#v %d %d\n", n, n, len(n), cap(n))

	m := a[2:3]
	// cap - start  10 - 2
	fmt.Printf("%T %#v %d %d\n", m, m, len(m), cap(m))

	nums = a[:]
	fmt.Printf("%T %#v %d %d\n", nums, nums, len(nums), cap(nums))

	// 删除

}
