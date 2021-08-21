package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type T struct {
	Content string `json:"content"`
}

func UnescapeUnicodeCharactersInJSON(jsonRaw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(jsonRaw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func main() {
	t := T{Content: "<>"}
	b, _ := json.Marshal(t)
	fmt.Println(string(b))
	c, _ := UnescapeUnicodeCharactersInJSON(b)
	fmt.Println(string(c))

}
