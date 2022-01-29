// 错误panic处理，捕获
// 手动创建一个错误panic

package main

import "fmt"

func divFunc(num1 int, num2 int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("divFunc发生错误")
		}
	} () // 匿名函数的调用
	res := num1 / num2
	return res
}
