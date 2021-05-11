package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go signalHandle() //用go程执行信号量处理函数
	fmt.Println("======")
	wg.Wait()
}

func signalHandle() {
	for {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGHUP)
		sig := <-ch

		fmt.Printf("Signal received: %v", sig)
		switch sig {
		default:
			fmt.Printf("get sig=%v\n", sig)
			fmt.Println("1231312")
		case syscall.SIGHUP:
			fmt.Println("get sighup\n") //Utils.LogInfo是我自己封装的输出信息函数
		case syscall.SIGINT:
			fmt.Println("\n----88---")
			os.Exit(1)
		case syscall.SIGQUIT:
			fmt.Println("SIGQUIT")
		case syscall.SIGUSR1:
			fmt.Println("SIGUSR1\n")
		case syscall.SIGUSR2:
			fmt.Println("SIGUSR2\n")
		}
		fmt.Println("1231312====")
	}
}
