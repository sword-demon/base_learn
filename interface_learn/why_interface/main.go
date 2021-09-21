package main

import "fmt"

type dog struct {
}

type cat struct {
}

func (d *dog) say() {
	fmt.Println("汪汪汪")
}

func (c *cat) say() {
	fmt.Println("喵喵喵")
}

// 接口不管你是什么类型，它只管你要实现什么方法
// 定义一个类型，一个抽象的类型，只要实现了say方法 这个方法的类型都可以称之为sayer类型
type sayer interface {
	say()
}

type person struct {
	name string
}

func (p *person) say() {
	fmt.Println("啊啊啊啊啊")
}

func click(arg sayer) {
	arg.say() // 不管传进来的是什么，都要调用say方法
}

func main() {
	d1 := &dog{}
	click(d1)

	c1 := &cat{}
	click(c1)

	p1 := &person{
		name: "无解",
	}
	click(p1)
}
