package main

import "fmt"

func fakeModifyArray(arr [3]int) {
	arr[0] = 1
	fmt.Println("testModifyArray", arr)
}

func realModifyArray(arr *[3]int){
	// 使用地址，可以修改数组的值；内部影响外部
	(*arr)[0] = 1
	fmt.Println("testModifyArrayAddr", *arr)
}


type person struct {
	name string
	age int
}

func fakeModifyStruct(p person) {
	p.name = "changeName"
}

// 使用指针类型，真正改变结构体
func realModifyStruct(p *person) {
	p.name = "changeName"
}