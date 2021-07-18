package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	v1 = iota
	v2
	v3
	v4
)

func main() {
	name := "wujie"
	nickname := name

	fmt.Println(name, &name)
	fmt.Println(nickname, &nickname)

	name = "无佩奇"

	fmt.Println(name, &name)
	fmt.Println(nickname, &nickname)

	// 九九乘法表
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println("")
	}

	fmt.Println(v1, v2, v3, v4)

	username := "无解"

	// 获取字符串的长度，字节长度
	fmt.Println(len(username))

	// 字符串转换为一个字节集合
	byteStr := []byte(username)
	// [230 151 160 232 167 163] 以10进制的方式进行展示
	fmt.Println(byteStr)

	// 字节的集合转换为字符串
	byteList := []byte{230, 151, 160, 232, 167, 163}
	targetString := string(byteList)
	fmt.Println(targetString)

	// rune的集合
	// 将一个字符串转换为unicode字符集码点的集合
	tempSet := []rune(name)
	fmt.Println(tempSet) // [26080 20329 22855]

	fmt.Println(tempSet[0], strconv.FormatInt(int64(tempSet[0]), 16))
	fmt.Println(tempSet[1], strconv.FormatInt(int64(tempSet[1]), 16))
	fmt.Println(tempSet[2], strconv.FormatInt(int64(tempSet[2]), 16))

	// 长度的处理
	runeLength := utf8.RuneCountInString(username)
	fmt.Println(runeLength) // 2

	// 判断字符串是否是以无开头
	result := strings.HasPrefix(username, "无")
	fmt.Println(result)

	var builder strings.Builder
	builder.WriteString("哈哈哈哈")
	builder.WriteString("去你的吧")
	value := builder.String()
	fmt.Println(value)

}
