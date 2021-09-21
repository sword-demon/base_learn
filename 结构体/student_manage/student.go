package main

import "fmt"

type student struct {
	id    int // 学号是唯一的
	name  string
	class string
}

// 学员管理的类型
type studentMgr struct {
	allStudents []*student
}

func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

func newStudent(id int, name, class string) *student {
	// 返回的是一个指针 需要取址符
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

// 1. 添加学生
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

// 2. 编辑学生
func (s *studentMgr) editStudent(newStu *student) {
	for i, v := range s.allStudents {
		// 当学号相同时，就表示找到了要修改的学生
		if newStu.id == v.id {
			// 根据切片的索引直接把旧学生修改为新学生
			s.allStudents[i] = newStu
			return
		}
	}
	// 说明输入的学生没有找到
	fmt.Printf("未找到学号为%d的学生", newStu.id)
}

// 3. 展示所有学生
func (s *studentMgr) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号: %d 姓名: %s 班级: %s\n", v.id, v.name, v.class)
	}
}
