package main

import (
	"fmt"
	"sync"
	"time"
)

// 锁本质上是将并行的代码串行化了，
// 使用lock肯定会影响性能
// 即使是涉及锁，那么也应该尽量保证并行
// 我们有两组协程，其中一组负责写数据，另一组负责读数据，WEB系统中绝大部分场景都是读多写少
// 虽然有多个goroutine，但是仔细分析我们会发现，读写协程之间应该并发，读和写之间应该串行，读和读之间也不应该并行
// 读写锁

func main() {
	var rwlock sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(6)
	// 写的goroutine
	go func() {
		time.Sleep(3 * time.Second)
		defer wg.Done()
		// 加写锁，写锁会防止别的写锁获取，和读锁获取
		rwlock.Lock()
		defer rwlock.Unlock()
		fmt.Println("get read-write lock")
		time.Sleep(5 * time.Second)
	}()

	// 读的goroutine
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()

			// 加读锁，读锁不会阻止别人的读
			for {
				rwlock.RLock()
				time.Sleep(500 * time.Millisecond)
				fmt.Println("get read lock")
				rwlock.RUnlock()
			}

		}()
	}

	wg.Wait()
}
