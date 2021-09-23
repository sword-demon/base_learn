package main

import (
	"fmt"
	"reflect"
)

func typeOf(x interface{}) {
	// 类型断言
	// 反射来知道传进来的参数是啥类型
	v := reflect.TypeOf(x) // 返回类型
	//fmt.Println(v, v.Name(), v.Kind()) // float32 float32 float32
	fmt.Println(v, v.Kind()) // float32 float32 float32
	fmt.Printf("%T\n", v)    // *reflect.rtype
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v, %T\n", v, v)
}

func main() {

	var a float32 = 1.234
	typeOf(a)

	var b int8 = 12
	typeOf(b)

	var c []int
	typeOf(c) // []int slice    .Name() 返回为空 切片和map以及指针类型的变量 .Name() 都返回为空

	reflectValue(a) // 1.234, reflect.Value
}
