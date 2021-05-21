# Cond 的基本用法

type Cond

func NewCond(l Locker) *Cond

func (c *Cond)BroadCast()
func (c *Cond)Signal()
func (c *Cond)Wait()


Signal 方法 
允许调用者 Caller 唤醒一个等待此 Cond 的 goroutine。如果此时没有等待的 goroutine，显然无需通知 waiter；
如果 Cond 等待队列中有一个或者多个等待的 goroutine，则需要从等待队列中移除第一个 goroutine 并把它唤醒。
在其他编程语言中，比如 Java 语言中，**Signal 方法也被叫做 notify 方法。**

Broadcast 方法  ，允许调用者 Caller 唤醒所有等待此 Cond 的 goroutine。如果此时没有等待的 goroutine，显然无需通知 waiter；
如果 Cond 等待队列中有一个或者多个等待的 goroutine，则清空所有等待的 goroutine，并全部唤醒。
在其他编程语言中，比如 Java 语言中，**Broadcast 方法也被叫做 notifyAll 方法。**

Wait 方法 会把调用者 Caller 放入 Cond 的等待队列中并阻塞，**直到被 Signal 或者 Broadcast 的方法从等待队列中移除并唤醒。**