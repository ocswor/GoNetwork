package main

import (
	"fmt"
	"runtime"
)

//当一个goroutine发生阻塞，Go会自动地把与该goroutine处于同一系统线程的其他goroutines转移到另一个系统线程上去，以使这些goroutines不阻塞

func showNumber(i int) {
	fmt.Println(i)
}

func main() {
	n := runtime.NumCPU()
	fmt.Println("cpu 核数:", n)
	runtime.GOMAXPROCS(1) // 如果是单核服务器的话  则
	for i := 0; i < 10; i++ {
		go showNumber(i)
	}

	// runtime.Gosched() //这个函数的作用是让当前goroutine让出CPU，好让其它的goroutine获得执行的机会。同时，当前的goroutine也会在未来的某个时间点继续运行。
	//如果是单核服务器的话   如果没有runtime.Gosched() 那么其他的goroutine 则没有机会运行则会直接退出
	fmt.Println("Haha")
}
