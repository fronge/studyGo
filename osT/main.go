package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ReadFromBufIo() {
	// bufio
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)
	for {

		line, err := reader.ReadString('\n') // 字符用单引号
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}

		fmt.Print(line)
	}
}

func readFromFileByIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("error:%v", err)
	}
	fmt.Print(string(ret))

}

func main() {
	// 简单的读取文件
	// fileObj, err := os.Open("./main.go")
	// if err != nil {
	// 	fmt.Printf("open file falled, err:%v", err)
	// }
	// defer fileObj.Close()
	// // 读文件
	// var tmp [128]byte
	// for {
	// 	n, err := fileObj.Read(tmp[:])
	// 	if err == io.EOF {
	// 		fmt.Println("读取完成...")
	// 		return
	// 	}
	// 	if err != nil {
	// 		fmt.Printf("error:%v", err)
	// 		return
	// 	}
	// 	fmt.Println(string(tmp[:n]))
	// }

	// ReadFromBufIo()
	readFromFileByIoutil()

}
