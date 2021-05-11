package main

import "fmt"

// 函数类型

// 匿名函数
// 一次执行函数

func f1() {
	fmt.Println("HELLO")
}

func f2() int {
	return 1
}

// 参数要求，返回值要求
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

// 不能传给f3
func f4(x, y int) int {
	return x + y
}

// 函数当成参数，函数当成返回值
func f5(x func() int) func(int, int) int {
	ret := func(a, b int) int {
		return a + b
	}
	return ret
}

func f6(x func() int) func(string, string) int {
	return func(name, age string) int {
		return 1
	}
}

func main() {

}
