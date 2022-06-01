package main

import (
	"fmt"
	"net"
)

// tcp server 端
func processConn(conn net.Conn) {

}

func main() {
	// 1，本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:2000 failded , err", err)
		return
	}
	// 2,监听端口，等待连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("accept failed, err:", err)
		return
	}
	// 3，与客户端通信
	var tem [128]byte
	n, err := conn.Read(tem[:])
	if err != nil {
		fmt.Println("read from conn failed, err", err)
		return
	}
	fmt.Println(string(tem[:n]))
}
