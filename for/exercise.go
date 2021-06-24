package main

import "fmt"

func main() {
	//guess()
	oneToHundred()
	outerSeven()
}

/**
猜数字
*/
func guess() {
	data := 66
	flag := true
	for flag {
		var userInputNum int
		fmt.Println("请输入一个数")
		fmt.Scanln(&userInputNum)

		if userInputNum > data {
			fmt.Println("太大了")
		} else if userInputNum < data {
			fmt.Println("太小了")
		} else {
			fmt.Println("答对了")
			flag = false
		}
	}
}

func oneToHundred() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

func outerSeven() {
	for i := 0; i < 10; i++ {
		if i != 7 {
			fmt.Println(i)
		}
	}
}

func oldNum() {
	for i := 0; i < 100; i++ {
		if i%2 == 1 {
			fmt.Println(i) // 输出奇数
		}
	}
}

func sum100() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum = sum + i
	}

	fmt.Println(sum)
}

func reverseNum() {
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
}
