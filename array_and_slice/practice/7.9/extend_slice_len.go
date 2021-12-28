package main

/**
给定 array_and_slice s[]int 和一个 int 类型的因子 factor，扩展 s 使其长度为 len(s) * factor。
*/
import "fmt"

func main() {
	slice := make([]int, 2)
	extendSlice := Extend(slice, 5)
	fmt.Printf("array_and_slice len is %d \n", len(slice))
	fmt.Printf("extendSlice len is %d ", len(extendSlice))
}

func Extend(slice []int, factor int) []int {
	var n = len(slice) * factor
	slice = make([]int, n)
	return slice
}
