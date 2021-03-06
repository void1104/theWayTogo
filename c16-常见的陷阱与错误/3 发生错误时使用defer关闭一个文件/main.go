package main

func main() {
	// 如果你在一个for循环内处理一系列文件，你需要使用defer确保文件在处理完毕后被关闭
	//for _, file := range files {
	//	if f, err = os.Open(file); err != nil {
	//		return
	//	}
	//	// 这是错误的方式，当循环结束时文件没有关闭，因为defer是在函数终止时执行
	//	//defer f.close()
	//	// 对文件进行操作
	//	//f.Process(data)
	//}

	// 更好的做法：
	//for _, file := range files {
	//	if f, err = os.Open(file); err != nil {
	//		return
	//	}
	//	// 对文件进行操作
	//	f.Process(data)
	//	// 关闭文件
	//	f.Close()
	//}
}
