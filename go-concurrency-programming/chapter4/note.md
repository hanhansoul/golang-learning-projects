# Go的并发机制

## 原理探究

go不推荐使用共享内存的方式传递数据，而推荐使用channel在多个goroutine之间传递数据，并且还会保证整个过程的并发安全性。

### 线程实现模型

1. M：machine，一个M代表一个内核线程，或称工作线程
2. P：processor，一个P代表执行一个go代码片段所需的资源
3. G：goroutine，一个G代表一个go代码片段

