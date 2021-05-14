package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func oneDemo() {
	i := int64(1)
	j := 1
	println(&j)
	println(unsafe.Sizeof(j)) // 占 8个字节
	var iPtr *int
	println(&i)
	// iPtr = &i // 错误
	iPtr = (*int)(unsafe.Pointer(&i))
	fmt.Printf("%d\n", *iPtr)
	fmt.Println(iPtr)
}

func twoDemo() {
	str1 := "hello world"
	hdr1 := (*reflect.StringHeader)(unsafe.Pointer(&str1)) // 注1
	fmt.Printf("str:%s, data addr:%d, len:%d\n", str1, hdr1.Data, hdr1.Len)

	str2 := "abc"
	hdr2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	hdr1.Data = hdr2.Data // 注2
	hdr1.Len = hdr2.Len   // 注3
	fmt.Printf("str:%s, data addr:%d, len:%d\n", str1, hdr1.Data, hdr1.Len)
}
func main() {
	oneDemo()
	twoDemo()
}
