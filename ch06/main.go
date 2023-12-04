package main

import (
	"fmt"
	"time"
)

// 函数参数传递的时候，值传递，引用传递，go均使用的是值传递 指针
//func add(a, b int) (sum int, err error) {
func add(desc string, items ...int) (sum int, err error) {
	// var datas [...]string
	// a = 3
	// sum = a + b
	for _, value := range items {
		sum += value
	}

	return sum, err
}

// 省略号


func funForever() {
	for {
		time.Sleep(time.Second)
		fmt.Println("doing")
	}
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
	sum, _ := add(a, b, 3, 4)

	fmt.Println(sum)
}
