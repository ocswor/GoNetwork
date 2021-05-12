package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// 简单获取gid
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	//fmt.Println(string(buf[:n]))

	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			groutine_id := GoID()
			fmt.Println(groutine_id)
			wg.Done()
		}()
	}
	wg.Wait()

}
