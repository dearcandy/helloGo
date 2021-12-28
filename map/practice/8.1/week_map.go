package main

import "fmt"

/**
创建一个 map 来保存每周 7 天的名字，将它们打印出来并且测试是否存在 Tuesday 和 Hollyday。
*/
func main() {
	var isPresent bool
	weekMap := make(map[string]string)
	weekMap["Sunday"] = "周日"
	weekMap["Monday"] = "周一"
	weekMap["Tuesday"] = "周二"
	weekMap["Wednesday"] = "周三"
	weekMap["Thursday"] = "周四"
	weekMap["Friday"] = "周五"
	weekMap["Saturday"] = "周六"
	_, isPresent = weekMap["Tuesday"]
	for key, value := range weekMap {
		fmt.Printf("%s Chinese is %s \n ", key, value)
	}
	fmt.Println(isPresent)
}
