package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

func main() {
	m := sync.Map{}

	// 新增
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i)
		m.Store(key, i)
	}

	// 删除
	m.Delete("key_8")

	// 改
	m.Store("key_1", 999)

	// 查询
	res, ok := m.Load("key_2")
	if ok {
		fmt.Printf("%d数字类型", res.(int)) // 类型断言
	}

	// 遍历
	m.Range(func(key, value interface{}) bool {
		k := key.(string)
		v := value.(int)
		if strings.HasSuffix(k, "3") {
			log.Printf("不想要3")
			return false
		} else {
			log.Printf("[sync.map.Range][遍历][key=%s][v=%d]", k, v)
			return true
		}
	})

	// 先获取值再删掉
	s1, ok := m.LoadAndDelete("key_7")
	log.Printf("key_7 LoadAndDelete : %v", s1)
	s2, ok := m.Load("key_7")
	log.Printf("key_7 LoadAndDelete : %v", s2)

	actual, ok := m.LoadOrStore("key_8", 158)
	if ok {
		log.Printf("原来的值是%v\n", actual)
	} else {
		log.Printf("原来没有，实际是%v\n", actual)
	}
}
