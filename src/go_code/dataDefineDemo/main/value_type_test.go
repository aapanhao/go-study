//测试值类型返回是否正确

package main

import "testing"

func TestFakeModifyArray(t *testing.T) {
	arr := [3]int{0, 1, 2}
	fakeModifyArray(arr)
	if arr != [3]int{0, 1, 2} {
		t.Fatalf("错误，改变了数组的值，%v", arr)
	} else {
		t.Logf("成功，数组值没有被改变")
	}
}

func TestRealModifyArray(t *testing.T) {
	arr := [3]int{0, 1, 2}
	realModifyArray(&arr)
	if arr != [3]int{0, 1, 2} {
		t.Logf("成功，改变了数组的值，%v", arr)
	} else {
		t.Fatalf("成功，数组值没有被改变")
	}
}

func TestFakeModifyStruct(t *testing.T) {
	p := person{"pan", 19}
	oriP := p
	fakeModifyStruct(p)

	if p != oriP {
		t.Fatalf("错误，改变了结构体的值，原始值%v, 调用方法后的值%v", oriP, p)
	} else {
		t.Logf("成功，结构体的值没有被改变，原始值%v, 调用方法后的值%v", oriP, p)
	}
}

func TestRealModifyStruct(t *testing.T) {
	p := person{"pan", 19}
	oriP := p
	realModifyStruct(&p)

	if p != oriP {
		t.Logf("成功，改变了结构体的值，原始值%v, 调用方法后的值%v", oriP, p)
	} else {
		t.Fatalf("错误，结构体的值没有被改变，原始值%v, 调用方法后的值%v", oriP, p)
	}
}

// 运行指令
// go test
// go test -v  输出详细信息
