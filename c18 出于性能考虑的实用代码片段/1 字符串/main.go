package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	// (1) 修改字符串：
	str := "hello"
	c := []byte(str)
	c[0] = 'c'
	s2 := string(c) // s2 == "cello"
	fmt.Println(s2)
	// (2) 获取字符串子串
	substr := str[2:4]
	fmt.Println(substr)
	// (3) for-range 遍历字符串

	// (4) 获取一个字符串的字节数：len(str)
	fmt.Println(utf8.RuneCountInString("你好")) // 最快速
	fmt.Println(len([]rune("你好")))
	// (5) 如何连接字符串
	var buffer bytes.Buffer // 最快：
	tmpStrs := []string{"你好", "hello", ",", "世界"}
	for _, v := range tmpStrs {
		buffer.WriteString(v)
	}
	fmt.Println(buffer.String())
	// Strings.Join()
	// +=
}
