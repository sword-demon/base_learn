package main

import "fmt"

type Person struct {
	name  string
	age   int
	hobby []string
}

func main() {
	// 1. 按照字段的先后顺序
	var p1 = Person{"wujie", 19, []string{"篮球"}}

	// 2. 关键字
	var p2 = Person{name: "wujie", age: 19, hobby: []string{"饺子"}}

	// 3. 先声明再赋值
	var p3 Person
	p3.name = "wujie"
	p3.age = 19
	p3.hobby = []string{"女人"}

	fmt.Println(p1.name, p1.age, p1.hobby)
	fmt.Println(p2.name, p2.age, p2.hobby)
	fmt.Println(p3.name, p3.age, p3.hobby)
}
