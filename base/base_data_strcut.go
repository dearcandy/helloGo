package main

import (
	"fmt" //fmt 包实现了格式化 IO（输入 / 输出）的函数
	"math"
	_ "os"
)

// 定义一个常量
const constTest = "const\"" +
	"Test"

var dividend = 7
var divisor = 2

var Pi float64

func init() {
	Pi = 4 * math.Atan(1) // init() function computes Pi
}

func testFun(a int, b int) int {
	return a / b
}

func main() {
	q := testFun(dividend, divisor)
	fmt.Println(q)
	fmt.Println("test", dividend/divisor)
}
