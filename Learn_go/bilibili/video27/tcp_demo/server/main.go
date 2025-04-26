package main

import (
	"bufio"
	"fmt"
	"net"
)

//tcp server demo

func process(conn net.Conn) {
	defer conn.Close() // 处理完之后要关闭这个链接
	// 针对当前的链接做数据的发送和接收操作
	reader := bufio.NewReader(conn)
	for {
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from failed,err:%v\n", err)
			break
		}
		recv := string(buf[:n])
		fmt.Println("接收到的数据：", recv)
		if _, err := conn.Write([]byte("ok")); err != nil {
			fmt.Printf("write failed: %v\n", err) //把收到的数据返回给客户端
			break
		}
	}
}
func main() {
	// 1.开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed,err:%v\n", err)
		return
	}
	for {
		// 等待客户端来建立链接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed,err:%v\n", err)
			continue
		}
		// 3.启动一个单独的goroutine去处理链接
		go process(conn)
	}
}
