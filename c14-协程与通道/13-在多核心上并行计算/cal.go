package main

import "runtime"

const NCPU = 4

/**
DoAll()创建一个sem通道，每个并行计算都将在对其发送完成信号;在一个for循环中NCPU个协程被启动了，
每个协程会承担1/NCPU的工作量。

DoAll()会在for循环中等待NCPU个协程完成：sem通道就像一个信号量
*/
func DoAll() {
	sem := make(chan int, NCPU)
	for i := 0; i < NCPU; i++ {
		go DoPart(sem)
	}
}

/**
每一个DoPart()协程都会向sem通道发送完成信号
 */
func DoPart(sem chan int) {
	sem <- 1
}

func main() {
	runtime.GOMAXPROCS(NCPU)
	DoAll()
}
