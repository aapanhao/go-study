// 判断语句、循环语句
// 异常处理
// 结构体
// 接口

package main

import "fmt"

func testRecover() int {
	res := -1
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("divFunc发生错误", err)
		}
	}()
	num1, num2 := 10, 0
	res = num1 / num2
	// 匿名函数的调用
	return res
}

func main() {

	// 位运算
	fmt.Println("位运算，2与1=", 2&1)
	fmt.Println("位运算，3与2=", 3&2)

	res := testRecover()
	fmt.Println("testRecover返回值", res)

}
