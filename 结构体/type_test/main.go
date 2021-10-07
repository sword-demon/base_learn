package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

// 定义结构体

type Course struct {
	name string
	price decimal.Decimal
}

// 定义接口

type Callable interface {

}

// 定义函数别名

type handle func(str string)

func main() {
	// 给一个类型定义别名
	type myByte = byte
	var b myByte
	fmt.Printf("%T\n", b)

	// 基于一个已有的类型定义一个新的类型
	type myInt int
	var i myInt
	fmt.Printf("%T\n", i)



}
