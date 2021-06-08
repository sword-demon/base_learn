package main

import "fmt"

// for 循环
func main() {
	// 1. 基本for循环
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	// 2. 省略初始语句，必须保留初始语句后面的分号
	//var i = 0
	//for ; i < 10; i++ {
	//	fmt.Println(i)
	//}

	// 3. 省略初始语句和结束语句
	//var i = 10
	//for i > 0 {
	//	fmt.Println(i)
	//	i--
	//}

	// 4. 死循环
	//for {
	//	fmt.Println("hello world")
	//}

	// 5. break跳出循环
	for i := 0; i < 5; i++ {

		if i == 3 {
			//break
			continue // 继续下一次循环
		}
		fmt.Println(i)
	}
}
