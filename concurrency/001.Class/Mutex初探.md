# Mutex 互斥锁

>在Go 语言里面 互斥锁就是Mutex
## 1.互斥锁解决问题场景

并发问题,多个goroutine 并发更新同一个资源,像计数器,同时更新用户账户信息,秒杀系统;往同一个buffer中并发写入数据等等.如果没有互斥锁就会
出现一些异常情况,比如计数器的计数不准确,用户的账户可能透支,秒杀系统出现超卖,buffer数据混乱,等等,后果都很严重.

## 2.临界区
学习并发控制机制,需要先理解一下什么是临界区
>在并发编程中，如果程序中的一部分会被并发访问或修改，那么，为了避免并发访问导致的意想不到的结果，这部分程序需要被保护起来，
>这部分被保护起来的程序，就叫做临界区。

如果很多线程同步访问临界区，就会造成访问或操作错误，这当然不是我们希望看到的结果。
所以，我们可以使用互斥锁，限定临界区只能同时由一个线程持有.

## 3.可以考虑实现一个秒杀接口 测试一下互斥锁效果. 

## 4.go race detector
```
go run -race main.go
```
增加 -race 参数可以在代码运行时,监控对共享变量的非同步访问.
    1.race 限制 只有在运行时,并且是代码触发执行了data race 才能检测到,假如程序启动,并没有运行bug所在的代码,并不能检测到.
    
### 5.这个工具没明白
编译的时候可以看到 目前还看不明白
```
go tool compile -race -S main.go
```
