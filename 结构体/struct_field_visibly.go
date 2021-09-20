package main

import (
	"encoding/json"
	"fmt"
)

// 结构体字段的可见性
// 如果go语言中包里定义的标识符的首字母是大写的，那就是对外可见的

// JSON序列化

// student 学生
type student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

// NewStudent student的构造函数
func NewStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}

// class 班级
type class struct {
	Title    string    `json:"title"`
	Students []student `json:"student_list" db:"student" xml:"ss"`
}

func main() {
	c1 := class{
		Title:    "火箭101",
		Students: make([]student, 0, 20),
	}

	// 往班级c1中添加学生
	for i := 0; i < 10; i++ {
		// 造10个学生
		tmpStu := NewStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}

	fmt.Printf("%#v\n", c1)

	// JSON序列化，Go语言的数据 -> JSON格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("JSON 序列化失败", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)

	fmt.Println("==================")

	// JSON反序列化，将满足JSON格式的字符串转变成go语言中的数据
	var c2 class
	err = json.Unmarshal([]byte(data), &c2) // 修改c2的值，传一个指针
	if err != nil {
		fmt.Println("json unmarshal failed, ", err)
		return
	}
	fmt.Printf("%#v\n", c2)
}
