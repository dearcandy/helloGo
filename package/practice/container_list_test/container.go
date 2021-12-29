package main

import (
	"container/list"
	"fmt"
)

/**
使用 container_list_test/list 包实现一个双向链表，将 101、102 和 103 放入其中并打印出来。
*/
func main() {
	l := list.New()
	element := l.PushBack(103)
	l.PushFront(101)
	l.InsertBefore(102, element)

	// 遍历链表并打印它包含的元素
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
