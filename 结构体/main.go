package main

import "fmt"

type Person1 struct {
	name   string
	age    int
	hobby  [2]string
	num    []int
	parent map[string]string
}

func main() {
	p1 := Person1{
		name:   "二狗子",
		age:    12,
		hobby:  [2]string{"裸奔", "大保健"},                                   // 拷贝了一份
		num:    []int{69, 19, 99, 38},                                    // 未拷贝
		parent: map[string]string{"father": "wujie", "mother": "dwqdqw"}, // 未拷贝
	}

	p2 := p1

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("")

	//p1.hobby[0] = "搓澡"
	//p1.num[0] = 12
	p1.parent["father"] = "我"
	fmt.Println(p1)
	fmt.Println(p2)

}
