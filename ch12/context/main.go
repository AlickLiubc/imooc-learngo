package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 我们的新需求，我可以主动退出监控程序
// 共享变量

// func cpuInfo(stop chan struct{}) {
func cpuInfo(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出CPU监控")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("CPU的信息")
		}

		//if stop {
		//	break
		//}

	}

}

func main() {
	// 渐进式的方式
	// 有一个goroutine监控CPU的信息
	wg.Add(1)

	// context包提供了3种函数，
	// withCancel， withTimeout, withValue
	// 如果你的goroutine，函数中，如果希望被控制
	// 超时、传值，但是我不希望影响原来的接口
	// 函数传参中第一个参数放ctx
	ctx, cancel := context.WithCancel(context.Background())

	//var stop = make(chan struct{})

	go cpuInfo(ctx)

	time.Sleep(6 * time.Second)
	// stop = true
	// stop <- struct{}{}
	cancel()
	wg.Wait()

	fmt.Println("监控完成")
}
