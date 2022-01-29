// 各个类型的创建与区别：值类型：数组、结构体；指针类型：切片、集合、管道
// 单元测试方法的使用

package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	arr := [...]int{0, 1, 2}
	fmt.Println("arr的值", arr)
	// 相等
	fmt.Println("arr和新声明的数组是否相等：", arr == [3]int{0, 1, 2})

	p1 := Person{"pan", 20}
	p2 := p1
	// 相等
	fmt.Println("p1 和 p2 是否相等：", p1 == p2)
	fmt.Println("p1 和 p2 地址是否相等：", &p1 == &p2)

	slice := make([]int, 10)
	// 是指向同一地址空间
	oriSlice := slice
	slice[0] = 1
	fmt.Println(&slice[0])
	// 相等
	fmt.Println(&slice[0] == &oriSlice[0])
	fmt.Println(slice, oriSlice)
}
