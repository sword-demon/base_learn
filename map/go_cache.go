package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

// 生产环境的web缓存应用
/*
维护用户信息的模块
在mysql中有一个user表  更新 不会太频繁
正常情况下是用orm gorm xorm 去db中查询
查询qps很高
在一定时间内获取到旧数据也能容忍
*/

type user struct {
	Name  string
	Email string
	Phone string
}

var (
	DefaultInterval = time.Minute * 1
	UserCache       = cache.New(DefaultInterval, DefaultInterval)
)

// GetUser 优先去本地缓存中查，有就返回，没有再去db查询，暂时用下面的方法
func GetUser(name string) user {
	if res, found := UserCache.Get(name); found {
		log.Printf("本地缓存中找到对应的用户[name: %v][value: %v]\n", name, res.(user))
		return res.(user)
	} else {
		res := HttpGetUser(name)
		log.Printf("本地缓存中没找到对应的用户，去远端查询[name: %v][value: %v]\n", name, res)
		// 更新本地的缓存
		UserCache.Set(name, res, DefaultInterval)
		return res
	}
}

// HttpGetUser 远端查询用户数据
func HttpGetUser(name string) user {
	// mock 模拟接口查询
	u := user{
		Name:  name,
		Email: "qq.com",
		Phone: "12321312312",
	}
	return u
}

// 查询方法

func queryUser() {
	for i := 0; i < 10; i++ {
		userName := fmt.Sprintf("user_name_%d", i)
		GetUser(userName)
	}
}

func main() {
	//c := cache.New(30*time.Second, 5*time.Second)
	//
	//c.Set("k1", "v1", 31*time.Second)
	//
	//res, ok := c.Get("k1")
	//fmt.Println(res, ok)
	//
	//time.Sleep(time.Second * 32)
	//
	//res, ok = c.Get("k1")
	//fmt.Println(res, ok)

	log.Printf("第1次queryUser")
	queryUser()
	log.Printf("第2次queryUser")
	queryUser()
	time.Sleep(61 * time.Second)
	log.Printf("第3次queryUser")
	queryUser()
}
