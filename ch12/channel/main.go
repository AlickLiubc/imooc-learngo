package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		不要通过共享内存来通信，而需要通过通信来实现内存共享
		php、python、java、多线程编程的时候，
		更常用的方式是使用一个全局变量
		也会提供消息队列机制、python-queue
		java
		消费者和生产者之间的关系。
		channel 再加上语法糖让使用channel更加简单
	*/
	var msg chan string

	// channel的初始化值，如果为0的话
	// 你放值进去会阻塞

	// 有缓冲和无缓冲的channel
	// 无缓冲的channel
	msg = make(chan string, 0)
	// 有缓冲的channel
	// msg = make(chan string, 1)
	go func(msg chan string) {
		// go有一种happen-before的机制
		data := <-msg
		fmt.Println(data)
	}(msg)

	// 无缓冲的channel
	// msg = make(chan string, 0)
	// 放值到channel中
	msg <- "bobby"
	time.Sleep(time.Second * 10)

	// waitgroup 如果少了done调用，容易deadlock
	// 无缓冲的channel也容易出现

	// data := <-msg
	// fmt.Println(data)
}
