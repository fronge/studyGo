package main

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.AllowURLRevisit = true
	once := sync.Once{}
	for i := 0; i < 10; i++ {
		once.Do(func() {
			c.OnRequest(func(_ *colly.Request) {
				fmt.Println(i)
			})
			// c.OnHTML(`html`, func(h *colly.HTMLElement) {
			// 	fmt.Println(i)
			// })
		})
		c.Visit("http://www.baidu.com")
	}
}
