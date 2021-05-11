package main

// range
// 类似于迭代器 可以遍历数组、字符串、map等等
// range会复制对象，而不是不是直接在原对象上操作。 浅拷贝
// 如果数据部分是一个指针，指向地址，复制对象的时候只是把指针的值复制，如果修改的话会修改原值

func main() {
	// 遍历数组
	// a := [3]int{1, 2, 3}
	// for i, n := range a {
	// 	fmt.Println(i, n)
	// }

	// 切片
	// 	b := []int{2, 3, 4}
	// 	for i, n := range b {
	// 		fmt.Println(i, n)
	// 	}

	//map 遍历
	// c := map[string]string{"Hello": "world"}
	// for k, v := range c {
	// 	fmt.Println(k, v)
	// }

}
