package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
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
		time.Sleep(time.Second)
		connect.Write([]byte(msg))
	}
}
