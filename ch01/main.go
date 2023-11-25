package main

import "fmt"

// 全局变量和局部变量
//var name = "bobby"
//var age = 18
//var ok bool

var (
	name = "bobby"
	age  = 16
	ok   bool
)

func main() {
	// var name int
	// name = 1

	// var name = 1

	age := 1
	var age2 int
	var name2 string

	fmt.Println(age)
	fmt.Println(age2)
	fmt.Println(name2)

	// 多变量定义
	var user1, user2, user3 = "bobby1", 1, "bobby3"
	fmt.Println(user1, user2, user3)

	// 注意：
	// 变量必须先定义才能使用
	// go语言是静态语言，要求变量的类型和赋值类型一致
	// 变量名不能冲突
	// 简洁变量定义不能用于全局变量
	// 变量是有零值的
}
