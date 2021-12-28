package main

/**
当一个函数在其函数体内调用自身，则称之为递归。
最经典的例子便是计算斐波那契数列，即前两个数为 1，从第三个数开始每个数均为前两个数之和
重写本节中生成斐波那契数列的程序并返回两个命名返回值，即数列中的位置和对应的值，例如 5 与 4，89 与 10。
*/
import "fmt"

func main() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
