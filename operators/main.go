package main

import "fmt"

// go中的运算符
func main() {
	// 1. 算数运算符
	a := 10
	b := 20
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a++
	fmt.Println(a)
	a--
	fmt.Println(a)

	// 2. 关系运算符
	fmt.Println(10 > 2)
	fmt.Println(10 != 2)
	fmt.Println(4 <= 5)

	// 3. 逻辑运算符
	fmt.Println(10 > 5 && 1 == 1) // true && true => true

	fmt.Println(!(10 > 5))       // false
	fmt.Println(1 > 5 || 1 == 1) // false || true => true 一真必真

	// 4. 位运算符
	fmt.Println(2 << 2)
	fmt.Println(1 & 5) // 001 101 => 001 1 都为1才为1
	fmt.Println(1 | 5) // 001 101 => 101 5 有一个为1就为1
	fmt.Println(1 ^ 5) // 001 101 => 100 4 每个位上的都不一样

	fmt.Println(1 << 2) // 001 => 100 	4
	fmt.Println(4 >> 2) // 100 => 1 	1

	fmt.Println(1 << 10) // 1024

	// 5. 赋值运算符
	var d int
	d = 5
	d += 5 // d = d + 5
	fmt.Println(d)
}
