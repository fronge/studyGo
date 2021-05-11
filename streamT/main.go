package main

import (
	"context"
	"fmt"
	"log"
	pd "studyGo/streamT/stream"

	"google.golang.org/grpc"
)

type Streammsg struct {
	Text string
	Code int
}

func main() {
	conn, err := grpc.Dial("127.0.0.1:50001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pd.NewStreamServiceClient(conn)
	r, err := c.SimpleFun(context.Background(), &pd.RequestData{Text: "111"})
	fmt.Println(r)
	fmt.Println(err)
}
