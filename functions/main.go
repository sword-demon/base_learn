package main

import (
	"fmt"
	"strings"
)

type Writer interface {
	Write([]byte) (int, error)
}

// Greeting 打招呼 函数 参数类型
func Greeting(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}

// WriteGreeting 将问候信息写入文件
//func WriteGreeting(name string, f *os.File) error {
//	greeting := Greeting(name)
//	_, err := f.Write([]byte(greeting))
//	if err != nil {
//		return err
//	}
//	return nil
//}

func WriteGreeting(name string, w Writer) error {
	greeting := Greeting(name)
	_, err := w.Write([]byte(greeting))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var sb strings.Builder
	err := WriteGreeting("Jon", &sb)
	if err != nil {
		return
	}
}