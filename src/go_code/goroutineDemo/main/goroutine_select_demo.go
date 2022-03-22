// select的实践
// 在管道不close的情况下，使用select，遍历管道

package main

import "fmt"

func putChanData(myChan chan int) {
	for i:=1;i<10;i++{
		myChan <- i
	}
}

func getNotCloseChanData(myChan chan int) {
	for {
		select {
			case v:= <- myChan:
				fmt.Println(v)
			default:
				fmt.Println("no data")
				return
		}
	}
}
