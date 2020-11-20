package main

import (
	"fmt"
	"net/http"
)

func trM() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "http://www.lagou.com/utrack/trackMid.html", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("referer", "https://www.lagou.com/")
	request.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/981.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36")

	response, _ := client.Do(request)
	defer response.Body.Close()
	//检出结果集
	header := response.Header
	fmt.Println(header)
}
func main() {
	trM()
}
