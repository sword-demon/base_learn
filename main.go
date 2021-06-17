// 定义了包的名字，声明当前go文件属于哪个包，
package main

import (
	"fmt"
	"github.com/xinwangqilin/base_learn/api"
)

// 定义一个函数 main函数，程序开始执行的函数，每一个可执行的函数必须包含一个main函数
func main() {
	// 这是单行的注释

	/*
	多行注释，块注释
	 */

	fmt.Println("hello world")
	// 调用city.go 的Add方法，因为他们属于同一个包，包名相同，直接去调用即可
	Add()

	// 调用api/baidu.go 的Baidu() 方法 ，只要包.方法名即可
	api.Baidu()
}