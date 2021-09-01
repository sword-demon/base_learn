/*
 * @Author: wxvirus
 * @Date: 2021-08-26 23:59:51
 * @LastEditTime: 2021-08-28 19:57:21
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /base_learn/map/main.go
 */
package main

import "fmt"

func main() {
	data := map[string]string{"n1": "wujie", "n2": "dwqdqw"}
	val := len(data)
	fmt.Println(val) // 2

	info := make(map[string]string, 10) // 默认有10个空余的地方来创建键值对  虚拟的容量
	fmt.Println(info)

	info["n1"] = "无解"
	info["n2"] = "Adqwdqw"

	fmt.Println(len(info)) // 2
	// fmt.Println(cap(info)) // 报错

	v7 := make(map[string][2]map[string]string) // 键是字符串，值是一个数组，数组里每个元素都是一个map
	v7["n1"] = [2]map[string]string{{"name": "wujie", "age": "18"}, {"name": "wujie1", "age": "19"}}
	fmt.Println(v7)

	data1 := map[string]string{"n1": "wujie", "n2": "dwqdqw"}

	// 循环
	for key, value := range data1 {
		fmt.Println(key, value)
	}

	Recurring()
}

// Recurring 切片的元素去重
func Recurring() {
	s1 := []string{"abc", "def", "abc", "okok"}

	m := make(map[string]struct{})
	for _, i := range s1 {
		m[i] = struct{}{}
	}
	s2 := make([]string, 0)
	for k := range m {
		s2 = append(s2, k)
	}
	fmt.Println(s1, s2)
}
