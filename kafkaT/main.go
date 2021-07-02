package main

import (
	"fmt"
	"studyGo/kafkaT/kafka"
)

func init() {
	kafka.InitKafka()
}
func main() {
	// kafka.InitKafka()
	p, o, err := kafka.SendToKfK("topic_test", "12313")
	if err != nil {
		fmt.Println("error to kfk")
	}
	fmt.Println("send kfk partition:%d,offset:%d", p, o)
}
