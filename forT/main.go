package main

import "fmt"

func test() {
	for i := 1; i <= 0; i++ {
		fmt.Print(i, "\n")
	}
}

func main() {
	test()
}
