package main

import "testing"

func TestReflectDemo(t *testing.T) {
	//num := 10
	//transformVariable(num)
	//variableReflectType(num)
}

func TestReflectStructDemo(t *testing.T) {

	type Student struct {
		Name  string `json:"my_name"`
		Age   int    `json:"my_age"`
		Hobby string
	}

	s := Student{"pan", 18, "fly"}

	reflectStructField(s)

}
