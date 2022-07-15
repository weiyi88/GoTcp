package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"package/protocol"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		recvStr, err := protocol.Decode(reader)
		if err == io.EOF {
			return // 结尾就退出
		}
		if err != nil {
			fmt.Println("decode faild ,err:", err)
			return
		}
		fmt.Println("收到信息：", string(recvStr))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen faild :", err)
		return
	}
	defer listen.Close()

	for {
		connect, err := listen.Accept() // 持续接收
		if err != nil {
			fmt.Println("accept failed:", err)
			continue
		}
		go process(connect)
	}
}
