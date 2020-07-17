package main

import "fmt"

// 自定义类型
type myInt int

func main() {
	m := myInt(10)
	fmt.Println(m)
}
