package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 1, 与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial 127.0.0.1:20000 fail error", err)
		return
	}
	//2，发送数据
	//var msg string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请你说话：")
		msg, _ := reader.ReadString('\n') // 读到换行字符
		msg = strings.TrimSpace(msg)      // 缩减空格
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}
