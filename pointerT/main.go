package main

import "fmt"

// 不能对指针进行操作
// & 取内存地址
// * 根据地址取值

// make 和 new 的区别
// new 很少用，一半给基本数据类型申请内存 string/int 返回的是对应类型的指针 （*string, *int）
// make 用来分配内存， 给 map，slice，chan申请内存的，make返回的是对应的这三个类型的本身

func newPoint() {
	var a2 = new(int) // new 申请一个内存地址
	fmt.Println(a2)
	*a2 = 100
	fmt.Println(*a2) // 100
}
func testPoint() {
	n := 10
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	fmt.Println(*p)

}

func testKong() {
	// 空指针  没有内存地址
	var a *int
	var i int = 100
	*a = 100
	//会报错的
	fmt.Println(i)

}

// 用指针更改一个值
func testNum() {
	var num int = 9
	var ptr *int
	ptr = &num
	*ptr = 10
	fmt.Println(num)
}

func main() {
	testKong()
}
