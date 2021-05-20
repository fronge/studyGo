package main

import (
	"fmt"
	"context"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, cannel chan bool){
	defer wg.Done()
	for{
		select{
		default:
			fmt.Println("hello")
		case <-cannel:
			return
		}
	}
}

func workerTwo(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	for{
		select{
		case<-ctx.Done():
			fmt.Println("===EXIT===")
			return ctx.Err()
		default:
			fmt.Println("hello")
			time.Sleep(1*time.Second)
		}
	}
}


func main(){
	// cancel := make(chan bool)
	// var wg sync.WaitGroup
	// for i:=0;i<10;i++{
	// 	wg.Add(1)
	// 	go worker(&wg,cannel)
	// }
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	var wg sync.WaitGroup
	for i:=0; i<10;i++{
		wg.Add(1)
		go workerTwo(ctx,&wg)

	}
	time.Sleep(time.Second)
	cancel()
	wg.Wait()
}