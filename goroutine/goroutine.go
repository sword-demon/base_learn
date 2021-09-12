/*
 * @Author: your name
 * @Date: 2021-09-11 10:42:34
 * @LastEditTime: 2021-09-11 14:12:27
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /base_learn/goroutine/goroutine.go
 */

package main

import (
	"fmt"
	"time"
)

func action() {
	fmt.Println("test goroutine")
}

func main1() {
	go action()

	// 沉睡两秒
	time.Sleep(2 * time.Second)
}
