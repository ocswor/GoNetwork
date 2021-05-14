Go 的 const 语法提供了“隐式重复前一个非空表达式”的机制
iota 的值 从 0 开始递增

```
const (
	mutexLocked      = 1 << iota // 加锁标识位置
	mutexWoken                   // 唤醒标识位置
	mutexStarving                //锁饥饿标识位置
	mutexWaiterShift = iota      //标识waiter的起始bit位置

)
```
mutexLocked    == 1<<0   1     2的0次方
mutexWoken     == 1<<1   2     2的1次方
mutexStarving  == 1<<2   4     2的2次方
mutexLocked    == 3      3