# 为什么会有读写锁

互斥锁 忽略锁的类型,当所有的操作是读的时候,没必要一个个等待锁.可以并发的访问共享变量.将串行的的读变成并行读.提高操作的性能.

# 读写操作锁 优先级

1.读优先  当所有的读释放锁,写才能获取到锁.
2.写优先  如果有写操作锁,会阻止新的reader 获取到锁.新请求的writer 也会等待已经存在的reader都释放之后才能获取.
写优先级 是针对新来的请求而言的.这种设计主要是为了避免饥饿.[Go 标准库中的rwMutex 就是写优先级]

写的优先权 是针对后来的reader 不要和它抢. 

3.不指定优先级


```
type RwMutex struct{

	w Mutex //互斥锁解决多个wirter竞争
	writerSem uint32 // writer信号量
	readerSem uint32 // reader 信号量
	readerCount uint32 // reader的数量
	readerWait int32   // writer等待完成的reader的数量

}

```
readerCount 这个字段有双重含义

负数表示 当前有wirter锁  同时代表了reader的数量
# 易错场景
1.rwmutex 也是不可重入锁.重入导致死锁
2.不可以复制

#扩展一个trylock RwMutex锁

3.Demo 这个方案更简洁


