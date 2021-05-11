package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func readFile() {
	b, err := ioutil.ReadFile("./j.json")
	if err != nil {
		fmt.Println(err)
	}
	whitelist := map[string]map[string]int{}
	err = json.Unmarshal(b, &whitelist)
	if err != nil {
		fmt.Println(err)
	}

	for key, value := range whitelist {
		fmt.Println("key:", key, "code:", value["code"], "hot:", value["hot"])
	}
}

func main() {
	readFile()
}
