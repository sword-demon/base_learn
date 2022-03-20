package test

import (
	"fmt"
	"testing"
)

// 执行在main函数之前 多个 init 则执行顺序为自上而下
func init() {
	fmt.Println("1白日依山尽，花钱似海流；囊中无一物，急的要跳楼")
}

func init() {
	fmt.Println("2白日依山尽，花钱似海流；囊中无一物，急的要跳楼")
}

func TestFirstTry(t *testing.T) {
	t.Log("my first try")
}

// 斐波那契数列
func TestFibList(t *testing.T) {
	var (
		a = 1
		b = 1
	)
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

// 交换两个变量的值
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	//tmp := a
	//a = b
	//b = tmp

	// 简洁的方式
	a, b = b, a
	t.Log(a, b)
}
