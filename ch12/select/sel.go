package main

import (
	"fmt"
	"time"
)

// var done bool
// var lock sync.Mutex

// channel是多线程安全的, channel需要初始化
// 很大时候并不会多个goroutine写入同一个channel
var done = make(chan struct{})

func g1(ch chan struct{}) {
	time.Sleep(2 * time.Second)
	//lock.Lock()
	//defer lock.Unlock()

	//done = true
	//done <- struct{}{}
	ch <- struct{}{}
}

func g2(ch chan struct{}) {
	time.Sleep(3 * time.Second)

	//lock.Lock()
	//defer lock.Unlock()
	//
	//done = true
	//done <- struct{}{}
	ch <- struct{}{}
}

func main() {
	// select 类似于switch case语句
	// 但是select的功能和我们linux里面提供的io的select,poll,epoll
	// select主要作用于多个channel

	// 现在有个需求, 我们现在有2个goroutine都在执行，但是呢，在主的goroutine中，
	// 当某一个执行完成以后，这个时候会立马知道

	g1Channel := make(chan struct{}, 1)
	g2Channel := make(chan struct{}, 2)

	g1Channel <- struct{}{}
	g2Channel <- struct{}{}

	go g1(g1Channel)
	go g2(g2Channel)

	//<- done

	// 我要监控多个channel，任何一个channel返回都知道
	// 1. 某一个分支就绪了，就执行该分支
	// 2. 如果两个都就绪了，先执行哪个？答：随机的。
	//    目的是什么？答：防止饥饿。
	// 应用场景：
	timer := time.NewTimer(time.Second)
	for {
		select {
		case <-g1Channel:
			fmt.Println("g1 done")
		case <-g2Channel:
			fmt.Println("g2 done")
		case <-timer.C:
			//time.Sleep(10 * time.Millisecond)
			fmt.Println("timeout")
			return
		}
	}

	//for {
	//	if done {
	//		fmt.Println("done")
	//		time.Sleep(10 * time.Millisecond)
	//		return
	//	}
	//}
	//fmt.Println("done")

}
