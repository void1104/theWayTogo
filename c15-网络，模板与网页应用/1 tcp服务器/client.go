package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/**
在网络编程中net.Dial非常重要，其简洁地抽象了网络层和传输层。
所以不管是IPv4还是IPv6，TCP或UDP都可以使用这个公共接口。
 */

func main(){
	// 打开连接
	conn, err := net.Dial("tcp","localhost:50000")
	if err != nil {
		// 由于目标计算机拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName,"\r\n")	// Windows平台下用 "\r\n"
	// 给服务器发送信息直到程序退出
	for{
		fmt.Println("What to send to the server?Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input,"\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
	}
}