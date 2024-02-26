package main

import (
	"fmt"
	"sync"
	"time"
)

var done bool
var lock sync.Mutex

func g1() {
	time.Sleep(time.Second)
}

func g2() {
	time.Sleep(2 * time.Second)

	lock.Lock()
	defer lock.Unlock()

	done = true
}

func main() {
	// select 类似 switch case语句
	// 但是select的功能和我们linux里面的select,poll,epoll
	// select主要作用于多个channel

	// 现在有个需求, 我们现在有2个goroutine都在执行，但是呢，在主的goroutine中，
	// 当某一个执行完成以后，这个时候会立马知道

	go g1()
	go g2()

	for {
		if done {
			fmt.Println("done")
			time.Sleep(10 * time.Millisecond)
			return
		}
	}

}
