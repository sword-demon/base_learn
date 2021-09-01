/*
 * @Author: wxvirus
 * @Date: 2021-09-01 20:45:11
 * @LastEditTime: 2021-09-01 21:14:26
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /base_learn/type/main.go
 */
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	// 读第二次，啥也读不到，也不会报错
	body, err := io.ReadAll(r.Body) // io.ReadCloser io操作设计成只能读一次
	if err != nil {
		fmt.Fprintf(w, "read body failed : %v", err)
		return
	}
	// 类型转换将 []byte 转换成string
	fmt.Fprintf(w, "read the data: %s \n", string(body))
}

func getBodyIsNil(w http.ResponseWriter, r *http.Request) {
	// GetBody 原则上是可以多次读取，但是在原生的http.Request里，这个是个nil
	if r.GetBody == nil {
		fmt.Fprint(w, "GetBody is nil\n")
	} else {
		fmt.Fprintf(w, "GetBody not nil\n")
	}
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Fprintf(w, "query is %v\n", values)
}

func wholeUrl(w http.ResponseWriter, r *http.Request) {
	// 唯一肯定会有的值  Path 其他数值可能都会没有
	data, _ := json.Marshal(r.URL) // json序列化
	fmt.Fprintf(w, string(data))
}

func getHeader(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v\n", r.Header) // %v 显示结构体
}

func form(w http.ResponseWriter, r *http.Request) {
	// 啥都没有
	fmt.Fprintf(w, "before parse form %v\n", r.Form)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parse form error %v\n", r.Form)
	}
	// 才会有
	fmt.Fprintf(w, "after parse form %v\n", r.Form)
}

func main() {
	fmt.Println("http库的学习")
}
