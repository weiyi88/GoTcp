package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})

	if err != nil {
		fmt.Println("listen failed ,err:", err)
		return
	}
	defer conn.Close()

	// 不需要建立连接，直接发送数据
	var data [1024]byte

	for {
		n, addr, err := conn.ReadFromUDP(data[:]) //	读到的字节，对方地址
		if err != nil {
			fmt.Println("read from UDP failed,err", err)
			return
		}
		fmt.Println(data[:])
		reply := strings.ToUpper(string(data[:n])) // 全部大写
		// 发送数据
		conn.WriteToUDP([]byte(reply), addr)

	}
}
