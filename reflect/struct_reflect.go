package main

import (
	"fmt"
	"reflect"
)

// 结构体反射

type student struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score" ini:"s_score"`
}

func (s student) Study() string {
	msg := "好好学习，天天向上"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "早睡早起"
	fmt.Println(msg)
	return msg
}

// 根据反射区获取结构中的方法的函数
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod()) // 有多少个方法
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name: %s\n", t.Method(i).Name)
		fmt.Printf("method: %s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args []reflect.Value
		v.Method(i).Call(args) // 调用方法
	}

	// 通过方法名获取结构体的方法
	t.MethodByName("Sleep") // (Method, bool)
	methodObj := v.MethodByName("Sleep") // (Value)
	fmt.Println(methodObj.Type()) // func() string
}

func main() {
	stu1 := student{
		Name:  "无解",
		Score: 90,
	}

	// 通过反射去获取结构体中所有字段信息
	t := reflect.TypeOf(stu1)
	// name: student, kind: struct
	fmt.Printf("name: %v, kind: %v\n", t.Name(), t.Kind())

	// 遍历结构体变量的所有字段
	for i := 0; i < t.NumField(); i++ {
		// i 就是结构体字段的索引
		// 根据结构体字段的索引去取字段
		fieldObj := t.Field(i)
		// name:Name,type:string,tag:json:"name" ini:"s_name"name:Score,type:int,tag:json:"score" ini:"s_score"
		fmt.Printf("name:%v,type:%v,tag:%v", fieldObj.Name, fieldObj.Type, fieldObj.Tag)

		// name s_name
		// score s_score
		fmt.Println(fieldObj.Tag.Get("json"), fieldObj.Tag.Get("ini"))
	}

	// 根据名字去取结构体中的字段
	fieldObj2, ok := t.FieldByName("Score")
	if ok {
		// name:Score,type:int,tag:json:"score" ini:"s_score"
		fmt.Printf("name:%v,type:%v,tag:%v\n", fieldObj2.Name, fieldObj2.Type, fieldObj2.Tag)
	}

	fmt.Println("=============")

	printMethod(stu1)
}
