package chapter3

import (
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"sync"
	"testing"
	"text/tabwriter"
	"time"
)

/**
sync包

## WaitGroup

当你不关心并发操作的结果，或者你有其他方法来收集它们的结果时，WaitGroup是等待一组并发操作完成的好方法。

可以将WaitGroup视为一个并发一安全的计数器：调用通过传人的整数执行add方法增加计数器的增量，并调用Done方位对计数器进行递减。Wait阻塞，直到计数器为零。

注意，添加的调用是在他们帮助跟踪的goroutine之外完成的。如果我们不这样做，我们就会引入一种竞争条件。

## 互斥锁和读写锁

临界区是你程序中需要独占访问共享资源的区域。Mutex提供了一种安全的方式来表示对这些共享资源的独占访问。
为了使用一个资源，channel通过通信共享内存，而Mutex通过开发人员的约定同步访问共享内存。你可以通过使用
Mutex对内存进行保护来协调对内存的访问。

可能存在需要在多个并发进程之间共享内存的情况，但可能这些进程不是都需要读写此内存。如果是这样，你可以利用
不同类型的互斥对象：sync.RWMutex。

sync.RWMutex在概念上和互斥是一样的：它守卫着对内存的访问，然而，RWMutex让你对内存有了更多控制。你可以
请求一个锁用于读处理，在这种情况下你将被授予访问权限，除非该锁被用于写处理。这意味着，任意数量的读消费者可
以持有一个读锁，只要没有其他事物持有一个写锁。


*/

func TestExample21(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()
	wg.Wait()
	fmt.Println("All goroutines complete.")
}

func TestExample22(t *testing.T) {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}
	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}

func TestExample23(t *testing.T) {
	var count int
	var lock sync.Mutex
	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}
	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}
}

func TestExample24(t *testing.T) {
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1)
		}
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}
		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRwMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}
}

/**
假设我们有一个固定长度为2的队列，还有10个我们想要推送到队列中的项目。我们想要在有房间的情况下尽快排队，所以就希望在队列中有空间时能立即得到通知。

在这个例子中，我们还有一个新方法Signal。这是Cond类型提供的两种方住中的一种，它提供通知goroutine阻塞的调用Wait，条件已经被触发。另一种方法叫做Broadcast。

运行时内部维护一个FIFO列表，等待接收信号；Signal发现等待最长时间的goroutine并通知它，而Broadcast向所有等待的goroutine发送信号。
Broadcast可以说是这两种方怯中比较有趣的一种，因为它提供了一种同时与多个goroutine通信的方法。

此外，与利用channel相比，Cond类型的性能要高很多。

*/
func TestExample25(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)
	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		c.Signal()
	}
	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}

func TestExample26(t *testing.T) {
	type Button struct {
		Clicked *sync.Cond
	}

	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}
	subscribe := func(c *sync.Cond, fn func()) {
		// 允许我们注册函数处理来自条件的信号，每个处理程序都在自己的 gorountine 上运行
		// 并且订阅不会退出，直到 gorountine 被确认运行为止。
		var gorountineRunning sync.WaitGroup
		gorountineRunning.Add(1)
		go func() {
			gorountineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		gorountineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast()
	// 在 Clicked Cond 调用 Broadcast，所以三个处理程序都将运行

	clickRegistered.Wait()
}

/**
once
sync.Once是一种类型，它在内部使用一些sync原语，以确保即使在不同的goroutine上，也只会调用一次Do方法处理传递进来的函数。
这确实是因为我们将调用sync.Once方式执行Do方法。
sync.Once只计算调用Do方法的次数，而不是多少次唯一调用Do方法。这样，sync.Once的副本与所要调用的函数紧密耦合， 我们再次
看到如何在一个严格的范围内合理使用sync包中的类型以发挥最佳效果。
我建议你通过将sync.Once包装在一个小的语法块中来形式化这种耦合：要么是一个小函数，要么是将两者包装在一个结构体中。
*/
func TestExample27(t *testing.T) {

}

func TestExample28(t *testing.T) {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}
	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()
}

func TestExample29(t *testing.T) {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	// 用 4kb 初始化 pool
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorks = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorks)
	for i := numWorks; i > 0; i-- {
		go func() {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}

func connectToService() interface{} {
	time.Sleep(time.Second)
	return struct{}{}
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()
		wg.Done()
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			connectToService()
			fmt.Fprintln(conn, "")
			conn.Close()
		}
	}()
	return &wg
}

func TestExample210(t *testing.T) {

}