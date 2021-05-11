package main

import (
	"bytes"
	"strconv"
)

func cf(a string) (s string) {
	var e bytes.Buffer
	r, _ := strconv.ParseInt(a[0:2], 16, 0)
	for n := 4; n < len(a)+2; n += 2 {
		i, _ := strconv.ParseInt(a[n-2:n], 16, 0)
		e.WriteString(string(i ^ r))
	}
	return e.String()
}

func main() {
	email := cf("5b3a6c686d6b38381b6a6d6875383436")
	print(email)
	print("\n")
}
