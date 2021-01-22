package main

/**
使用带缓冲区的通道可以实现，缓冲区容量就是同时处理请求的最大数量
	- 超过MAXREQS的请求将不会被同时处理，因为当信号通道表示缓冲区已满时,handle函数会阻塞且不再处理
      其他请求，直到某个请求从sem中被移除。sem就像一个信号量，这一专业术语用于在程序中表示特定条件的标志变量。
*/

const MAXREQS = 50

var sem = make(chan int, MAXREQS)

type Request struct {
	a, b   int
	replyc chan int
}

func process(r *Request) {
	// do something
}

// 通过这种方式，应用程序可以通过使用缓冲通道（通道被用作信号量）使协程同步其对该资源的使用，从而充分利用有限的资源（如内存）
func handler(r *Request) {
	sem <- 1 // 当sem已满时，就会阻塞在这里，直到sem有空位为止
	process(r)
	<-sem
}

func server(service chan *Request) {
	for {
		request := <-service
		go handler(request)
	}
}

func main() {
	service := make(chan *Request)
	go server(service)
}
