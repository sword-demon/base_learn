package main

import (
	"fmt"
	"os"
)

// 学员信息管理系统

// 1. 添加学员信息
// 2. 编辑学员信息
// 3. 显示所有学员信息

func main() {
	sm := newStudentMgr()
	// 一直执行
	for {
		// 1. 打印系统菜单
		showMenu()
		// 2. 等待用户选择要执行的选项
		var input int
		fmt.Println("请输入你要操作的序号:")
		fmt.Scanf("%d\n", &input)
		fmt.Println("用户输入的是: ", input)
		// 3. 执行用户选择的动作
		switch input {
		case 1:
			// 添加学员
			newStu := getInput()
			sm.addStudent(newStu)
		case 2:
			// 编辑学员
			newStu := getInput()
			sm.editStudent(newStu)
		case 3:
			// 展示所有学员
			sm.showStudent()
		case 4:
			// 退出
			os.Exit(0)
		}
	}
}

func showMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1. 添加学员信息")
	fmt.Println("2. 编辑学员信息")
	fmt.Println("3. 显示所有学员信息")
	fmt.Println("4. 退出系统")
}

// 获取用户输入的学员信息
func getInput() *student {
	fmt.Println("请按要求输入学员信息")
	var (
		id    int
		name  string
		class string
	)
	fmt.Print("请输入学员的学号: ")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学员的姓名: ")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学员的班级: ")
	fmt.Scanf("%s\n", &class)

	// 就能拿到用户输入的学员的信息
	student := newStudent(id, name, class) // 调用构造函数
	return student
}
