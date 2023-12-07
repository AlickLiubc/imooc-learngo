package main

import (
	"fmt"
	"time"
)

// 函数参数传递的时候，值传递，引用传递，go均使用的是值传递 指针
//func add(a, b int) (sum int, err error) {
func add(items ...int) (sum int, err error) {
	// var datas [...]string
	// a = 3
	// sum = a + b
	for _, value := range items {
		sum += value
	}

	return sum, err
}

// 省略号


func add2 (a, b int) {
	fmt.Printf("sum is: %d\r\n", a + b)
}


func callback(y int, f func(int, int)) {
	f(y, 2)
}

func cal(myfunc func(items ...int) int ) int {
	//switch op {
	//	case "+":
	//		return func() {
	//			fmt.Println("这是加法")
	//		}
	//	case "-":
	//		return func() {
	//			fmt.Println("这是减法")
	//		}
	//	default:
	//		return func() {
	//			fmt.Println("这不是加法，也不是减法")
	//		}
	//}
	return myfunc()
}

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
	// funcVar := add

	// a := 1
	// b := 2
	// sum, _ := add(a, b, 3, 4)
	// sum, _ := funcVar(a, b, 3, 4)

	// fmt.Println(sum)

	sum := cal(func(items ...int) int {
		sum := 0
		for _, value := range items {
			sum += value
		}
		return sum
	})

	fmt.Println(sum)

	// callback(1, add2)

	localFunc := func(a, b int) {
		fmt.Printf("total is: %d\r\n", a + b)
	}

	// 匿名函数
	callback(1, localFunc)

	// 闭包
	// 有个需求，有个函数每次调用一次返回结果值都增加1次
}
