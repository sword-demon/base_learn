package main

import "fmt"

type Course struct {
	name  string
	price int
	url   string
}

func main() {
	// go语言不支持面向对象

	var c = Course{
		name:  "django",
		price: 100,
		url:   "https://www.immoc.com",
	}

	fmt.Println(c.name, c.url, c.price)
}
