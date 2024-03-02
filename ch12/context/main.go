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
	// 这里能拿到一个请求的id，链路id
	fmt.Printf("traceid: %s\r\n", ctx.Value("traceid"))
	// 记录一些日志，这次请求是哪个traceid打印的

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
	// ctx1, cancel := context.WithCancel(context.Background())
	// ctx2, _ := context.WithCancel(ctx1)

	// var stop = make(chan struct{})

	// 2. timeout 主动超时
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)

	// 3. WithDeadline 在时间点cancel
	// go cpuInfo(ctx)

	// 4.withValue
	valueCtx := context.WithValue(ctx, "traceid", "qjw12")
	go cpuInfo(valueCtx)

	// time.Sleep(6 * time.Second)
	//// stop = true
	//// stop <- struct{}{}
	// cancel()
	wg.Wait()

	fmt.Println("监控完成")
}
