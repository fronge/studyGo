package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 并发goroutine
// 使用goroutine:在调用函数前面加一个go

// func hello() {
// 	fmt.Println("Hello Goroutine!")
// }

// func main() {
// 	go hello()
// 	fmt.Println("main goroutine done!")
// 	time.Sleep(time.Second) // 等待创建 否则主线程直接关闭了
// }

// 多个goroutine
// var wg sync.WaitGroup //

// func hello(i int) {
// 	defer wg.Done() // goroutine结束就登记-1
// 	fmt.Println("Hello Goroutine!", i)
// }

// func main() {
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1) // 启动一个goroutine就登记+1
// 		go hello(i)
// 	}
// 	wg.Wait() // 等待所有等级的goroutine都结束
// }

// 如果主协程退出了，其他任务也结束
// func main() {
// 	// 合起来写
// 	go func() {
// 		i := 0
// 		for {
// 			i++
// 			fmt.Printf("new goroutine: i = %d\n", i)
// 			time.Sleep(time.Second)
// 		}
// 	}()
// 	i := 0
// 	for {
// 		i++
// 		fmt.Printf("main goroutine: i = %d\n", i)
// 		time.Sleep(time.Second)
// 		if i == 1 {
// 			break
// 		}
// 	}
// }

// goroutine 可增长的栈
// goroutine 与线程的区别
// 一个goroutine的栈在其生命周期开始时只有很小的栈（典型情况下2KB），goroutine的栈不是固定的，他可以按需增大和缩小，goroutine的栈大小限制可以达到1GB

// goroutine 调度
// GMP调度
// G：goroutine
// 就是个goroutine的，里面除了存放本goroutine信息外 还有与所在P的绑定等信息。
// P:管理者
// P管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界），P会对自己管理的goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停、运行后续的goroutine等等）当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了会去其他P的队列里抢任务。

// M: machine
// M（machine）是Go运行时（runtime）对操作系统内核线程的虚拟， M与内核线程一般是一一映射的关系， 一个groutine最终是要放到M上执行的；

// 所有的调度都是有GO语言层面实现的，由P调度

// runtime.GOMAXPROCS 用来设置CPU

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%v\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%v\n", i)
	}
}

func main() {
	runtime.GOMAXPROCS(4)
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
