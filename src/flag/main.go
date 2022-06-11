package main

import (
	"flag"
	"fmt"
	"os"
)

// 获取命令行参数
//os.Args[ 运行脚本本身，第一个参数，第二个参数....]
func OsArgs() {
	fmt.Printf("%s\n", os.Args)
	for k, v := range os.Args {
		fmt.Println(k, "===>", v)
	}
}

// flag 获取命令行参数
// 可以 -标志符   --标志符  ， = 或直接跟值，  获取命令行参数
//go run main.go -name Jack --age=99
func Flag() {
	// 创建一个标志位参数
	name := flag.String("name", "Aring", "请输入名字")
	age := flag.Int("age", 23, "请输入真实年龄")
	// 使用flag
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
}

func main() {
	//OsArgs()
	Flag()
}
