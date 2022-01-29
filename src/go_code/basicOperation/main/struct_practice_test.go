package main

import (
	"fmt"
	"testing"
)

func TestPersonStruct(t *testing.T) {
	person := GetPerson("pan", 20)
	name := person.GetName()
	fmt.Printf("person的类型%T, 值%v, name值%v", person, person, name)
}

func TestStudentStruct(t *testing.T) {
	student := GetStudent("pan", 20, "一班")
	name := student.GetName()
	class := student.GetClass()
	fmt.Printf("person的类型%T, 值%v, name值%v, class值:%v", student, student, name, class)
}