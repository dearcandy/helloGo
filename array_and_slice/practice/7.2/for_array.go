package main

/**
for_array.go: 写一个循环并用下标给数组赋值（从 0 到 15）并且将数组打印在屏幕上。
*/
import "fmt"

func main() {
	var arr [16]int

	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	fmt.Println(arr)
}
