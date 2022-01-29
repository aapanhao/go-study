// 接口的声明
// 结构体实现接口
// 多态的使用，基于接口编程

package main

import "fmt"

// Flyable 定义接口
type Flyable interface {
	// 定义接口的方法
	fly()
}
type Duck struct {
	Name string
	Age  int
}

type RedDuck struct {
	Duck
}

// 结构体实现接口，把接口方法都实现了才称为实现
// 这里不能是指针类型，必须是结构体类型
func (d RedDuck) fly() {
	fmt.Println("i can fly")
}

type RubberDuck struct {
	Duck
}

func (d RubberDuck) fly() {
	fmt.Println("i mock fly")
}

// 基于接口编程
// 实现了多态(根据传入类型，判断具体实现哪个方法)
// 因为RedDuck和RubberDuck都实现了Flyable接口
func action(duck Flyable) {
	fmt.Printf("我是什么类型, %T\n", duck)
	duck.fly()
}
