package main

import "fmt"

// 两个知识盲区 位操作符 以及const

// 复制Mutex 定义的常量
const (
	mutexLocked      = 1 << iota // 加锁标识位置
	mutexWoken                   // 唤醒标识位置
	mutexStarving                //锁饥饿标识位置
	mutexWaiterShift = iota      //标识waiter的起始bit位置
	test
	de
)

func main() {
	fmt.Println(mutexLocked)
	fmt.Println(mutexWoken)
	fmt.Println(mutexStarving)
	fmt.Println(mutexWaiterShift)
	fmt.Println(test)
	fmt.Printf("de:%d\n", de)

	fmt.Println(1 << 0)
	fmt.Println(1 << 1)
	fmt.Println(1 << 2)
	fmt.Println(1 << 62) //int 是 8个字节，8位，也就是64位
	//最高位为符号位，0代表正数，1代表负数。剩余位存储实际值。
	//0 (62个0) 1 也是就说 1最大可以移动62位

	i:= 1 << 62
	println(&i)
	j:= 1<<2
	println(&j)
	a := mutexLocked | mutexWoken | mutexStarving
	fmt.Printf("a:%d\n", a) // 7

	b := mutexLocked | mutexWoken //int 占8个字节，这里用一个字节 8位来表示，其他字节都是0
	// mutexLocked: 00000001    mutexWoken: 00000010          #1或上任何值结果都为1
	// 结果 : 00000011
	fmt.Printf("b:%d\n", b) // 3

	var old = 0

	c := old & (mutexLocked | mutexWoken | mutexStarving) // 0与上任何值结果都为0
	fmt.Printf("c:%d\n", c)                               // 0

}
