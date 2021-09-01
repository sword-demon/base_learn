/*
 * @Author: your name
 * @Date: 2021-09-01 21:20:13
 * @LastEditTime: 2021-09-01 21:47:51
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /base_learn/type/server.go
 */
package main

import (
	"fmt"
	"net/http"
)

// Server 接口里的方法不需要 func
// 尽量使用接口，以实现面向接口编程
// 接口是一组行为的抽象 一般接口普遍定义为大写，实现一般定义为小写
type Server interface {
	// Route 设定一个路由
	Route(pattern string, handleFunc http.HandlerFunc)
	// Start 启动我们的服务器
	Start(address string) error
}

// sdkHttpServer 基于http库实现的
type sdkHttpServer struct {
	Name string
}

// Route 做注册路由
func (s *sdkHttpServer) Route(pattern string, handleFunc http.HandlerFunc) {
	panic("implement me")
}

func (s *sdkHttpServer) Start(address string) error {
	panic("implement me")
}

type Header map[string][]string

type Node struct {
	// 自引用只能使用指针，指针的大小是固定的
	left  *Node
	right *Node
}

// 方法接收器，给结构体加方法
type User struct {
	Name string
	Age  int
}

// 方法接收器 结构体不会影响自身的状态
func (u User) ChangeName(newName string) {
	u.Name = newName
}

// 指针接收器，指针会对内部的状态做出影响
func (u *User) changeAge(newAge int) {
	u.Age = newAge
}

func UseFunc() {
	u := User{
		Name: "Tome",
		Age:  10,
	}

	u.ChangeName("Tom changed!") // Tome
	u.changeAge(100)             // 100
	fmt.Printf("%v\n", u)        // {Time 100}

	up := &User{
		Name: "jerry",
		Age:  12,
	}

	up.ChangeName("jerry changed") // jerry
	up.changeAge(120)              // 120
}

// 遇事不决 => 用指针

// type Handle func()

// func (h Handle) Hello() {

// }
