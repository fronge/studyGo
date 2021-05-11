package main

import "fmt"

type Animal struct{
	Name string `json "name"`
	Sex int `json "name"`
}

type People struct {
	*Animal	
}
func main()  {
	p:=People{&Animal{Name:"123",Sex:1}}
	fmt.Println(p.name)
}