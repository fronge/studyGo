package main

import (
	"fmt"
)

// 函数
// 函数定义 第一个括号为参数， 第二个括号放返回值
// 参数可以命名也可以不命名
// 命名返回值相当于在函数中声明了变量
// 使用命名返回值return后可以省略
// Go语言没有默认参数值，也没有任何方法可以通过参数名指定形参，因此形参和返回值的变量名对于函数调用者而言没有意义。

func sum(x int, y int) (ret int) {
	ret = x + y
	return
}

// 没有参数没有返回值
func f0() {
	fmt.Println("f0")
}

// 有参数无返回值函数
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 无参数有返回值
func f2() int {
	return 1
}

// 多个返回值
func f3() (string, int) {
	return "哈哈", 10
}

// 参数类型简写：参数中连续多个参数的类型一致时，将非最后一个参数的类型省略
func f4(x, y, z int) {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

// 可变参数
func f5(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y是切片 []int
	y = append(y, 12)
	fmt.Println(y)
}

// 可变参数为空接口类型
func f6(a ...interface{}){
	fmt.Println(a...)
}

// defer 会把它后面的语句延迟到函数即将结束的时候执行 多用于函数结束之前释放资源：文件句柄，数据库连接，socket连接etc.
// 最想定义的defer 最后执行 后进先出
// return不是原子操作 分两步：返回值赋值，RET指令
// defer执行时机：return之前，返回值赋值之后 => 返回值赋值 运行defer RET指令
func deferDemo() {
	fmt.Println("START")
	defer fmt.Println("AAAAAAAAAAAA")
	defer fmt.Println("BBBBBBBBBBBB")
	fmt.Println("END")
}

// ========关于defer的题 =============
// 默认选择题
func a() int {
	x := 5
	defer func() {
		x++ // 改的是x 不是返回值
	}()
	return x
}

// 返回值为x
func b() (x int) {
	defer func() {
		x++
	}()
	return 5 //最终return 6
}

func c() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 最终return 5
}

// 函数参数是一个副本
func d() (x int) {
	defer func(x int) {
		x++ // 这个x 不是外面的x
	}(x)
	return 5 // 最终return 5
}

// 最后传入的变变量为 3 所以都输出 3
func deferForOne(){
	for i := 0; i<3;i++{
		defer func(){println(i)}()
	}
}
// 重新声明变量
func deferForTwo(){
	for i := 0; i<3;i++{
		i := i
		defer func(){println(i)}()
	}
}
// 将变量传入
func deferForTree(){
	for i := 0; i<3;i++{
		defer func(i int){println(i)}(i)
	}
}

type A = struct {
	Name string
	Msg  string
}

func TestCArgs(aa *A) {
	aa.Name = "hhhhh"
}
func TestArgs(aa A) {
	aa.Name = "hhhhh"
}

func testArg() {
	aa := A{
		"aaa",
		"bb",
	}

	TestArgs(aa)
	fmt.Println(aa)
	TestCArgs(&aa)
	fmt.Println(aa)

}

func twice(x []int){
	for i:= range x{
		x[i] *=2
	}
}

type IntSliceHeader struct{
	Data []int
	Len int
	Cap int
}

func twiceTwo(x IntSliceHeader){
	for i:=0;i<x.Len;i++{
		x.Data[i]*=2
	}
}
//
type interface IntSliceHeader{}

func main() {
	a := []int{1,2,}
	twice(a)
	fmt.Println(a)
	b := IntSliceHeader{[]int{1,2,},2,2}
	twiceTwo(b)
	fmt.Println(b)

}
