package main

import "fmt"

func main() {
	// 在if的时候创建的变量只能在语句块中使用
	// if n := 3; n > 3 {
	// 	fmt.Printf("%v>3\n", n)
	// } else if n < 3 {
	// 	fmt.Printf("%v<3\n", n)
	// } else {
	// 	fmt.Printf("%v=3\n", n)
	// }

	// // switch
	// a := 1
	// switch a {
	// case 0:
	// 	fmt.Printf("a:%v=0\n", a)
	// case 1:
	// 	fmt.Printf("a:%v=1\n", a)
	// case 2:
	// 	fmt.Printf("a:%v=0\n", a)
	// default:
	// 	fmt.Printf("a:%v is default!\n", a)
	// }

	// // type switch
	// var x interface{}
	// switch i := x.(type) {
	// case nil:
	// 	fmt.Printf("x type is %T", i)
	// case int:
	// 	fmt.Printf("x type is int")
	// case func(int) float64:
	// 	fmt.Println("x type is func(int) float64")
	// case bool, string:
	// 	fmt.Printf("x type is bool or string")
	// default:
	// 	fmt.Println("未知类型")
	// }

	// c := 10
	// switch {
	// case c > 0 && c < 5:
	// 	fmt.Println("0<c<5")
	// case c == 5:
	// 	fmt.Println("c=5")
	// default:
	// 	fmt.Println("default:c>5")
	// }

	// select
	/*
			每个case都必须是一个通信
		    所有channel表达式都会被求值
		    所有被发送的表达式都会被求值
		    如果任意某个通信可以进行，它就执行；其他被忽略。
		    如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
		    否则：
		    如果有default子句，则执行该语句。
			如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
	*/
	// 使用场景
	// 用于处理异步IO操作
	// select会监听case语句中channel的读写操作，当case中channel读写操作为非阻塞状态（即能读写）时，将会触发相应的动作。
	// select中的case语句必须是一个channel操作

	// select 写法
	// 	select {
	//     case communication clause  :
	//        statement(s);
	//     case communication clause  :
	//        statement(s);
	//     /* 你可以定义任意数量的 case */
	//     default : /* 可选 */
	//        statement(s);
	// }

	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received", i1, " from c1\n ")
	case c2 <- i2:
		fmt.Printf("send", i2, " to c2\n ")
	case i3, ok := (<-c3):
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Println("c3 is close")
		}
	default:
		fmt.Println("no communication\n")
	}

}
