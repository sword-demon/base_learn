package main

import "fmt"

// 结构体的继承

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动~\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal // 匿名嵌套一个结构体的指针
}

func (d *Dog) wolf() {
	fmt.Printf("%s会叫\n", d.name)
}

func main() {
	dog1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "乐乐",
		},
	}

	dog1.wolf() // 狗会叫
	dog1.move() // 狗会移动
}
