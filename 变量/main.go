package main

import "fmt"

const (
	v1 = iota
	v2
	v3
	v4
)

func main() {
	name := "wujie"
	nickname := name

	fmt.Println(name, &name)
	fmt.Println(nickname, &nickname)

	name = "666"

	fmt.Println(name, &name)
	fmt.Println(nickname, &nickname)

	// 九九乘法表
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println("")
	}

	fmt.Println(v1, v2, v3, v4)
}
