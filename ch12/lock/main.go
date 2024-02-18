package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
锁-资源竞争
*/
var total int32
var wg sync.WaitGroup

// var lock sync.Mutex

// lock2 := lock
// 锁能复制么？
// 复制后就失去了锁的效果了
func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		//lock.Lock()
		// 竞争
		//total += 1
		//lock.Unlock()
		atomic.AddInt32(&total, 1)
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		//lock.Lock()
		//total -= 1
		//lock.Unlock()
		atomic.AddInt32(&total, -1)
	}
}

func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)
}
