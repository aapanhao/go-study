package main

import "testing"

func TestRealModifySlice(t *testing.T) {
	slice := make([]int, 5)
	realModifySlice(slice)
	if slice[0] != 9 {
		t.Fatalf("错误，没有修改切片的值,%v", slice)
	} else {
		t.Logf("正确，修改了切片的值,%v", slice)
	}
}

func TestRealModifyMap(t *testing.T) {
	myMap := make(map[string]string)
	myMap["name"] = "pan"
	realModifyMap(myMap)
	if myMap["name"] != "changeName" {
		t.Fatalf("错误，没有修改集合的值,%v", myMap)
	} else {
		t.Logf("正确，修改了集合的值,%v", myMap)
	}
}

func TestRealModifyChannel(t *testing.T) {
	channel := make(chan int, 10)
	realModifyChannel(channel)
	if i := <-channel; i != 9 {
		t.Fatalf("错误，没有修改管道的值,%v, %T", channel, channel)
	} else {
		t.Logf("正确，修改了管道的值,%v, %T", channel, channel)
	}
}
