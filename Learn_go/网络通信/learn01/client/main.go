package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var exitChan chan bool = make(chan bool, 1) // 声明并初始化了一个名为 exitChan 的通道

func main() {
	fmt.Printf("准备开启客户端…………")

	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
	}

	go proc(conn)
	if <-exitChan { // 如果通道中收到一个true
		fmt.Println("客户端退出连接")
	}
}

// 通过标准流进行数据传递
func proc(conn net.Conn) {
	// 最后执行关闭连接
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin) // 从标准输入读取
	for {
		line, err := reader.ReadString('\n') // 每次读取一行
		if err != nil {
			fmt.Println(err)
			exitChan <- true
			break
		}
		line = strings.Trim(line, "\r\n")
		if line == "bye" {

			exitChan <- true //向通道发送一个true
			break
		}

		len, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println(err)
			exitChan <- true
			break
		}
		fmt.Println("写入的字节数：", len)

	}

	// 获取远程地址

}
