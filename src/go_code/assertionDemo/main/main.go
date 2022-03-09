package main

import "fmt"

func main() {
	var i interface{} = 10
	t1 := i.(interface{})
	fmt.Println(t1)

	//fmt.Println(i.(type))
}
