package main

import (
	"errors"
	"fmt"
)

// 错误哲学 每一个错误都是一个值
// defer + recover
// recover 内置函数可以捕获到异常
func test() {
	defer func() {
		err := recover()
		if err != nil { // 捕获到异常
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 2
	res := num1 / num2
	fmt.Println("res=", res)
}

// 自定义异常
// errors.New("error info") 会返回一个error类型的值 但是并不是直接抛出
// panic 内置函数 接受一个interface{}类型的值， 可以接受一个error类型变量，输出错误信息，并退出程序
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("read file error...\n") //
	}
}

func test02() {
	defer func() {
		err := recover() // 获取到错误
		if err != nil {
			fmt.Println("error:", err)
		}
	}()
	err := readConf("config2.ini")
	fmt.Println(err)
	if err != nil {
		panic(err) // 错误抛出
	}
	fmt.Println("test02() go on...")
}

func main() {
	test02()
}
