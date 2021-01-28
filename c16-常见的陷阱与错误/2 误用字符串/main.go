package main

/**
当需要对一个字符串进行频繁的操作时，谨记在GO语言中字符串是不可变的（类似java和c#）.
使用诸如a+=b形式连接字符串效率低下，尤其在一个循环内部使用这种形式，会导致大量的内存开销和拷贝
*：应该使用一个字符数组代替字符串，将字符串内容写入一个缓存中。
*/

//var b bytes.Buffer
//...
//for condition {
//b.WriteString(str) // 将字符串str写入缓存buffer
//}
//return b.String()
