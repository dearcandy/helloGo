package main

import "fmt"

/**
编写一个函数，接收两个整数，然后返回它们的和、积与差。编写两个版本，一个是非命名返回值，一个是命名返回值。
*/
func main() {
	fmt.Println(ReturnVal1(2, 3))
	fmt.Println(ReturnVal2(2, 3))

}

func ReturnVal1(n1 int, n2 int) (sum int, product int, difference int) {
	sum = n1 + n2
	product = n1 * n2
	difference = n1 - n2
	return
}

func ReturnVal2(n1 int, n2 int) (int, int, int) {
	return n1 + n2, n1 * n2, n1 - n2
}
