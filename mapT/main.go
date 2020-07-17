package main

import (
	"fmt"
)

func main() {
	// var m1 map[string]string  //
	// fmt.Println(m1 == nil) // 未初始化 没有在内存中开劈空间
	// m1["name"] = "费时"  // 报错

	// var m2 map[string]int
	m2 := make(map[string]int, 10)
	m2["name"] = 10
	m2["age"] = 12
	fmt.Println(m2) // map[age:12 name:10]

	// 取值
	// v, ok := m2["哈哈"]
	// if !ok {
	// 	fmt.Println("查无此K")
	// } else {
	// 	fmt.Println(v)
	// }

	// 遍历key
	// for k := range m2 {
	// 	fmt.Println(k)
	// }

	// 遍历map
	// for k, v := range m2 {
	// 	fmt.Println(v) // value
	// 	fmt.Println(k) // key
	// }

	// 删除
	// delete(m2, "name")
	// fmt.Println(m2) // map[age:12]

	//  // // 按照顺序遍历
	// var scoreMap = make(map[string]int, 200)
	// for i := 0; i < 100; i++ {
	// 	key := fmt.Sprintf("stuf%02d", rand.Intn(100)) // 生成stu开头的字符串
	// 	value := rand.Intn(100)
	// 	scoreMap[key] = value
	// }
	// fmt.Println(scoreMap)
	// var keys = make([]string, 0, 200)
	// for key := range scoreMap {
	// 	keys = append(keys, key)
	// }
	// sort.Strings(keys)  // 将keys 排序
	// for _, key := range keys {
	// 	fmt.Println(key, scoreMap[key])
	// }

}
