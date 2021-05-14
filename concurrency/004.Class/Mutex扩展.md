1、排外锁

为 Mutex 添加一个 TryLock 的方法，也就是尝试获取排外锁。

>当一个 goroutine 调用这个 TryLock 方法请求锁的时候，如果这把锁没有被其他 goroutine 所持有，
>那么，这个 goroutine 就持有了这把锁，并返回 true；如果这把锁已经被其他 goroutine 所持有，
>或者是正在准备交给某个被唤醒的 goroutine，那么，这个请求锁的 goroutine 就直接返回 false，不会阻塞在方法调用上。

简单总结就是先尝试获取锁,获取到锁，则进入正常业务流程，没有获取到锁，则直接返回，不用等待锁。


1.Demo 
封装了一个tryLock 排外锁

2.Demo 
封装过程中 iota用法   与或操作用法

3.Demo
CompareAndSwapInt32 原子操作 替换地址的value

4.Demo unsafe.pointer操作


