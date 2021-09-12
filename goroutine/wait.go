/*
 * @Author: your name
 * @Date: 2021-09-11 14:08:30
 * @LastEditTime: 2021-09-11 14:12:20
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /base_learn/goroutine/wait.go
 */

package main

import (
	"fmt"
	"sync"
)

func main2() {
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(num int) {
			fmt.Printf("Goroutine Test %d\n", num)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
