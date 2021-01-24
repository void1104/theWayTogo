package main

import (
	"./rpc_object"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func main() {
	// 服务器端产生一个rpc_objects.Args类型的对象calc
	calc := new(rpc_object.Args)
	// 并用rpc.Register(object)注册
	// 也可以按名称注册对象，例如：rpc.RegisterName("Calculator", calc)
	rpc.Register(calc)
	// 调用HandlerHTTP()
	rpc.HandleHTTP()
	// 用net.Listen在指定的地址上启动监听
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	// 以协程启动http.Serve(listener, nil)，会为每一个进入listener的HTTP连接创建新的服务线程
	go http.Serve(listener, nil)
	// 使服务器在一段时间内保持运行状态（保证go程走完）
	time.Sleep(1000e9)
}
