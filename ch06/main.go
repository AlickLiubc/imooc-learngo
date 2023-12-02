package main

import "fmt"

// 函数参数传递的时候，值传递，引用传递，go均使用的是值传递
func add(a, b int) (int, error) {
	a = 3
	return a + b, nil
}

func main() {
	// go函数支持普通函数、匿名函数、闭包
	/*
		go中函数是"一等公民"
		1.函数本身是可以当做变量的
		2.匿名函数 闭包
		3.函数可以满足接口
	*/
	a := 1
	b := 2
	sum, _ := add(a, b)

	fmt.Println(sum)
}
