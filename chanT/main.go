package main

import (
	"fmt"
	"math/rand"
)

// channel 实现多个goroutine

// 1、Go语言的并发模型是CSP (Communicating Sequential Processes）
// 2、Go多并发时 通过通信共享内存而不是通过共享内存而实现通信

// goroutine是Go程序并发的执行体，channel就是它们之间的连接

// channel 特点
// 1、先入先出
// 2、每一个通道都是一个具体类型的导管
// 3、channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制
// 4、是一种引用类型
// 5、空值是nil

// channel 声明
// 1、var ch chan int // 声明一个传递int类型的channel
// 2、make(chan 元素类型, [缓冲大小])  // 缓冲大小是可选

// 操作
// 通道有发送（send）、接收(receive）和关闭（close）三种操作
// 1、发送和接收都使用<-符号

// 使用实例
// 1、无缓冲通道 又叫 同步通道
// 无缓冲的通道只有在有人接收值的时候才能发送值
// func recv(c chan int) {
// 	fmt.Println("3")
// 	ret := <-c
// 	fmt.Println("接收成功", ret)
// }

// func main() {
// 	ch := make(chan int)
// 	fmt.Println("1")
// 	go recv(ch)
// 	fmt.Println("2")
// 	ch <- 10
// 	fmt.Println("4")
// 	fmt.Println("发送成功")
// 	// close(ch)
// }

// 2、有缓冲通道
// 只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量
// 使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量 ==> 很少用
// 内置的close()函数关闭channel
// func main() {
// 	ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
// 	ch <- 10
// 	fmt.Println("发送成功")
// 	x := <-ch
// 	fmt.Println(x)
// }

// 优雅的循环取值
// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
// 	// 开启goroutine将0~100的数发送到ch1中
// 	go func() {
// 		for i := 0; i < 100; i++ {
// 			ch1 <- i
// 		}
// 		close(ch1)
// 	}()
// 	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
// 	go func() {
// 		for {
// 			i, ok := <-ch1 // 通道关闭后再取值ok=false
// 			if !ok {
// 				break
// 			}
// 			ch2 <- i * i
// 		}
// 		close(ch2)
// 	}()
// 	// 在主goroutine中从ch2中接收值打印
// 	for i := range ch2 { // 通道关闭后会退出for range循环
// 		fmt.Println(i)
// 	}
// }

// 三、单向通道
// 1.chan<- int是一个只能发送的通道，可以发送但是不能接收；
// 2.<-chan int是一个只能接收的通道，可以接收但是不能发送。
// func counter(out chan<- int) {
// 	for i := 0; i < 100; i++ {
// 		out <- i
// 	}
// 	close(out)
// }

// // out chan<- int只能投放， in <-chan int 只能取
// func squarer(out chan<- int, in <-chan int) {
// 	for i := range in {
// 		out <- i * i
// 	}
// 	close(out)
// }
// func printer(in <-chan int) {
// 	for i := range in {
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
// 	go counter(ch1)
// 	go squarer(ch2, ch1)
// 	printer(ch2)
// }

//  多个任务 worker pool（goroutine池）
type Job struct {
	// id
	Id int
	// 需要计算的随机数
	RandNum int
}

type Result struct {
	// 这里必须传对象实例
	job *Job
	// 求和
	sum int
}

func main() {
	// 需要2个管道
	// 1.job管道
	jobChan := make(chan *Job, 128)
	// 2.结果管道
	resultChan := make(chan *Result, 128)
	// 3.创建工作池
	createPool(3, jobChan, resultChan)
	// 4.开个打印的协程
	go func(resultChan chan *Result) {
		// 遍历结果管道打印
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id,
				result.job.RandNum, result.sum)
		}
	}(resultChan)
	var id int
	// 循环创建job，输入到管道
	for {
		id++
		// 生成随机数
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}
}

// 创建工作池
// 参数1：开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range jobChan {
				// 随机数接过来
				r_num := job.RandNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				// 想要的结果是Result
				r := &Result{
					job: job,
					sum: sum,
				}
				//运算结果扔到管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}

// 备注
// 操作已经关闭的通道会引发 panic
// 关闭、发送、接收都会引发 panic
