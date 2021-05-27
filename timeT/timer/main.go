package main

import(
	"fmt"
	"time"
	"math/rand"
)

func AfterFunc() {
	fmt.Println("AfterFunc调用时间:",time.Now())
}

func main(){
	fmt.Println("开始时间:", time.Now())
	t2 := time.AfterFunc(time.Duration(1)*time.Second, AfterFunc) // 4秒之后另起一个go执行AfterFunc
	fmt.Println(t2.Reset(time.Duration(10)*time.Second))
	
	// 只有一次
	t := time.NewTimer(time.Second)
	for{
		select{
			case <-t.C:
				fmt.Println("T1 到时了:%v",time.Now())
			// default:
			// 	fmt.Println("----")
			// 	time.Sleep(time.Second)
		}
		
		randTime := rand.Intn(5)
		fmt.Println("随机时间:",randTime)
		// 重置
		
		fmt.Println("外面见~",t.Reset(time.Duration(randTime)*time.Second))
	}
}