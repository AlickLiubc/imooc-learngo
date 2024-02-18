package main

import (
	"fmt"
	"time"
)

// python,java,php 多线程编程，多进程编程，主要问题是耗费内存
// 内存，线程切换，web2.0 用户级的线程，绿程，轻量级线程，协程 asyncio-python swoole netty

// 内存占用小(2k)，切换快, go语言的协程，go语言诞生之后就只有协程可用 goroutine，非常方便

func asyncPrint() {
	for {
		time.Sleep(time.Second)
		fmt.Println("bobby")
	}
}

// 主协程
func main() {
	// 主死随从
	// go asyncPrint()

	// 匿名函数启动goroutine
	//go func() {
	//	for {
	//		time.Sleep(time.Second)
	//		fmt.Println("bobby")
	//	}
	//}()

	// 1.闭包
	// 2.for循环的问题 for循环的时候，每个变量会重用
	// 每次for循环的时候，i变量会被重用
	// 当我进行到第2轮for循环的时候，这个i就变了
	for i := 0; i < 100; i++ {
		// 方式1
		//tmp := i
		//go func() {
		//	fmt.Println(tmp)
		//}()

		// 方式2 - 更好
		go func(i int) {
			fmt.Println(i)
		}(i) // 值传递
	}

	fmt.Println("main goroutine")
	time.Sleep(10 * time.Second)
}
