#Once

#实现Once 

初版的Once

```
type Once struct {
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32((*uint32)(unsafe.Pointer(&o.done))) == 1 {
		return
	}
	f()
	atomic.AddUint32((*uint32)(unsafe.Pointer(&o.done)), 1)
}
```
这种情况 会有一个问题 当f() 很耗时的时候,并发调用Do的时候会 重复的初始化

如果 是这样的话
```
func (o *Once) Do(f func()) {
	if atomic.LoadUint32((*uint32)(unsafe.Pointer(&o.done))) == 1 {
		return
	}
    atomic.AddUint32((*uint32)(unsafe.Pointer(&o.done)), 1)
	f()
}
```
则会出现获取到初始化为空的内容.

所以要引入一个新的函数slowDo()

从逻辑上讲,这里有两个逻辑
一个是 初始化入口 要限制,  一个是开始实际执行初始化 要限制
```
func (o *Once) doSlow(f func()) {
	o.Lock()
	defer o.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
```


封装一个 可以查看one是否已经初始化成功的once
封装一个 once初始化失败 重复初始化的once