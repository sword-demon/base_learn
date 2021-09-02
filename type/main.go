/*
 * @Author: wxvirus
 * @Date: 2021-09-01 20:45:11
 * @LastEditTime: 2021-09-01 21:14:26
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /base_learn/type/main.go
 */
package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	//fmt.Println("type关键字学习")
	// 抽象出一个基于http的一个web服务
	server := NewHttpServer("test-server")
	//server.Route("GET", "/", SignUp)
	server.Route(http.MethodGet, "/", SignUp)
	err := server.Start(":8080")

	// 创建一个error的方法
	er := errors.New("abc")
	fmt.Println(er)

	if err != nil {
		// 处理err => 一步一判错那种
		panic(err) // 快速失败 一般用于表达非常严重不可恢复的错误
	}
}
