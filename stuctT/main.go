package main

import "fmt"

type Author struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Post struct {
	ID        int64      `json:"id"`
	Content   string     `json:"content"`
	Author    Author     `json:"author"`
	Published bool       `json:"published"`
	Label     []string   `json:"label"`
	NextPost  *Post      `json:"nextPost"`
	Comments  []*Comment `json:"comments"`
}

// func main() {
// 	// 打开json文件
// 	fh, err := os.Open("a.json")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer fh.Close()
// 	// 读取json文件，保存到jsonData中
// 	jsonData, err := ioutil.ReadAll(fh)
// 	fmt.Println(jsonData)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	var post Post
// 	// 解析json数据到post中
// 	err = json.Unmarshal(jsonData, &post)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(post)
// }

// func main() {
// 	post := &Post{}
// 	author := Author{
// 		0,
// 		"Name",
// 	}
// 	// author.ID = 0
// 	author.Name = "Name"

// 	comment := &Comment{
// 		0,
// 		"content comment",
// 		"author comment",
// 	}

// 	postp := Post{
// 		0,
// 		"content post",
// 		author,
// 		true,
// 		[]string{"linux", "shell"},
// 		post,
// 		[]*Comment{comment},
// 	}

// 	p, _ := json.MarshalIndent(postp, "", "\t")
// 	fmt.Println(string(p))
// }

type Animal struct {
	name string
}

func (a *Animal) move() {
	a.name = "10101"
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet   int8
	Animal //通过嵌套匿名结构体实现继承
}

func main() {
	a := Animal{"aa"}
	d := Dog{4, a}
	a.move()
	// a.name = "bb"
	fmt.Println(d.name)

}
