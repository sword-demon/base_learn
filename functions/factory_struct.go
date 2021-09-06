package main

import "fmt"

type file struct {
	fd   int
	name string
}

// NewFile 对外提供的一个创建File的方法
// 结构体工厂的方式
func NewFile(fd int, name string) *file {
	return &file{fd, name}
}

func main() {
	f := NewFile(10, "./test.txt")
	fmt.Println(f)
}
