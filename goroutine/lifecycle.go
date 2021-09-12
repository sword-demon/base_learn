/*
 * @Author: your name
 * @Date: 2021-09-11 14:12:37
 * @LastEditTime: 2021-09-11 14:32:24
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /base_learn/goroutine/lifecycle.go
 */

package main

import (
	"context"
	"fmt"
	"sync"
)

func main3() {
	// 初始化一个context
	parent := context.Background()
	// 生成一个取消的context
	ctx, cancel := context.WithCancel(parent)

	runTime := 0

	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine Done")
				return
			default:
				fmt.Printf("Goroutine Running Times: %d\n", runTime)
				runTime += 1
			}
			if runTime > 5 {
				cancel()
				wg.Done()
			}
		}
	}(ctx)
	wg.Wait()
}
