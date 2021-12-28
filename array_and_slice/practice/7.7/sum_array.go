package main

import "fmt"

/**
a) 写一个 Sum 函数，传入参数为一个 32 位 float 数组成的数组 arrF，返回该数组的所有数字和。
如果把数组修改为切片的话代码要做怎样的修改？如果用切片形式方法实现不同长度数组的的和呢？
b) 写一个 SumAndAverage 方法，返回两个 int 和 float32 类型的未命名变量的和与平均值。
*/
func main() {
	var arrF = []float32{1.2, 1.3, 1.4}
	fmt.Printf("arrF's Sum is %f \n", Sum(arrF))
	sum, average := SumAndAverage(arrF)
	fmt.Printf("arrF's int type Sum is %d, int type average is %d ", sum, average)

}

func Sum(arrF []float32) float32 {
	var sum float32
	for _, value := range arrF {
		sum += value
	}
	return sum
}

func SumAndAverage(arrF []float32) (int, int) {
	var sum float32
	for _, value := range arrF {
		sum += value
	}
	return int(sum), int(sum / float32(len(arrF)))
}
