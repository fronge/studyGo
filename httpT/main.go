package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	Name string
	Age  int
}

type Ret struct {
	Code  int
	Param string
	Msg   string
	Data  []Data
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	fmt.Println(m)
	data := Data{Name: "why", Age: 18}
	ret := new(Ret)
	id := r.FormValue("id")

	ret.Code = 0
	ret.Param = id
	ret.Data = append(ret.Data, data)
	ret.Data = append(ret.Data, data)
	ret.Data = append(ret.Data, data)
	ret_json, _ := json.Marshal(ret)

	io.WriteString(w, string(ret_json))

}

func HeServer(w http.ResponseWriter, r *http.Request) {
	r.Form
	io.WriteString(w, "hello world")
}
func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<div color=\"#bfa\">1111</div> <input>`)
}

func main() {
	http.HandleFunc("/data/t", HelloServer)
	http.HandleFunc("/hello", HeServer)
	http.HandleFunc("/", Index)

	fmt.Println("----start server http://0.0.0.0:8800 ------")
	err := http.ListenAndServe("0.0.0.0:8800", nil)
	if err != nil {
		fmt.Println("htpp listen failed")
	}
}
