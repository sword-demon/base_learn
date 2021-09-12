/*
 * @Author: your name
 * @Date: 2021-09-11 14:32:39
 * @LastEditTime: 2021-09-11 16:00:30
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /base_learn/goroutine/first_select.go
 */

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func serverApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	})

	// 主线程 pending 8080端口 阻塞
	if err := http.ListenAndServe("0.0.0.0.8080", mux); err != nil {
		// 不建议在生产环境调用
		log.Fatal(err) // 会进行退出 整个程序就会退出 无条件终止
	}

}

func serverDebug() {
	http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux)
}

func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		<-stop // wait for stop signal 接收停止信号
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func main4() {

	//done := make(chan error, 2)
	//stop := make(chan struct{})
	//go func() {
	//	done <- serverDebug(stop)
	//}()
	//
	//go func() {
	//	done <- serverApp(stop)
	//}()
	//
	//var stopped bool
	//for i := 0; i < cap(done); i++ {
	//	if err := <-done; err != nil {
	//		fmt.Println("error: %v\n", err)
	//	}
	//	if !stopped {
	//		stopped = true
	//		close(stop)
	//	}
	//}

	// 异步执行 它如果退出，你会不知道
	// go serverDebug()

	// // 如果它因为别的原因断开了，就会产生问题

	// // 改进
	// go serverApp()

	// select {}

	// if err := http.ListenAndServe(":3000", nil); err != nil {
	// 	log.Fatal(err)
	// }

	// go func() {
	// 	if err := http.ListenAndServe(":3000", nil); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// // 空的select会永远阻塞，避免监听的代码一运行就挂掉了
	// select {}
}
