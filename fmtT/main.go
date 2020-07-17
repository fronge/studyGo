package main

import (
	"fmt"
	"time"
)

type Test struct {
	Name string
	Age  int
}

func test() {
	d := time.Now()
	ss := fmt.Sprintf("aa:%v", d)
	fmt.Println(ss, "\n")
	fmt.Println(d)
}
func main() {
	test()
	// a := Test{Name: "哈哈", Age: 10}
	// fmt.Printf("%T", a) // type
	// fmt.Printf("%v", a) // 值得默认表示 value
	// fmt.Printf("%v", a) // 值得默认表示 value 输出结构体会添加字段名称

	// b := true
	// fmt.Printf("bool:%t", b) // 布尔值

	// 关于整数
	// c := 10000
	// fmt.Printf("%b", c) // 二进制
	// fmt.Printf("%c", c) // 数字对应的unicode 码的值
	// fmt.Printf("%d", c) // 10进制的值
	// fmt.Printf("%o", c) // 8进制的值
	// fmt.Printf("%x", c) // 16进制的值 用a-f表示
	// fmt.Printf("%X", c) // 16进制的值 用A-F表示
	// fmt.Printf("%X", c) // 16进制的值 用A-F表示
	// fmt.Printf("%U", c) // 表示为Unicode格式 U+03E8
	// d := 2
	// fmt.Printf("%q", c) // 单引号的字面值 '\x02'

	// 浮点型
	// d := 1.11
	// fmt.Printf("%b", d) // 无小数部分、二进制指数的科学计数法 4998995586381251p-52
	// fmt.Printf("%e", d) // 科学计数法 1.110000e+00
	// fmt.Printf("%E", d) // 科学计数法 11.110000E+00
	// fmt.Printf("%f", d) // 有小数部分但无指数部分 1.110000
	// fmt.Printf("%F", d) // ==%f 1.110000
	// fmt.Printf("%g", d) // 根据实际情况采用%e或%f格式 1.11
	// fmt.Printf("%G", d) // 根据实际情况采用%E或%F格式 1.11

	// 字符串和[]byte
	// e := "哈哈哈"
	// fmt.Printf("%s", e) // 直接输出字符串或者[]byte
	// fmt.Printf("%q", e) // 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
	// fmt.Printf("%x", e) // 每个字节用两字符十六进制数表示（使用a-f）
	// fmt.Printf("%X", e) // 每个字节用两字符十六进制数表示（使用A-F）

	// 指针
	// f := "哈哈"
	// g := &f
	// fmt.Printf("%p", g) // 0xc00008e1e0

	// h := 12.2345 // 宽度:设置输出总长度  精度：控制小数部分
	// fmt.Printf("%f", h)  // 默认宽度 默认精度
	// fmt.Printf("%9f", h)  // 宽度9 默认精度
	// fmt.Printf("%3.2f", h) // 宽度3 精度2

	fmt.Println("\n")
}
