package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Printf("准备开启客户端…………")

	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
	}

	go proc(conn)
}

// 通过标准流进行数据传递
func proc(conn net.Conn) {
	// 最后执行关闭连接
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)  // 从标准输入读取
	line, err := reader.ReadString('\n') // 读取到换行符
	if err != nil {
		fmt.Println(err)
	}
	len, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println(err)
	}

	// 获取远程地址
	fmt.Println("连接成功", conn.RemoteAddr().String(), "写入的字节数：", len)
}
