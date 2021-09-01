package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// Context 目前不会扩充，所以使用结构体
type Context struct {
	W http.ResponseWriter // 本身就是一个接口
	R *http.Request       // 是一个结构体
}

// ReadJson interface{} 可以接收任何类型的参数
func (c *Context) ReadJson(req interface{}) error {
	// 帮我读Body
	// 帮我反序列化
	r := c.R
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		return err
	}
	return nil
}

// WriteJson 核心返回内容方法
func (c *Context) WriteJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = c.W.Write(respJson)
	return err
}

func (c *Context) OkJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}

func (c *Context) SystemErrJson(resp interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, resp)
}

func (c *Context) BadRequestJson(resp interface{}) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}

// NewContext 对外暴露的一个创建Context的方法，不让外部知道如何生成一个Context
func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		R: request,
		W: writer,
	}
}
