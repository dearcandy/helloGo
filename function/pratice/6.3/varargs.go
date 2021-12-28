package main

import "fmt"

/**
写一个函数，该函数接受一个变长参数并对每个元素进行换行打印。
一个接受变长参数的函数可以将这个参数作为其它函数的参数进行传递：
*/

func main() {
	Varargs("arg1", "arg2", "arg3")
}

func Varargs(args ...string) {
	for i := 0; i < len(args); i++ {
		fmt.Println(args)
	}

}
