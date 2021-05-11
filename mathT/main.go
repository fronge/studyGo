package main

import (
	"fmt"
)

func Mod() {
	m := 102 % 100
	fmt.Println(m)
}

func Wei() {
	var e int = 0
	var a int = 1
	var b int = 2
	var c int = 3
	fmt.Println(b | e)
	fmt.Println(a | a)
	fmt.Println(a | b)
	fmt.Println(a | c)
}

func main() {
	Wei()
}
