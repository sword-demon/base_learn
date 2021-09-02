package main

import (
	"net/http"
)

// Routable 单独的接口
type Routable interface {
	// Route 组合语法
	Route(method string, pattern string, handleFunc func(ctx *Context))
}

// Handler 定义接口
type Handler interface {
	http.Handler // 负责http请求
	Routable // 组合一个 路由
}

type HandlerBasedOnMap struct {
	// key 应该是 method + url
	handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedOnMap) Route(method string, pattern string, handleFunc func(c *Context)) {

	key := h.key(method, pattern)
	h.handlers[key] = handleFunc

	// 需要解决重复注册的问题
}

func (h *HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	key := h.key(request.Method, request.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Not found"))
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
