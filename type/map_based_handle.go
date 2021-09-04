package main

import (
	"net/http"
)

// Routable 单独的接口
type Routable interface {
	// Route 组合语法
	Route(method string, pattern string, handleFunc handlerFunc)
}

// Handler 定义接口
type Handler interface {
	// ServeHTTP 	上一个代码: http.Handler // 负责http请求
	ServeHTTP(c *Context)
	Routable // 组合一个 路由
}

type HandlerBasedOnMap struct {
	// key 应该是 method + url
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedOnMap) Route(method string, pattern string,
	handleFunc handlerFunc) {

	key := h.key(method, pattern)
	h.handlers[key] = handleFunc

	// 需要解决重复注册的问题
}

func (h *HandlerBasedOnMap) ServeHTTP(c *Context) {

	key := h.key(c.R.Method, c.R.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Not found"))
	}
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}

// 在实现类之上加一个这个
// 确保 HandlerBasedOnMap 一定是实现了 Handler 的接口
var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		handlers: make(map[string]func(ctx *Context)), // 预估容量，考虑默认值的问题
	}
}
