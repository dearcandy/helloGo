package main

import (
	"fmt"
)

func main() {
	PrintG(26)
}

func PrintG(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("G")
		}
		fmt.Println("")
	}
}
