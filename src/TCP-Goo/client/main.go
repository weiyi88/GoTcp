package main

import (
	"fmt"
	"net"
	"package/protocol"
	"strconv"
)

func main() {
	connect, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dail failed:", err)
		return
	}
	defer connect.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello 这是第 ` + strconv.Itoa(i) + "条信息"
		// 调用协议 编码数据
		connect.Write([]byte(msg))
		b, err := protocol.Encode(msg)
		if err != nil {
			fmt.Println("encode failed ,err:", err)
			return
		}
		connect.Write(b)
	}
}
