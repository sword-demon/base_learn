package main

import "log"

type Person struct {
	Name   string
	Age    int
	Labels map[string]string
}

var p = Person{Name: "xiao"}

func main() {
	p1 := Person{
		Name:   "小贾",
		Age:    10,
		Labels: map[string]string{},
	}

	p2 := Person{"无解", 20, map[string]string{}}

	log.Printf("%+v", p)
	log.Printf("%+v", p1)
	log.Printf("%+v", p2)

	p3 := new(Person) // 指针 内存初始化为零值并返回其指针
	log.Printf("%+v", p3)
}
