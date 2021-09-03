package main

import (
	"fmt"
	"time"
)

type Filter func(c *Context)

//type Filter1 interface {
//	Filter(c *Context)
//}
//
//type server struct {
//	filters []Filter1 // 遍历进行调度
//}
//
//type Interceptor interface {
//	Before(C *Context)
//	After(C *Context)
//	Surrounding(c *Context)
//	OnError()
//	OnResponse()
//	OnReturn()
//}

type FilterBuilder func(next Filter) Filter

var _ FilterBuilder = MetricsFilterBuilder

func MetricsFilterBuilder(next Filter) Filter {
	return func(c *Context) {

		start := time.Now().Nanosecond()
		// 调用下一个filter
		next(c)
		end := time.Now().Nanosecond()

		fmt.Printf("执行时间: %d", end-start)
	}
}
