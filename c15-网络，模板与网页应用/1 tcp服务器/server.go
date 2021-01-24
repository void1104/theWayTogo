package main

import (
	"fmt"
	"net"
)

/**
一个（web）服务器应用需要响应众多客户端的并发请求：
Go会为每一个客户端产生一个协程用来处理请求。我们
需要使用net包中网络通信的功能。它包含了处理TCP/IP
以及UDP协议，域名解析等方法。
*/

func main() {
	fmt.Println("Starting the server...")
	// 创建listener,其实现了服务器的基本功能：用来监听和接收来自客户端的请求
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return // 终止程序
	}
	// 监听并接收来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		// 每个请求连接都启动一个go程去处理
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		// 用512字节的缓存data来读取客户端发送的数据
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return // 终止程序
		}
		fmt.Printf("Received data: %v\n", string(buf[:len]))
	}
}
