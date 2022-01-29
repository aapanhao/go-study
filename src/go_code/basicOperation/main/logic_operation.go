// 判断与循环
// 取和，1到100的所有偶数

package main

func sumFunc1(start int, end int) (sum int) {
	sum = 0
	for i := start; i <= end; i++ {
		if i % 2 == 0 {
			sum += i
		}
	}
	return
}

func sumFunc2(start int, end int) (sum int) {
	sum = 0
	for i := start; i <= end; i++ {
		if i&1 == 0 {
			sum += i
		}
	}
	return
}