package main

import "fmt"

type errorStr struct {
	s string
}

func (e errorStr) Error() string {
	return e.s
}

func NewError(text string) error {
	return errorStr{text}
}

var ErrType = NewError("EOF")

func main() {
	if ErrType == NewError("EOF") {
		// 还去去比较字符串是否相等，这样其实还是不行的，还是需要去取地址进行比较
		fmt.Println("Error:", ErrType)
	}
}
