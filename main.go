// 定义了包的名字，声明当前go文件属于哪个包，
package main

// 告诉go 我需要fmt的这个包，实现了一些格式化输出的函数
import "fmt"

// 定义一个函数 main函数，程序开始执行的函数，每一个可执行的函数必须包含一个main函数
func main() {
	// 这是单行的注释

	/*
	多行注释，块注释
	 */
	fmt.Println("人生苦短，let's go")
}