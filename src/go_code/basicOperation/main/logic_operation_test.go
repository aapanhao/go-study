package main

import "testing"

func TestSumFunc(t *testing.T) {
	start, end := 1, 100
	res1 := sumFunc1(start, end)
	if res1 != 2550 {
		t.Fatalf("SumFunc1计算从%v到%v的偶数错误, 计算结果是%v", start, end, res1)
	} else {
		t.Logf("SumFunc1计算从%v到%v的偶数正确, 计算结果是%v", start, end, res1)
	}

	res2 := sumFunc1(start, end)
	if res2 != 2550 {
		t.Fatalf("SumFunc2计算从%v到%v的偶数错误, 计算结果是%v", start, end, res1)
	} else {
		t.Logf("SumFunc2计算从%v到%v的偶数正确, 计算结果是%v", start, end, res1)
	}
}
