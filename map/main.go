/*
 * @Author: wxvirus
 * @Date: 2021-08-26 23:59:51
 * @LastEditTime: 2021-08-27 00:03:29
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
}
