package main

import "fmt"

type xxx interface {
	// 空接口
	// 可以存储任意值
	// 空接口一般不需要提前定义
}

// 接口的嵌套
type animal interface {
	mover
	sayer
}

type mover interface {
	move()
}

type sayer interface {
	say()
}

type person struct {
	name string
	age  int
}

// 使用值接受者实现接口：类型的值和类型的指针都能保存到接口变量中
//func (p person) move() {
//	fmt.Printf("%s在炮...\n", p.name)
//}

// 使用指针接收者实现接口 只有类型指针能够保存到接口变量中
func (p *person) move() {
	fmt.Printf("%s在跑...\n", p.name)
}

func (p *person) say() {
	fmt.Printf("%s在叫...\n", p.name)
}

// 空接口的应用
// 1. 作为函数的参数
// 2. 空接口的类型可以作为map的value


func main() {
	var m mover
	//p1 := person{
	//	name: "无解",
	//	age:  12,
	//}
	p2 := &person{
		name: "带我去多无群",
		age:  18,
	}
	//m = p1 // 无法保存，因为p1是值类型，没有实现mover接口
	m = p2
	m.move()
	fmt.Println(m)

	// 定义一个空接口变量x x可以存储任意类型
	var x interface{}

	x = "hello"
	x = 100
	x=  false
	fmt.Println(x)

	ret, isTrue := x.(bool)
	if isTrue {
		fmt.Println(ret)
	} else {
		fmt.Println("false")
	}
}
