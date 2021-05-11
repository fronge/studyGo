package main

import (
	"fmt"
	"time"
)

// select && channal
func main() {
	ch := make(chan int, 4)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		default:
			fmt.Println(0)
		}

	}
	time.Sleep(5 * time.Second)
}
