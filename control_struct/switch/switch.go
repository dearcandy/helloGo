package main

import "fmt"

func main() {
	fmt.Println(Season(13))
}

func Season(month int) string {
	switch month {
	case 1, 2, 3:
		return "It's Spring"
	case 4, 5, 6:
		return "It's Summer"
	case 7, 8, 9:
		return "It's Autumn"
	case 10, 11, 12:
		return "It's Winter"
	default:
		return "The month is not isExist"
	}

}
