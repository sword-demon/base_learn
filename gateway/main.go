package gateway

import (
	"fmt"
	"github.com/xinwangqilin/base_learn/gateway/unpack"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	err = unpack.Encode(conn, "hello world 0 !!!")
	if err != nil {
		return
	}
}
