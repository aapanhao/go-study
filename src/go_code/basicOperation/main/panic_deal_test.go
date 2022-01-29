package main

import (
	"testing"
)

func TestDicFunc(t *testing.T) {
	num1 := 10
	num2 := 0
	res := divFunc(num1, num2)
	if res != -1 {
		t.Fatalf("错误，结果为%v， 期望为-1", res)
	} else {
		t.Logf("捕获了异常，返回值%v", res)
	}
}
