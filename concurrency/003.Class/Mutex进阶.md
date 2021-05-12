## 1.自己实现一个可以重入的互斥锁

因为 Mutex 的实现中没有记录哪个 goroutine 拥有这把锁。理论上，任何 goroutine 都可以随意地 Unlock 这把锁，
所以没办法计算重入条件.


## 3.模拟死锁的场景


在编写Makefile的时候可以检测一下 思索的情况 工具 *vet*
```
go vet main 
```