package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//guess()
	//oneToHundred()
	//outerSeven()
	//tag()

	var v1 int16 = 129
	v2 := int8(v1)

	fmt.Println(v2)

	v3 := 19

	result := strconv.Itoa(v3)                  // 只能转换int类型 int8/16 都是不允许的
	fmt.Println(result, reflect.TypeOf(result)) // 整型转换为字符串

	v4 := "666"
	// 内置了错误处理机制
	str4, err := strconv.Atoi(v4) // 第一个值是转换之后的结果，第二个结果就是是否转换成功
	if err == nil {
		fmt.Println("转换成功")
	} else {
		fmt.Println("转换失败")
	}
	fmt.Println(str4, reflect.TypeOf(str4))

	r1 := strconv.FormatInt(int64(v3), 2)	// 10进制转换成二进制，后面填几就转换成几进制
	fmt.Println(r1, reflect.TypeOf(r1))
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

func tag() {
f1:
	for i := 0; i < 3; i++ {
		for j := 1; j < 5; j++ {
			if j == 3 {
				continue f1
			}
			fmt.Println(i, j)
		}
	}
}
