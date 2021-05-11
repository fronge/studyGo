package main

import "fmt"

func test() {
	defer func() {
		fmt.Println("----------")
	}()
	s := []int{1, 2, 3, 4}
	for i := 0; i < 10; i++ {
		for _, j := range s {
			fmt.Println(j)
			return
		}
	}
}

func main() {
	test()
}
