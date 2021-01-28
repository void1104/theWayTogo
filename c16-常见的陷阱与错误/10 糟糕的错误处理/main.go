package main

/**
- 避免错误检测使代码变得混乱
	- 避免每调用一个方法就进行一次错误检测：会使程序逻辑和错误检测混杂
*/

/**
- 更好的方法，详细参考13.5小节
	- 缺点：
		- 1.所有的函数签名必须一致（也许可以用interface{}解决）
		- 2.效率也许存在问题
		- 3.所有调用方法外必须再包一层errorHandler()
*/
//func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
//	err := func() error {
//		if req.Method != "GET" {
//			return errors.New("expected GET")
//		}
//		if input := parseInput(req); input != "command" {
//			return errors.New("malformed command")
//		}
//		// 可以在此进行其他的错误检测
//		return nil
//	}()
//
//	if err != nil {
// 		w.WriteHeader(400)
// 		io.WriteString(w, err)
//      return
// }
//  doSomething()...
//}
