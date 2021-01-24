package main

import (
	"./rpc_object"
	"fmt"
	"log"
	"net/rpc"
)

const serverAddress = "localhost"

func main() {
	// 执行rpc.DialHTTP()连接到服务器后，就可以用client.Call()调用远程对象的方法
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	// Synchronous call
	args := &rpc_object.Args{N: 7, M: 8}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
}
