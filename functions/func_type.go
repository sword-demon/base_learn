package main

import (
	"fmt"
)

// 拆分
type f1 func(arg int) (int, bool)

func add99(arg int) (int, bool) {
	return arg + 99, true
}

func feature(data int, exec f1) int {
	data, flag := exec(data)
	if flag {
		return data
	} else {
		return 9999
	}
}

// 变长参数只能放在最后，且只能有一个
func do(num ...int) int {
	fmt.Println(num) // 切片 [1,2 ,3 ,4 ,5]
	sum := 0
	for _, value := range num {
		sum += value
	}
	return sum
}

func F1(n1 int, n2 int) func(int) string {
	return func(n1 int) string {
		fmt.Println("匿名函数")
		return "匿名函数"
	}
}

func main() {
	result := feature(123, add99)
	fmt.Println(result)

	r1 := do(1, 2, 3, 4, 5)
	r2 := do(1, 2, 3, 4, 5, 6)

	fmt.Println(r1, r2)

	// 匿名函数
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}

	// 执行匿名函数
	fmt.Println(res(1, 2))

	// 存储着5个函数
	var functionList []func()

	for i := 0; i < 5; i++ {
		// 没执行前，只是创建了函数，加到了切片里去
		function := func(arg int) func() {
			return func() {
				fmt.Println(arg + i)
			}
		}(i)
		functionList = append(functionList, function)
	}

	// 运行函数前，循环已经执行完，i已经等于5了
	functionList[0]() // 5 5
	functionList[1]() // 5 6
	functionList[2]() // 5 7
}
