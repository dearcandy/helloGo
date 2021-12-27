package main

import "fmt"

/**
使用递归函数从 10 打印到 1
*/
func main() {
	PrintNum(10)
}

func PrintNum(num int) {
	fmt.Println(num)
	if num > 1 {
		PrintNum(num - 1)
	}
}
