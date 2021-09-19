package concurrency

import (
	"fmt"
	"sync"
)

/*
单例模式
*/

type Single interface {
	Single()
}

type singleton struct {
	// 有很多字段
}

func (s *singleton) Single() {
	fmt.Println("I am single")
}

var instance *singleton
var instanceOnce sync.Once

func GetSingleInstance() *singleton {
	instanceOnce.Do(func() {
		instance = &singleton{}
	})
	return instance
}

type MyBiz struct {
	once sync.Once
}

// Init 让这个初始化方法只执行一次
func (m *MyBiz) Init() {
	m.once.Do(func() {

	})
}
