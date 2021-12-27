package main

/**
编写一个名字为 MySqrt 的函数，计算一个 float64 类型浮点数的平方根，如果参数是一个负数的话将返回一个错误。
编写两个版本，一个是非命名返回值，一个是命名返回值。
*/
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(MySqrt(3.0))
	fmt.Println(MySqrt1(3.0))
}

func MySqrt(value float64) (sqrt float64) {
	sqrt = math.Sqrt(value)
	return
}

func MySqrt1(value float64) float64 {
	return math.Sqrt(value)
}
