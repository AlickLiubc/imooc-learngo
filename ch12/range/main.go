package main

import (
	"fmt"
	"time"
)

func main() {
	var msg chan int

	msg = make(chan int, 2)
	go func(msg chan int) {
		for data := range msg {
			fmt.Println(data)
		}

		fmt.Println("all done")
	}(msg)

	// 放值到channel中
	msg <- 1
	// 放值到channel中
	msg <- 2

	// 其他的编程语言有很大的区别
	close(msg)

	// 已经关闭的channel不能继续放值
	// msg <- 3

	// 已经关闭的channel可以继续读取值
	d := <-msg
	fmt.Println(d)

	time.Sleep(3 * time.Second)
}
