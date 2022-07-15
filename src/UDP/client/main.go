package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务端失败,err:", err)
		return
	}
	defer socket.Close()

	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入内容:")
		msg, _ := reader.ReadString('\n') //直到读到\n结束
		socket.Write([]byte(msg))
		// 收到回复的数据
		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("redv reply msg failed ,err:", err)
			return
		}
		fmt.Println("收到回复信息:", string(reply[:n]))
	}
}
