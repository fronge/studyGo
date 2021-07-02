package main

import(
	"fmt"
	"time"
	"studyGo/Program/svc"
)

type Program struct {}

func (p *Program)Start() error {
	fmt.Println("application is begin.")
	go func() {
		for{
			time.Sleep(time.Second)
			fmt.Println("application is running.")
		}
	}()
	return nil
}

func (p *Program)Init() error {
	fmt.Println("application is init.")
	return nil
}

func (p *Program)Stop() error {
	fmt.Println("application is end.")
	return nil
}


func main() {
	p := &Program{}
	svc.Run(p)
}