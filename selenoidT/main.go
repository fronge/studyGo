package main

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
)

func main() {
	capabilities := selenium.Capabilities{
		"browserName": "firefox",
	}
	capabilities["browserName"] = "firefox"

	capabilities["enableVNC"] = true
	capabilities["screenResolution"] = "1280x1024x24"

	driver, err := selenium.NewRemote(capabilities, "http://localhost:4444/wd/hub")
	if err != nil {
		fmt.Println("----")
		fmt.Println(err)
		return
	}
	for {
		driver.Get("https://www.baidu.com")
		fmt.Println("======")
		time.Sleep(time.Minute)
	}

}
