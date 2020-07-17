package main

import (
	"fmt"
	"time"
)

func main() {
	sli := make([]int, 0)
	var i = 1
	for {
		i++

		sli = append(sli, i)
		if len(sli) > 4 {
			fmt.Println("----小于10----", sli)
			sli = (sli)[0:0]
		}
		time.Sleep(1 * time.Second)

	}
}
