package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnRequest(func(req *colly.Request) {
		fmt.Println("Visiting", req.URL)
		city := req.URL.Query().Get("city")
		fmt.Println("Visiting", city)
		req.Headers.Set("Cookie", "uuid=2e7283fe6e71410abce5.1628578664.1.0.0; mtcdn=K; lsu=; _lxsdk_cuid=17b2eda6b5bc8-0cc3ace2b4b36f-113a6054-13c680-17b2eda6b5bc8; IJSESSIONID=node01gtbt2ztg0q2g8drdwlnb4fol18090423; iuuid=156B7868C8A4C16CE2BC984983A5B34BB732C4F3C970F09E810BECCCD7647A45; cityname=%E4%B8%8A%E6%B5%B7; _lxsdk=156B7868C8A4C16CE2BC984983A5B34BB732C4F3C970F09E810BECCCD7647A45; u=2848485011; n=tqz586847076; lt=qMnTCq3MPInnEcpJmainL0KUOowAAAAAUg4AAH1dAJoVg-ltw1veuUcacY8Qd6Tt62ZZBl44xXbexcW2VY-LQy94ARJimqAOpqL4ow; mt_c_token=qMnTCq3MPInnEcpJmainL0KUOowAAAAAUg4AAH1dAJoVg-ltw1veuUcacY8Qd6Tt62ZZBl44xXbexcW2VY-LQy94ARJimqAOpqL4ow; token=qMnTCq3MPInnEcpJmainL0KUOowAAAAAUg4AAH1dAJoVg-ltw1veuUcacY8Qd6Tt62ZZBl44xXbexcW2VY-LQy94ARJimqAOpqL4ow; token2=qMnTCq3MPInnEcpJmainL0KUOowAAAAAUg4AAH1dAJoVg-ltw1veuUcacY8Qd6Tt62ZZBl44xXbexcW2VY-LQy94ARJimqAOpqL4ow; unc=tqz586847076; __mta=145193483.1628578672900.1628660641898.1628664425887.17; firstTime=1628670536896; _lxsdk_s=17b348515ca-204-947-09c%7C%7C2; ci=40; rvct=40%2C59%2C1%2C10")
		req.Headers.Set("Host", "apimobile.meituan.com")
		req.Headers.Set("Referer", "https://cd.meituan.com/")
		req.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("--OnResponse--")
		fmt.Print(string(r.Body))
	})
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("--OnError--")
		fmt.Println(e.Error())
	})
	fmt.Println("----")
	// c.Visit("https://www.meituan.com/ptapi/getLoginedUserInfo")
	c.Visit("https://apimobile.meituan.com/group/v4/poi/pcsearch/40?limit=32&offset=768&q=折扣")
	c.Wait()
}
