package main

import "fmt"

type MyError struct {
	Msg  string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d:%s", e.File, e.Line, e.Msg)
}

func test() error {
	return &MyError{"Something happened", "my_error.go", 42}
}

func main() {
	err := test()

	// 类型断言
	switch err := err.(type) { // 尽量避免使用 error types
	case nil:
	// call success, nothing to do
	case *MyError:
		fmt.Println("error occurred on line:", err.Line)
	default:
		// unknown error
	}
}
