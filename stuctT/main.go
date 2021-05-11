package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

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

type T struct {
	Version int
	Content string
}

func makeT(c string, v int) *T {
	if v == 0 {
		v = 1
	}
	return &T{
		Content: c,
		Version: v,
	}
}

func main() {
	// 打开json文件
	fh, err := os.Open("a.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fh.Close()
	// 读取json文件，保存到jsonData中
	jsonData, err := ioutil.ReadAll(fh)
	fmt.Println(jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}

	var post Post
	// 解析json数据到post中
	err = json.Unmarshal(jsonData, &post)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(post)
}
