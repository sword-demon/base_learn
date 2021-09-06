package main

import "fmt"

// MyInt 声明类型
type MyInt int

// DoSomething 指针方法 只能特定对象能调用
// 可以是指针，也可以是类型
// 不使用对象，可以使用 _ 代替
func (i *MyInt) DoSomething(a1, a2 int) int {
	return a1 + a2 + int(*i) // 根据指针取到它的值 => 转换成int
}

func Do(a1, a2 int) int {
	return a1 + a2
}

type Base struct {
	name string
}

type Son struct {
	Base
	age int
}

func (b *Base) m1() int {
	return 666
}

func (s *Son) m2() int {
	return 999
}

func main() {
	var v1 MyInt = 1

	result := v1.DoSomething(1, 2)
	fmt.Println(result)

	son := Son{age: 18, Base: Base{
		name: "wujie",
	}}

	son.m1()
	son.m2()
}
