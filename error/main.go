/*
 * @Author: your name
 * @Date: 2021-08-31 19:00:54
 * @LastEditTime: 2021-08-31 19:27:09
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /base_learn/error/main.go
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// type errWriter struct {
// 	io.Writer
// 	err error
// }

// func (e *errWriter) OverWriter(buf []byte) (int, error) {
// 	if e.err != nil {
// 		return 0, e.err
// 	}
// 	var n int
// 	n, e.err = e.Writer.Write(buf)
// 	return n, nil
// }

// func WriteResponseOver(w io.Writer, st Status, headers []Header, body io.Reader) error {
// 	ew := &errWriter{OverWriter: w}
// 	fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)

// 	for _, h := range headers {
// 		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
// 	}

// 	fmt.Fprint(ew, "\r\n")
// 	io.Copy(ew, body)

// 	return ew.err
// }

// func AuthenticateRequest(r *Request) error {
// 	return authenticate(r.User)
// }

// CountLines 读取内容的行数
func CountLines(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)
	for {
		_, err = br.ReadString('\n')
		// 先加加
		lines++
		if err != nil {
			// 读完了，中途报错了
			break
		}
	}
	// 统计结果是否是无效的
	if err != io.EOF {
		return 0, err
	}
	return lines, nil
}

// CountLinesSuper 改进
func CountLinesSuper(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	lines := 0

	// 一行一行的读，内部报错，就返回false
	for sc.Scan() {
		lines++
	}

	// 一定要返回sc的内部的err
	return lines, sc.Err()
}

func testPanic() {
	// 延迟执行
	defer func() {
		// data 是下面的panic传递的东西 => Boom
		if data := recover(); data != nil {
			fmt.Printf("hello, panic: %v\n", data)
		}
		fmt.Println("恢复之后从这里继续执行")
	}()

	panic("Boom")
	fmt.Println("这里肯定不会执行")
}

func main() {
	path := "./"
	f, err := os.Open(path)
	if err != nil {
		// handle error
	}
	fmt.Println(f)

	// do stuff

	testPanic()
}
