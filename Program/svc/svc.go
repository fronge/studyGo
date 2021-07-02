package svc

import (
	"os"
	"os/signal"
)

type Service interface{
	Init() error
	Start() error
	Stop() error
}

var msgChan = make(chan os.Signal, 1)

func Run(service Service) error {
	if err := service.Init(); err != nil{
		return err
	}
	if err := service.Start(); err != nil{
		return err
	}
	// 监听control + c
	signal.Notify(msgChan, os.Interrupt, os.Kill)
	<-msgChan
	return service.Stop()
}

func Interrupt(){
    msgChan<-os.Interrupt
}