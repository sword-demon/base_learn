package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	// 光声明map类型，不初始化，a就是初始化值为nil
	var a map[string]int
	fmt.Println(a == nil) // true

	// map的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil) // false

	// 添加键值对
	a["wujie"] = 100
	a["dwqdqw"] = 200
	fmt.Printf("a:%#v\n", a)
	fmt.Printf("type:%T\n", a)

	// 声明map的同时完成初始化
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("b:%#v\n", b)
	fmt.Printf("type: %T\n", b)

	// 判断某个键存不存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["张三"] = 70
	scoreMap["李四"] = 71

	// 判断王五在不在scoreMap中
	v, ok := scoreMap["王五"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("不存在")
	}

	// map遍历
	// map是无序的，键值对和添加的顺序是无关的
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	// 删除键值对
	delete(scoreMap, "李四")

	//fmt.Println(scoreMap)

	// 添加50个键值对
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100) // 0~99的随机整数
		scoreMap[key] = value
	}

	//for k, v := range scoreMap {
	//	fmt.Println(k, v)
	//}

	// 按照key从小到大的顺序去遍历scoreMap
	// 1. 先取出所有的key 存放到切片中
	keys := make([]string, 0, 100)
	// 2. 对取出的key做排序
	for k := range scoreMap {
		keys = append(keys, k)
	}
	// 对key做排序
	sort.Strings(keys) // keys 目前是有序的切片s
	// 3. 按照排序后的key对scoreMap进行排序
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

	// 元素类型为map的切片
	mapSlice := make([]map[string]int, 8, 8)
	// 还需要完成内部map元素的初始化
	mapSlice[0] = make(map[string]int, 8) // 完成了map的初始化
	mapSlice[0]["沙河"] = 100
	fmt.Println(mapSlice)

	// 值为切片的map
	var sliceMap = make(map[string][]int, 8) // 只完成了map的初始化

	value, ok := sliceMap["中国"]
	if ok {
		fmt.Println(value)
	} else {
		sliceMap["中国"] = make([]int, 8) // 完成了对切片的初始化
		sliceMap["中国"][0] = 100
		sliceMap["中国"][1] = 200
		sliceMap["中国"][2] = 300
	}

	// 遍历sliceMap
	for k, v := range sliceMap {
		fmt.Println(k, v)
	}

	// 统计字符串中每个单词出现的次数
	var s = "how do you do"
	var wordCount = make(map[string]int, 10)
	words := strings.Split(s, " ")
	for _, word := range words {
		wordValue, ok := wordCount[word]
		if ok {
			wordCount[word] = wordValue + 1
		} else {
			wordCount[word] = 1
		}
	}
	for k, v := range wordCount {
		fmt.Println(k, v)
	}
}
