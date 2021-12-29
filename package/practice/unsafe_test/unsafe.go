package main

import (
	"fmt"
	"unsafe"
)

/**
通过使用 unsafe 包中的方法来测试你电脑上一个整型变量占用多少个字节。
*/
func main() {
	var num int
	fmt.Println(unsafe.Sizeof(num))
}
