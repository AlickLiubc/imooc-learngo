package main

import "fmt"

func deferReturn() (ret int) {
	defer func() {
		ret++
	}()

	return 10
}


func main() {
	//// 连接数据库，打开问题件，开始锁
	//// 无论如何，最后都要记得去关闭数据库，关闭文件，
	//// 解锁
	//defer fmt.Println("1")
	//defer fmt.Println("2")
	//
	//defer func() {
	//
	//}()
	//
	//fmt.Println("main")
	//return

	ret := deferReturn()
	fmt.Printf("ret = %d\r\n", ret)
}
