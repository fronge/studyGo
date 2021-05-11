package main

import "fmt"

const (
	U int = iota
	K
	M = iota + 3
)

func main() {
	fmt.Println(U)
	fmt.Println(K)
	fmt.Println(M)
}
