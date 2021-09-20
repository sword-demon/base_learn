package main

import (
	"fmt"
)

// 方法的定义实例

// Person 是一个结构体
type Person struct {
	name string
	age  int8
}

// NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream 接受者: 定义方法  是属于具体类型的
func (p *Person) Dream() {
	fmt.Printf("%s的梦想是学好go语言!\n", p.name)
}

// SetAge 修改年龄的方法
// 指针类型的接受者
func (p *Person) SetAge(newAge int8) {
	// 函数传参是值拷贝
	p.age = newAge
}

// People 结构体的匿名字段
// 字段必须要求唯一的，相同类型的字段只能有一个
type People struct {
	string
	int
}

// 嵌套结构体

type Peo struct {
	Name   string
	Gender string
	Age    int
	//Address Address // 嵌套另外一个结构体
	Address // 类型名和名称一样可以直接写类型名
	Email   // 嵌套另一个结构体
}

type Address struct {
	Province   string
	City       string
	UpdateTime string
}

type Email struct {
	Addr       string
	UpdateTime string
}

func main() {
	p1 := NewPerson("无解", 18)

	//(*p1).Dream()

	//p1.Dream()

	fmt.Println(p1.age)

	p1.SetAge(int8(20))

	fmt.Println(p1.age)

	p2 := People{
		"无解",
		18,
	}

	// 只能根据字段的类型去访问
	fmt.Println(p2.string, p2.int)

	p3 := Peo{
		Name:   "无解",
		Gender: "男",
		Age:    18,
		Address: Address{
			Province:   "江苏省",
			City:       "苏州市",
			UpdateTime: "2021",
		},
		Email: Email{
			Addr:       "1221@qq.com",
			UpdateTime: "2022",
		},
	}

	fmt.Printf("%#v\n", p3)
	fmt.Println(p3.Name, p3.Gender, p3.Age, p3.Address.Province, p3.Address.City)
	fmt.Println(p3.Province)

	fmt.Println(p3.Address.UpdateTime) // 带有结构体冲突的字段
	fmt.Println(p3.Email.UpdateTime)   // 需要加上你访问的哪个具体的结构体的名称
}
