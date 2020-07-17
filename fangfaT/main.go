package main

import (
	"fmt"
)

// 方法
// 语法格式
// func (variable_name variable_data_type) function_name() [return_type]{
// 	/* 函数体*/
//  }

type dog struct {
	Name string
	Age  int
}

// 构造函数
func newDog(name string) dog {
	return dog{
		Name: name,
	}
}

// 方法是作用于特定类型的函数
// (d Dog) =>d 方法的接收器(receiver)，相当于python 中的self，在go中多用构造体的首字母的小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪~", d.Name)
}

// 接受者：值接收者，指针接收者 一般都是用指针接收者
// 指针接收者使用场景:
// 1、需要修改接收者中的值
// 2、接收者的拷贝代价比较大的大对象
// 3、保证一致性，如果有个别的方法使用了指针接收者，那么其他的也使用指针接收者
type persen struct {
	Name string
	Age  int
}

func newPersen(name string, age int) persen {
	return persen{
		Name: name,
		Age:  age,
	}
}

// 值接收者
func (p persen) guonian() {
	p.Age++
}

// 指针接收者
func (p *persen) zhenguonian() {
	p.Age++
}

func (p *persen) dream() {
	fmt.Printf("%v拯救世界", p.Name)
}

// // ###########################
// 自定义类型添加方法
// 不能给别的包类型添加方法，只能给自己包类型定义方法
type myInt int

func (m myInt) hello() {
	fmt.Println("hello")
}

func main() {
	// d1 := newDog("大熊")
	// d1.wang()

	// p := newPersen("人", 20)
	// p.guonian()
	// fmt.Printf("%v", p.Age) // 20
	// p.zhenguonian()
	// fmt.Printf("%v", p.Age) // 21
	// p.dream()

	m := myInt(100)
	m.hello()

}
