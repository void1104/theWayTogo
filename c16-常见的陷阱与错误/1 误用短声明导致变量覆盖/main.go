package main

import "fmt"

/**
注意：代码块内的短变量声明会把代码块外的变量覆盖
*/

func check1() (int, error) {
	// do something
	return 0, nil
}

func check2() (int, error) {
	// do something
	return 0, nil
}

func shadow() (err error) {
	x, err := check1() // x是新创建变量，err是被赋值
	if err != nil {
		fmt.Println(x)
		return // 正确返回err
	}
	if y, err := check2(); err != nil { // y 和 if语句中err被创建
		return // if语句中的err覆盖外面的err，所以错误的返回nil！
	} else {
		fmt.Println(y)
	}
	return
}

func main() {
	shadow() // 报错：err is shadowed during return
}
