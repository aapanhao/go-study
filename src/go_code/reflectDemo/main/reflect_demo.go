// 案例一：为了使用而使用，将一个变量，通过空接口方式（函数参数是空间口），转换为reflect.Value，再转换为interface{}，再转换为变量
// 案例二：反射获取结构体的tag
package main

import (
	"fmt"
	"reflect"
)

func transformVariable(i interface{}) {
	val := reflect.ValueOf(i)
	fmt.Printf("value 值%v, 类型%T, %v\n", val, val, val.Kind())

	// 转换为接口类型
	iVal := val.Interface()
	fmt.Printf("iVal 值%v, 类型%T\n", iVal, iVal)

	// 如果想要使用整数类型的操作，一种方式是需要类型断言
	num := 1 + iVal.(int)
	fmt.Printf("num 值%v, 类型%T\n", num, num)

	// 如果想要使用整数类型的操作，另一种方式是是用reflect.Value.Int()
	num2 := 2 + val.Int()
	fmt.Printf("num2 值%v, 类型%T\n", num2, num2)
}

func variableReflectType(i interface{}) {
	iType := reflect.TypeOf(i)
	fmt.Printf("iType 值%v, 类型%T, %v\n", iType, iType, iType.Kind())
}

// 反射获取结构体的tag
func reflectStructField(i interface{}) {
	val := reflect.ValueOf(i)

	// 通过反射 Value 获取字段值
	num := val.NumField()
	fmt.Printf("%v的字段有%v个\n", val, num)
	for i := 0; i < num; i++ {
		fmt.Printf("第%v个字段的值为%v\n", i+1, val.Field(i))
	}

	// 通过反射 Type 获取字段值
	iType := reflect.TypeOf(i)
	num = iType.NumField()
	fmt.Printf("%v的字段有%v个\n", iType, num)
	for i := 0; i < num; i++ {
		fmt.Printf("第%v个字段的值为%v\n", i+1, iType.Field(i))
		fmt.Println()
		fmt.Printf("第%v个字段的tag为%v\n", i+1, iType.Field(i).Tag.Get("json"))
	}
	// 通过反射 Type 获取 结构体的方法
	iType.Method(0)

}
