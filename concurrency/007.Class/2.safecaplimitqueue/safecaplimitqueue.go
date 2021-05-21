package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type SafeSliceQueue struct {
	*sync.Cond
	data []interface{}
	cap  int
	logs []string
}

func NewSafeSliceQueue(capacity int) *SafeSliceQueue {
	return &SafeSliceQueue{
		Cond: &sync.Cond{L: &sync.Mutex{}},
		data: make([]interface{}, 0, capacity),
		cap:  capacity,
		logs: make([]string, 0),
	}
}

func (q *SafeSliceQueue) Enqueue(v interface{}) {
	q.L.Lock()
	defer q.L.Unlock()
	for len(q.data) == q.cap {
			q.Wait()
	}
	// FIFO 入队
	q.data = append(q.data, v)
	//记录操作日志
	q.logs = append(q.logs, fmt.Sprintf("Enqueue %v\n", v))
	// 通知其他waiter 进行Deque 或 Enqueue操作
	q.Broadcast()
}

// 出队

func (q *SafeSliceQueue) Dequeue() interface{} {
	q.L.Lock()
	defer q.L.Unlock()

	for len(q.data) == 0 {
		q.Wait()
	}
	// FIFO 出队
	v := q.data[0]
	q.data = q.data[1:]
	// 记录操作日志
	q.logs = append(q.logs, fmt.Sprintf("Dequeue %v\n", v))

	// 通知其他的waiter 进行Deque 或 Enqueue 操作
	q.Broadcast()
	return v
}

func (q *SafeSliceQueue) Len() int {
	q.L.Lock()
	defer q.L.Unlock()
	return len(q.data)
}

func (q *SafeSliceQueue) String() string {
	var b strings.Builder
	for _, log := range q.logs {
		b.WriteString(log)
	}
	return b.String()
}

func example() {
	var wg sync.WaitGroup
	// 容量为5 的阻塞队列
	queue := NewSafeSliceQueue(5)
	// 生成随机命令
	wg.Add(20)
	for i, cmd := range Commands(20, true) {
		// 0 表示入队 1表示出队
		if cmd == 0 {
			go func(id int) {
				defer wg.Done()
				queue.Enqueue(id)
			}(i)
		} else {
			go func(id int) {
				defer wg.Done()
				queue.Dequeue()
			}(i)
		}

		// 输出操作日志
		//fmt.Println(queue)
	}
	wg.Wait()
}

func Commands(N int, random bool) []int {
	if N%2 != 0 {
		panic("will deadlock!")
	}
	// 0 表示入队 1表示出队
	commands := make([]int, N)
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			commands[i] = 1
		}
	}
	if random {
		// shuffle algorithms
		for i := len(commands) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			commands[i], commands[j] = commands[j], commands[i]
		}
	}
	return commands
}

func demo() {
	//var wg = sync.WaitGroup{}
	queue := NewSafeSliceQueue(1)
	go func() {
		for true {
			fmt.Println("开始等待：")
			v := queue.Dequeue()

			fmt.Println("dequeue:", v)
		}
	}()

	t1 := time.Tick(3 * time.Second)

	go func() {
		//for true {
		//	i := 1
		//	select {
		//	case <-time.Tick(10):
		//		queue.Enqueue(i)
		//	}
		//	i++
		//}
		i := 1
		for true {
			select {
			case <-t1:
				//queue.Enqueue(2)
				fmt.Println("定时器：")
				queue.Enqueue(i)
			}
			fmt.Println("select end")
			i++
		}

	}()

	time.Sleep(time.Second * 1000)

}

func main() {
	//example()
	demo()
}
