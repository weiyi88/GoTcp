package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed , err:", err)
		}
		fmt.Println("收到信息：", string(buf[:n]))
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
