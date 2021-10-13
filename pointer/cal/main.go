package main

import "fmt"

func main() {
	var intVariables int = 100

	fmt.Printf("intVariables的值=%d,地址=%v\n", intVariables, &intVariables)

	var pointerVariables *int = &intVariables
	fmt.Printf("pointerVariables=%d,地址=%v\n", pointerVariables, &pointerVariables)

	// 修改变量的值
	*pointerVariables = 200

	fmt.Println(*pointerVariables)
}
