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
	// 需要加上 method POST PUT GET DELETE
	Route(method, string, pattern string, handleFunc func(c *Context))
	// Start 启动我们的服务器
	Start(address string) error
}

// sdkHttpServer 基于http库实现的
type sdkHttpServer struct {
	Name string
	handler *HandlerBasedOnMap
}

// Route 做注册路由
func (s *sdkHttpServer) Route(method string, pattern string, handleFunc func(c *Context)) {

	//http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
	//	//ctx := &Context{
	//	//	W: w,
	//	//	R: r,
	//	//}
	//	ctx := NewContext(w, r)
	//	handleFunc(ctx)
	//})
	key := s.handler.key(method, pattern)
	s.handler.handlers[key] = handleFunc

	// 需要解决重复注册的问题
}

func (s *sdkHttpServer) Start(address string) error {
	http.Handle("/", s.handler)
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) *sdkHttpServer {
	return &sdkHttpServer{
		Name: name,
	}
}

func NewServer() *sdkHttpServer {
	return &sdkHttpServer{}
}

// Factory 小型工程模式设计方法
type Factory func() Server

var factory Factory

func RegisterFactory(f Factory) {
	factory = f
}
func NewServerA() Server {
	return factory()
}
type Node struct {
	// 自引用只能使用指针，指针的大小是固定的
	left  *Node
	right *Node
}

// User 方法接收器，给结构体加方法
type User struct {
	Name string
	Age  int
}

// ChangeName 方法接收器 结构体不会影响自身的状态
func (u User) ChangeName(newName string) {
	u.Name = newName
}

// changeAge 指针接收器，指针会对内部的状态做出影响
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

func SignUp(ctx *Context) {
	req := &signUpReq{}

	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}
	resp := &commonResponse{
		Data: 123,
	}
	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
}

type signUpReq struct {
	// Tag 倾向于做声明式API
	Email             string `json:"email"` // 可以进行反序列化， 指定json输出的内容
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
