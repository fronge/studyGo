package main

import "fmt"

// 变量的作用域
// 查找顺序：
//   1 现在函数内部找
//   2 找不到在函数外面找，一直找到全局

// 函数内定义的变量只能在函数内部使用
// 语句块中的作用域
var x = 100 // 全局变量

func f1() {
	fmt.Println(x)
}

func main() {
	f1()
}
