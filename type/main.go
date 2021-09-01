/*
 * @Author: wxvirus
 * @Date: 2021-09-01 20:45:11
 * @LastEditTime: 2021-09-01 21:14:26
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /base_learn/type/main.go
 */
package main

func main() {
	//fmt.Println("type关键字学习")
	// 抽象出一个基于http的一个web服务
	server := NewHttpServer("test-server")
	server.Route("/", SignUp)
	server.Start(":8080")
}
