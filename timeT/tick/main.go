package main

import(
	"fmt"
	"time"
)
// 会调整时间间隔或者丢弃tick信息以适应反应慢的接收者。

func main(){
	t := time.NewTicker(time.Second)
	defer t.Stop()
	for{
		select{
			case <-t.C:
				fmt.Println("到时了:%v",time.Now())
		}
		time.Sleep(time.Duration(3)*time.Second)
		fmt.Println("外面见~")
	}
}