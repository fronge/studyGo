package main

import (
	"log"
	"net/rpc"

	"fmt"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)

	}
	var reply string
	err = client.Call("HelloService.Good", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
