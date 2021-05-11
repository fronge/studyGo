package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func(n int) {
		fmt.Println("n:", n)
		t := time.Duration(n) * time.Second
		time.Sleep(t)
		fmt.Println("-:", n)
		wg.Done()
	}(1)
	wg.Add(1)
	go func(n int) {
		fmt.Println("n:", n)
		t := time.Duration(n) * time.Second
		time.Sleep(t)
		fmt.Println("-:", n)
		wg.Done()
	}(2)
	wg.Add(1)
	go func(n int) {
		fmt.Println("n:", n)
		t := time.Duration(n) * time.Second
		time.Sleep(t)
		fmt.Println("-:", n)
		wg.Done()
	}(3)

	wg.Wait()

	fmt.Println("main exit...")
}
