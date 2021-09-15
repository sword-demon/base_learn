/*
 * @Author: your name
 * @Date: 2021-09-15 22:53:17
 * @LastEditTime: 2021-09-15 22:55:10
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /base_learn/channel/time_cancel_context.go
 */

package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond

func main() {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		// 告诉你的context是被取消了还是超时
		fmt.Println(ctx.Err())
	}
}
