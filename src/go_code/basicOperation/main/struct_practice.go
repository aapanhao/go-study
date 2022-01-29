// 结构体实践
// 结构体方法
// 继承

package main

type person struct {
	name string
	age  int
}

// GetPerson 工厂方法
func GetPerson(name string, age int) *person {
	return &person{name, age}
}

func (p *person) GetName() string {
	return p.name
}

// 继承
type student struct {
	person
	class string
}

func GetStudent(name string, age int, class string) *student {
	return &student{person{name, age}, class}
}

func (s *student) GetClass() string {
	return s.class
}
