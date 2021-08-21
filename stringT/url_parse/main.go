package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	u := "?a=1&b=3"
	Url, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	a := Url.Query().Get("a")
	// a := Parse(u)
	fmt.Println(a)
}

func Parse(str string) map[string]string {
	var m = make(map[string]string, 1)
	strs := strings.Split(str, "&")
	fmt.Println(strs)
	for _, s := range strs {
		sr := strings.Split(s, "=")
		fmt.Println(sr)
		if len(sr) > 1 {
			m[sr[0]] = sr[1]
		}
	}
	return m
}
