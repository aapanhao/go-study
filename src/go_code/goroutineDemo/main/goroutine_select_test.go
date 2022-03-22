package main

import "testing"

func TestSelectDemo(t *testing.T) {
	myChan := make(chan int, 10)
	putChanData(myChan)
	getNotCloseChanData(myChan)
}