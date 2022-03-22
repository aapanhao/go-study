package main

import (
	"fmt"
	"time"
)

func readNoCacheChan(myChan chan int) {
	for {
		fmt.Println("number:", <-myChan)
	}

}

func noCacheChan() {

	myChan := make(chan int)
	go readNoCacheChan(myChan)

	for i := 0; i < 4; i++ {
		myChan <- i
	}

	time.Sleep(time.Second)

}

func safeSend(myChan chan bool, i bool) (close bool) {
	fmt.Println("开始插入")

	defer func() {
		if recover() != nil {
			fmt.Println("发生异常")
			close = true
		}
	}()

	myChan <- i
	close = false
	fmt.Println("正常插入")
	return

}
func safeClose(ch chan bool) (justClosed bool) {
	defer func() {
		if recover() != nil {
			// 一个函数的返回结果可以在defer调用中修改。
			justClosed = false
		}
	}()

	// 假设ch != nil。
	close(ch)   // 如果ch已关闭，则产生一个恐慌。
	return true // <=> justClosed = true; return
}

func testSafeSend() {
	c := make(chan bool)
	go func() {
		for i := 0; i < 25; i++ {
			safeSend(c, true)
		}
	}()
	safeClose(c)
	time.Sleep(time.Second)

}
