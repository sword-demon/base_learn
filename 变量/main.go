package main

import "fmt"

func main() {
	name := "wujie"
	nickname := name

	fmt.Println(name, &name)
	fmt.Println(nickname, &nickname)

	name = "666"

	fmt.Println(name, &name)
	fmt.Println(nickname, &nickname)
}
