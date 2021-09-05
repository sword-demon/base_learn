package main

import (
	"errors"
	"fmt"
)

type errorString string

func (e errorString) Error() string {
	return string(e)
}

func New(text string) error {
	return errorString(text)
}

var ErrNameType = New("EOF")
var ErrStructType = errors.New("EOF")

func main() {
	if ErrNameType == New("EOF") {
		// 不鼓励使用 两个字符串相等的形式
		fmt.Println("Named Type Error")
	}

	if ErrStructType == errors.New("EOF") {
		// 条件不满足 不会返回
		fmt.Println("Struct Type Error")
	}
}
