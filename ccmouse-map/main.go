package main

import "fmt"

func map1() {
	m := map[string]string{
		"name":    "wujie",
		"course":  "golang",
		"site":    "wjstar",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // empty map

	var m3 map[string]int // nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting value")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")

	name, ok = m["name"]
	fmt.Println(name, ok)
}

func lengthOfNnRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			// 更新 start
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func map2() {
	fmt.Println(lengthOfNnRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNnRepeatingSubStr("bbbbbbb"))
	fmt.Println(lengthOfNnRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNnRepeatingSubStr(""))
	fmt.Println(lengthOfNnRepeatingSubStr("b"))
	fmt.Println(lengthOfNnRepeatingSubStr("abcdef"))
	fmt.Println(lengthOfNnRepeatingSubStr("这里是中文"))
	fmt.Println(lengthOfNnRepeatingSubStr("一二三二一"))
}

func main() {
	map2()
}
