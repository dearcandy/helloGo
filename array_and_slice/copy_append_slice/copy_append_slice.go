package main

import "fmt"

func main() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	// copy 内置函数将元素从源切片复制到目标切片。
	//（作为一种特殊情况，它还会将字节从字符串复制到字节切片。）
	// 源和目标可能会重叠。
	// Copy 返回复制的元素数，这将是 len(src) 和 len(dst) 中的最小值。
	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}
