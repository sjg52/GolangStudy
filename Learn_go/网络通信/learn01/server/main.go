package main

import (
	"fmt"
	"net"
)

// 服务器端
func main() {
	fmt.Printf("准备开启服务器…………")
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()
	//阻塞去监听listener
	for {
		conn, err := listener.Accept()

		go proc(conn)

		if err != nil {
			fmt.Println(err)
			return
		} else {
			// 获取远程地址
			fmt.Println(conn.LocalAddr())
		}
	}
}

func proc(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf) // 返回接收的字节数
		if err != nil {
			fmt.Println(err)
			return
		}
		// 服务端去读取客户端发送的信息
		fmt.Println(n)
		fmt.Println(string(buf[:n]))
	}
}
