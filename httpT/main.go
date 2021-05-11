package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func getD(w http.ResponseWriter, r *http.Request) {
	emailIDs, ok := r.URL.Query()["id"]
	if !ok {
		io.WriteString(w, "查库错误")
	}
	fmt.Println(emailIDs)
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

func HelloServer(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	fmt.Println(m)
	if strings.ToUpper(m) == "GET" {
		getD(w, r)
	}
	if strings.ToUpper(m) == "POST" {
		io.WriteString(w, "POST")
	}

	if strings.ToUpper(m) == "PUT" {
		if r.Form == nil {
			r.ParseMultipartForm(32 << 20)
		}
		for key, value := range r.Form {
			if key == "userName" {
				fmt.Println("userName:", value)
			} else if key == "passworld" {
				fmt.Println("passworld:", value)
			} else if key == "server" {
				fmt.Println("server:", value)
			}
		}
		io.WriteString(w, "PUT")
	}

	if strings.ToUpper(m) == "DELETE" {
		io.WriteString(w, "DELETED")
	}

}

func HeServer(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]
	if ok {
		fmt.Println(keys, "======")
	}

	io.WriteString(w, "=====")
}
func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<div color=\"#bfa\">1111</div> <input>`)
}

func main() {
	http.HandleFunc("/data", HelloServer)
	http.HandleFunc("/hello", HeServer)
	http.HandleFunc("/", Index)

	fmt.Println("----start server http://0.0.0.0:8800 ------")
	err := http.ListenAndServe("0.0.0.0:8800", nil)
	if err != nil {
		fmt.Println("htpp listen failed")
	}
}
