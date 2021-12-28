package main

import (
	"fmt"
	"reflect"
)

/**
在 main 函数中写一个用于打印 Hello World 字符串的匿名函数并赋值给变量 fv，然后调用该函数并打印变量 fv 的类型。
*/
func main() {
	fv := func() {
		fmt.Println("Hello World")
	}

	fv()
	fmt.Println(reflect.TypeOf(fv))
}
