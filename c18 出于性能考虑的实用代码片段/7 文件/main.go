package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 如何打开一个文件并读取：
	fmt.Println(os.Getwd()) // 获取当前路径
	file, err := os.Open("./c18 出于性能考虑的实用代码片段/7 文件/pjx.txt")
	if err != nil {
		fmt.Printf("file has error: %s", err)
		return
	}
	defer file.Close()
	iReader := bufio.NewReader(file)
	for {
		// 逐行读取
		str, err1 := iReader.ReadString('\n')
		if err1 != nil {
			return
		}
		fmt.Printf("The input was %s", str)
	}
	// 如何通过一个切片读写文件：
}
