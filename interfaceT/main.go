package main

import "fmt"

// 接口类型
// 接口类型是对其他类型行为的抽象和概况
/*
定义:
type interface_name interface {
	方法1(参数...)(返回值...)
	方法2(参数...)(返回值...)
	...
}
用来给变量、
一个变量如果实现了接口中规定的 所有 方法，那么这个变量就是实现了这个接口

1.接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
2.方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
3.参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。
*/

type cat struct {
	name string
}

type dog struct{}

//
type speaker interface {
	speak() // 只要实现了speak 方法的变量都是speaer类型
}

// 实现方法
func (c cat) speak() {
	fmt.Println("喵喵~")
}

func (d dog) speak() {
	fmt.Println("汪汪~")
}

func (c cat) say() {
	fmt.Println("猫说~")
}

func (c cat) move() {
	fmt.Println("猫走~")
}

func (d dog) say() {
	fmt.Println("狗说~")
}

func (c dog) move() {
	fmt.Println("狗走~")
}

//
func jiao(x speaker) {
	x.speak()
}

func gg(x animal) {
	x.move()
	x.say()
}

// 接口嵌套
type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type animal interface {
	Sayer
	Mover
}

// func main() {
// 	// var s speaker
// 	// s = cat{}
// 	// jiao(s)
// 	// s = dog{}
// 	// jiao(s)

// 	var an animal
// 	an = cat{name: "hahah"}
// 	an.move()
// 	an = dog{}
// 	an.move()

// }

// 空接口 & 断言
// 空接口没必要起名字
// 特点：
// 所有类型都实现了空接口，也就是任何类型的变量都可以传递进来
// interface 关键字
// interface{} 接口
func justifyT(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Printf("%v:==int==", v)
	case string:
		fmt.Printf("%v:==string==", v)
	case bool:
		fmt.Printf("%v:==bool==", v)
	default:
		fmt.Printf("%v:==default==", v)
	}
}

func main() {
	// 使用空接口实现可以保存任意值的字典
	var m1 map[string]interface{} // 空接口作为map值
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "哈哈"
	m1["age"] = 11
	// 断言 判断interface值类型 使用 x.(T) 其中 T为类型 x为变量 只能用在interfa

	justifyT(m1["age"])
	// fmt.Println(m1["name"].(int)) //  error

}
