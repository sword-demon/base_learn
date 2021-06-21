package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Println("请输入用户名:")

	//count, err := fmt.Scan(&name) // 把name变量的内存地址放进来，给内存地址指向的空间赋值的过程
	_, err := fmt.Scan(&name) // 使用下划线来代替不使用的变量
	if err == nil {           // 错误信息为空，代表没错，这里容易被单词误解
		fmt.Println(name)
	} else {
		fmt.Println("用户输入错误", err)
	}

	ln()
}

func ln() {
	var name string
	_, err := fmt.Scanln(&name)
	fmt.Println(err)
	fmt.Println(name)
}

func f() {
	var name string
	fmt.Println("请输入用户名")
	fmt.Scanf("我叫%s", &name)
	fmt.Println(name)
}

func solution() {
	reader := bufio.NewReader(os.Stdin)
	line, isPrefix, err := reader.ReadLine() // 读一行
	fmt.Println(line, isPrefix, err)
}
