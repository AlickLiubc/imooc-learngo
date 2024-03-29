package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}

	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Printf("num=%d\r\r", num)
	}
}

func main() {
	// 默认情况下，channel是双向的
	// 但是，我们经常一个channel作为参数进行传递，希望对方是单向使用

	//// 双向channel
	//var ch1 chan int
	//
	//// 单向channel，只能写入float64的数据
	//var ch2 chan<- float64
	//
	//// 单向的，只能读取
	//var ch3 <-chan int

	//c := make(chan int, 3)
	//// send-only
	//var send chan<- int = c
	//// read-only
	//var read <-chan int = c
	//
	//send <- 1
	//<-read

	c := make(chan int)
	go producer(c)

	go consumer(c)

	time.Sleep(10 * time.Second)
}
