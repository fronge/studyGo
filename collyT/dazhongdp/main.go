package main

import (
	"fmt"
	"os/exec"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnRequest(func(req *colly.Request) {
		fmt.Println("Visiting", req.URL)
		req.Headers.Set("Cookie", "_lxsdk_cuid=17b3973b7bbc8-0d36bd7f7f4631-113a6054-13c680-17b3973b7bbc8; _lxsdk=17b3973b7bbc8-0d36bd7f7f4631-113a6054-13c680-17b3973b7bbc8; Hm_lvt_602b80cf8079ae6591966cc70a3940e7=1628756491; _hc.v=74c4a267-fd9f-49db-619c-f162a97e7afe.1628756491; _dp.ac.v=8cfa2a11-f0e0-46b3-84f2-6e6221c820d9; dplet=fa5bf2ee9fd5ee1f72d477a100c02008; dper=52a8b9a1f5440b35c79f08fb1f9d4fca8fed4a922783501698d7ad33096023fe5a57c9eb7a62e332621a15a2ae368e45dd06f8a9ec832366f01ac41db361931d646e00d521565563713d0a6e19af5effbc0acdc2e80d8b432e1fb520949534c7; ll=7fd06e815b796be3df069dec7836c3df; ua=dpuser_0645025773; uamo=13614205855; fspop=test; cy=2; cye=beijing; Hm_lpvt_602b80cf8079ae6591966cc70a3940e7=1628756611; _lxsdk_s=17b3973b4e1-309-572-b13%7C%7C59")
		req.Headers.Set("Host", "www.dianping.com")
		req.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("--OnResponse--")
		// fmt.Print(string(r.Body))
	})
	c.OnHTML(`html`, func(h *colly.HTMLElement) {
		cssURL := h.ChildAttrs(`link[type="text/css"]`, "href")
		fmt.Println(cssURL)
	})
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("--OnError--")
		fmt.Println(e.Error())
	})
	c.Visit("http://www.dianping.com/shop/k4unEzHLSriL5STi")
	// c.Visit("http://www.dianping.com/search/keyword/2/0_%E6%8A%98%E6%89%A3%E5%BA%97/p3")
	c.Wait()
}

func Ziti(url string) {
	cmd := exec.Command("font", "7cc1dac6.woff")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return
	}

}
