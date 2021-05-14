package main

import (
	"fmt"
	"unsafe"
)

type W struct {
	b int32
	c int64
}

func main() {

	var w *W = new(W)
	//这时w的变量打印出来都是默认值0，0
	fmt.Println(w.b, w.c)
	fmt.Println("b offsetof :",unsafe.Offsetof(w.b))
	fmt.Println("c offsetof :",unsafe.Offsetof(w.c))

	//现在我们通过指针运算给b变量赋值为10

	//unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
	//而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象， uintptr 类型的目标会被回收；
	b := unsafe.Pointer(uintptr(unsafe.Pointer(w)) + unsafe.Offsetof(w.b))

	fmt.Println(b)
	*((*int)(b)) = 10
	//*b = 12 通用指针类型 不能用于赋值 和参与指针运算
	//此时结果就 变成了10，0
	fmt.Println(w.b,w.c)

}
