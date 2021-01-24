package rpc_object

/**
Go程序之间可以使用net/rpc实现相互通信，这是另一种客户端-服务器
应用场景。它提供了一种方便的途径，通过网络连接调用远程函数。
仅当程序运行在不同机器上时，这项技术才实用。
rpc包建立在gob包之上，实现了自动编码/解码传输的跨网络方法调用
*/

/**
服务器端需要注册一个对象实例，与其类型名一起，使之成为一项可见的服务：
它允许远程客户端跨越网络或其他I/O连接访问对象已导出的方法。总之就是
在网络上暴露类型的方法。
*/

type Args struct {
	N, M int
}

func (t *Args) Multiply(args *Args, reply *int) error {
	*reply = args.N * args.M
	return nil
}
