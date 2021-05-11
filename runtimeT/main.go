package main

import (
	"fmt"
	"runtime"
	"time"
)

// runtime.Gosched()  // 让出时间切片

// func main() {
// 	go func(s string) {
// 		for i := 0; i < 2; i++ {
// 			fmt.Println(s)
// 		}
// 	}("world")
// 	// 主协程
// 	for i := 0; i < 2; i++ {
// 		// 切一下，再次分配任务
// 		runtime.Gosched()
// 		fmt.Println("hello")
// 	}
// }

// runtime.Goexit() // // 让出时间切片
// func main() {
// 	go func() {
// 		defer fmt.Println("A.defer")
// 		func() {
// 			defer fmt.Println("B.defer")
// 			runtime.Goexit()             // 结束协程
// 			defer fmt.Println("C.defer") // 没有打印
// 			fmt.Println("B")             // 没有打印
// 		}()
// 		fmt.Println("A")
// 	}()
// 	for {

// 	}
// }

// runtime.GOMAXPROCS  Go运行时的调度器
func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(200) // 需要使用多少个OS线程来同时执行Go代码
	go a()
	go b()
	time.Sleep(time.Second)
}
