package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/**
带过期时间的map 定时清理
map做缓存用
希望缓存存货 5分钟
将加锁的时间控制在最低
*/

type MapCache struct {
	sync.RWMutex
	mp map[string]*item
}

type item struct {
	value int   // 值
	ts    int64 //时间戳 item 被创建出来的时间
}

func (c *MapCache) Get(key string) *item {
	c.RLocker()
	defer c.RUnlock() // 在return结束的时候解锁
	return c.mp[key]
}

func (c *MapCache) Set(key string, value *item) {
	c.Lock()
	defer c.Unlock() // 在return结束的时候解锁
	c.mp[key] = value
}

func (c *MapCache) Clean(timeDelta int64) {
	// 一直运行
	for {
		// 遍历map的中的key value
		// now - ts > 5分钟 5 * 60 认为key value过期

		toDelKeys := make([]string, 0) // 待删除的key
		now := time.Now().Unix()       // 返回int64 时间戳

		// 先加读锁 把所有待删除的k拿到
		c.RLock()
		for k, v := range c.mp {
			if now-v.ts > timeDelta {
				toDelKeys = append(toDelKeys, k)
			}
		}
		c.RUnlock()
		// 进行删除
		// 加写锁 降低加写锁的时间
		c.Lock()
		for _, k := range toDelKeys {
			log.Printf("删除过期数据[key:%s]", k)
			delete(c.mp, k)
		}
		c.Unlock() // 写锁释放

		time.Sleep(2 * time.Second) // 每2秒执行一次清理操作
	}
}

func (c *MapCache) CacheNum() int {
	c.RLock()
	defer c.RUnlock()
	keys := make([]string, 0)
	for k := range c.mp {
		keys = append(keys, k)
	}
	return len(keys)
}

func main() {
	c := MapCache{
		mp: make(map[string]*item),
	}

	// 让清理的任务异步执行
	// 没5秒运行一次，检查时间差大于30秒 item就删除
	go c.Clean(30)

	// 塞入缓存
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i)
		ts := time.Now().Unix()
		im := &item{
			value: i,
			ts:    ts,
		}
		log.Printf("设置缓存[item][key:%s][v:%v]", key, im)
		c.Set(key, im)
	}

	time.Sleep(33 * time.Second)

	// 打印一下
	log.Printf("缓存的数据量:%d", c.CacheNum())

	// 更新缓存
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key_%d", i)
		ts := time.Now().Unix()
		im := &item{
			value: i,
			ts:    ts,
		}
		log.Printf("更新缓存[item][key:%s][v:%v]", key, im)
		c.Set(key, im)
	}

	log.Printf("缓存的数据量:%d", c.CacheNum())
	// pending
	select {}
}
