#SafeMap


###[concurrent-map](https://github.com/orcaman/concurrent-map)
正如 这里 和 这里所描述的, Go语言原生的map类型并不支持并发读写。concurrent-map提供了一种高性能的解决方案:通过对内部map进行分片，降低锁粒度，从而达到最少的锁等待时间(锁冲突)

在Go 1.9之前，go语言标准库中并没有实现并发map。在Go 1.9中，引入了sync.Map。新的sync.Map与此concurrent-map有几个关键区别。标准库中的sync.Map是专为append-only场景设计的。因此，如果您想将Map用于一个类似内存数据库，那么使用我们的版本可能会受益。
你可以在golang repo上读到更多，这里 and 这里 译注:sync.Map在读多写少性能比较好，否则并发性能很差


### Go test
go test 命令，会自动读取源码目录下面名为 *_test.go 的文件，生成并运行测试用的可执行文件。
输出的信息类似下面所示的样子

也可以直接运行下面的命令test git上远程的包
```
go test "github.com/orcaman/concurrent-map"
```

