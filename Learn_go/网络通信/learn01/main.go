package main

import (
	"fmt"
	"net"
)

// main starts a TCP server on localhost:8080 and continuously accepts and logs incoming connections.
func main() {
	fmt.Printf("准备开启服务器…………")
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	//记得关闭
	defer listener.Close()
	//阻塞去监听listener
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(conn)
		}
	}
}
