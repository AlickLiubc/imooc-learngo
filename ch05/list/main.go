package main

import (
	"container/list"
	"fmt"
)

func main() {
	// var mylist list.List

	mylist := list.New()

	// 放在尾部
	mylist.PushBack("go")
	mylist.PushBack("grpc")
	mylist.PushBack("mysql")
	//fmt.Println(mylist)

	// 放在头部
	mylist.PushFront("gin")

	// 遍历打印值，正序
	for i := mylist.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	//// 反向遍历
	//for i := mylist.Back(); i != nil; i = i.Prev() {
	//	fmt.Println(i.Value)
	//}

	// 集合类型4种
	// 1.数组 - 不同长度的数组类型不一样
	// 2.切片 - 动态数组，用起来很方便，性能比较好
	// 3.map
	// 4.list
}
