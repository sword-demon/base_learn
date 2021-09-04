package main

import (
	"fmt"
	"sync"
	"time"
)

// 为了解决map线程不安全，自己加锁

// concurrentMap 造了一个类型
type concurrentMap struct {
	mp map[int]int
	sync.RWMutex
}

// Set 通过Set方法做原有map的赋值 m[key] = value
func (c *concurrentMap) Set(key, value int) {
	// 加写锁
	c.Lock()
	// 赋值
	c.mp[key] = value
	// 解锁
	c.Unlock()
}

// Get 通过Get方法做原有map的读取值 v := m[key]
func (c *concurrentMap) Get(key int) int {
	// 加读锁
	c.RLock()
	// 获取值
	res := c.mp[key]
	// 解读锁
	c.RUnlock()
	return res
}

func main() {
	c := concurrentMap{
		mp: make(map[int]int),
	}

	// 模拟线程循环的写map
	go func() {
		for i := 0; i < 10000; i++ {
			c.Set(i, i)
		}
	}()

	// 模拟线程循环读map
	go func() {
		for i := 0; i < 10000; i++ {
			res := c.Get(i)
			fmt.Printf("[cmap.get][%d=%d]\n", i, res)
		}
	}()

	time.Sleep(1 * time.Hour)
}
