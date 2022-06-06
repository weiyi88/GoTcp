package main

import (
	"fmt"
	"net"
	"protocol"
	"strconv"
	"time"
)

func main() {
	protocol.Test()
	return
	connect, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dail failed:", err)
		return
	}
	defer connect.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello 这是第 ` + strconv.Itoa(i) + "条信息"

		// 调用协议 编码数据
		protocol.Test()
		time.Sleep(time.Second)
		connect.Write([]byte(msg))
	}
}
