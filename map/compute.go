package main

import "time"

// 只声明
//var m1 map[string]string

// 声明又初始化
//var m2 = map[string]string{"a": "b"}

func main() {
	//m := make(map[string]string)
	//
	//m["app"] = "淘宝"
	//m["lang"] = "golang"
	//
	//// 删除
	//delete(m, "app")
	//fmt.Println(m)
	//
	//// 改
	//m["lang"] = "python"
	//
	//// 单变量形式
	//lang := m["lang"]
	//fmt.Println(lang)
	//
	//// 双变量赋值 ok -> boolean 存在就true 不存在就false
	//lang1, ok := m["lang"]
	//if ok {
	//	fmt.Println("lang存在", lang1)
	//} else {
	//	fmt.Println("lang不存在")
	//	m["lang"] = "Java"
	//}
	//
	//newMap := make(map[string]int)
	//keys := make([]string, 0)
	//
	//// 遍历map
	//for i := 0; i < 20; i++ {
	//	key := fmt.Sprintf("key_%d", i)
	//	keys = append(keys, key)
	//	newMap[key] = i
	//}
	//fmt.Println(newMap)
	//
	//for k, v := range newMap {
	//	fmt.Printf("[%s=%d]\n", k, v)
	//}
	//
	//// 切片排序
	//sort.Sort(sort.StringSlice(keys))
	//fmt.Println("切片排序", keys)
	//
	//// map遍历是无序的
	//// 有序key遍历
	//fmt.Println("有序遍历")
	//for _, k := range keys {
	//	fmt.Printf("[%s=%d]\n", k, newMap[k])
	//}
	//
	//// map的key建议不要使用float64类型, value可以是任意类型
	//
	//var doubleM map[string]map[string]string
	//// panic: assignment to entry in nil map
	//doubleM = make(map[string]map[string]string)
	//
	//v1 := make(map[string]string)
	//
	//v1["k1"] = "v1"
	//doubleM["m1"] = v1

	// golang的原生map是线程不安全的
	c := make(map[int]int)
	// 匿名 goroutine 循环写map
	go func() {
		for i := 0; i < 10000; i++ {
			c[i] = i
		}
	}()

	// fatal error: concurrent map read and map write
	// 循环读map
	//go func() {
	//	for i := 0; i < 10000; i++ {
	//		_ = c[i]
	//	}
	//}()

	// fatal error: concurrent map writes
	go func() {
		for i := 0; i < 10000; i++ {
			c[i] = i
		}
	}()

	time.Sleep(30 * time.Minute)
}
