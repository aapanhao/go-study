package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func loopPrint(ctx context.Context, i int, wg *sync.WaitGroup) {
	for {
		select {
		case v := <-ctx.Done():
			fmt.Println("协程执行完", i, v)
			wg.Done()
			return
		default:
			fmt.Println("正在执行", i)
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func testManualClose() {
	// 定义可取消的context
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(4)

	for i := 0; i < 4; i++ {
		go loopPrint(ctx, i, &wg)
	}

	time.Sleep(time.Second)
	cancel()

	wg.Wait()
	fmt.Println("退出主程序")

}



