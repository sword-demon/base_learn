package main

import "fmt"

// 切片
// 存储不定长的类型的元素

func slice1() {
	// 变长
	// 声明变量但未初始化 尚未申请内存空间

	var s1 []int
	var s2 []bool
	var s3 []string
	//var s4 []int

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Printf("%#v \n", s1) // nil
	fmt.Printf("%#v \n", s2)
	fmt.Printf("%#v \n", s3)

	// 切片是不能相互比较的
	//fmt.Println(s1 == s4)

	fmt.Println(len(s1)) // 求切片的长度 元素的个数
	fmt.Println(cap(s1)) // 求切片的容量 可以存放的元素的个数(底层数组能存储的元素)

	// 数组
	var a1 [3]int
	fmt.Printf("%#v \n", a1)

	var company []string // 元素代表员工
	// 租了 10个工位 代表切片的 cap
	// 1个老板 + 4个员工 = 5 长度 目前元素的数量
	// 增加了5个员工 切片中又放了5个元素 len = cap = 10
	// 又增加了8个人 切片中又增加8个元素
	// 公司换租了一个更大的办公地址 -> 切片底层的数组换了
	fmt.Println(cap(company))
}

func slice2() {
	// 切片表达式
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // 切片类型  a[low:high] 左取右不取
	// 容量从 最左边开始往后一直到数组的末尾就是切片的容量
	fmt.Printf("s:%v len(s):%v cap(s):%v \n", s, len(s), cap(s))

	s2 := a[3:]
	// 从索引为3开始切到最后
	fmt.Printf("len(s2): %v, cap(s2): %v \n", len(s2), cap(s2))

	s3 := a[:4]
	// 从索引0开始切，切到索引为4的，但不包含索引为4的
	fmt.Printf("len(s3): %v, cap(s3): %v \n", len(s3), cap(s3))

	s4 := a[:]
	// 从头切到尾
	fmt.Printf("len(s4): %v, cap(s4): %v \n", len(s4), cap(s4))

	fmt.Println("a == s4 ? 不等于 a就是数组类型，s4是切片类型")

	// 等价
	fmt.Println(a[2:])
	fmt.Println(a[2:len(a)])

	fmt.Println(a[:5])

	// 数组，字符串支持切片表达式

	ss := "hello"
	s6 := ss[:3]
	fmt.Println(s6)
	// 字符串切片之后还是字符串 所以没有容量
	fmt.Printf("%T \n", s6)

	// 切片再切片
	a1 := a[0:1]                      // [1]
	fmt.Println(a1, len(a1), cap(a1)) // [1] 1 5

	a2 := a1[:5] // 0 <= low < high <= cap(a1)
	fmt.Println(a2, len(a2), cap(a2))

	a3 := a1[4:5]
	fmt.Println(a3, len(a3), cap(a3))

	a[4] = 500 // 修改 数组中的值 会影响切片
	fmt.Println(a2)
	fmt.Println(a3)

}

func slice3() {
	a := []int{1, 2, 3, 4, 5}
	// 默认切片的容量时从切片的开始索引到数组的最后
	// max 影响切片的容量 它指定了最大能切到的索引位置
	s1 := a[0:1:2] // 0 <= low < high <= max(切到的最大索引)
	// 最终切片容量就是 max - low
	fmt.Println(s1, len(s1), cap(s1))
}

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func slice4() {
	fmt.Println("creating slice")
	var s []int // 初始化 Zero value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)

	printSlice(s2)
	printSlice(s3)
}

func slice5() {
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)
	// 拷贝切片
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("deleting element from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	// 删除头尾
	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from tail")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)
}

func main() {
	slice5()
}
