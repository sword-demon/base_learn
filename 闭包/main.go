package main

import (
	"fmt"
	"time"
)

// 匿名函数 + 定义它的上下文
// 可以访问定义以外的变量
// 闭包里面使用闭包外面的参数，其值是最终调用的时候确定下来的

func main() {
	i := 13
	a := func() {
		fmt.Printf("i is %d\n", i)
	}
	a()

	fmt.Println(ReturnClosure("Tome")())

	Delay()

	time.Sleep(time.Second)
}

func ReturnClosure(name string) func() string {
	return func() string {
		return "Hello" + name
	}
}

func Delay() {
	fns := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		// i最后=9，但是还要经历一个 +1 => i = 10
		fns = append(fns, func() {
			fmt.Printf("hello this is %d\n", i)
		})
	}

	for _, fn := range fns {
		fn()
	}
}
