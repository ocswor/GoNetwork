[参考链接](http://c.biancheng.net/view/124.html)   
[推荐](https://www.jianshu.com/p/04124b5aa0b2)
```
go test -v -bench=. benchmark_test.go
```
###代码说明如下：
* 第 1 行的-bench=.表示运行 benchmark_test.go 文件里的所有基准测试，和单元测试中的-run类似。    
* 第 4 行中显示基准测试名称，2000000000 表示测试的次数，也就是 testing.B 结构中提供给程序使用的 N。“0.33 ns/op”表示每一个操作耗费多少时间（纳秒）。
```
goos: windows
goarch: amd64
Benchmark_Add
Benchmark_Add-4         1000000000               0.359 ns/op
PASS
ok      command-line-arguments  0.793s

```

#### 自定义测试时间
```
go test -v -bench=. -benchtime=5s benchmark_test.go
```
```
$ go test -v -bench=. -benchtime=5s benchmark_test.go
goos: linux
goarch: amd64
Benchmark_Add-4           10000000000                 0.33 ns/op
PASS
ok          command-line-arguments        3.380s
```

#### 测试内存
```
D:\goproject\GoNetwork\concurrency\010.Gotest>go test -v -bench=Alloc -benchmem benchmark_test.go
goos: windows
goarch: amd64
Benchmark_Alloc
Benchmark_Alloc-4        9446401               126 ns/op              16 B/op          1 allocs/op
PASS
ok      command-line-arguments  1.738s

```
代码说明如下：
* 第 1 行的代码中-bench后添加了 Alloc，指定只测试 Benchmark_Alloc() 函数。
* 第 4 行代码的“16 B/op”表示每一次调用需要分配 16 个字节，“2 allocs/op”表示每一次调用有两次分配。