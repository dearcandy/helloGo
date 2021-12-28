package main

import "fmt"

/**
写一个 minSlice 方法，传入一个 int 的切片并且返回最小值，再写一个 maxSlice 方法返回最大值。
*/
func main() {
	slice := []int{2, 3, 5, 8, 4, 7, 6, 9}
	fmt.Printf("slice's min value is %d \n\n", MinSlice(slice))
	fmt.Printf("slice's max value is %d \n\n", MaxSlice(slice))

}

func MinSlice(slice []int) int {
	var min = slice[0]
	for i := 0; i < len(slice); i++ {
		var cur = slice[i]
		if cur < min {
			min = cur
		}
	}
	return min
}

func MaxSlice(slice []int) int {
	var max = slice[0]
	for i := 0; i < len(slice); i++ {
		var cur = slice[i]
		if cur > max {
			max = cur
		}
	}
	return max
}
