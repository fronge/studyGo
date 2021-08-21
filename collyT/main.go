package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
	// colly.AllowedDomains("dianping.com"),
	)
	c.OnRequest(func(req *colly.Request) {
		fmt.Println("Visiting", req.URL)
		req.Headers.Add("Host", "www.dianping.com")
		req.Headers.Add("Pragma", "no-cache")
		req.Headers.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
		req.Headers.Add("Accept-Encoding", "gzip, deflate")
		req.Headers.Add("Cache-Control", "no-cache")
		req.Headers.Add("Upgrade-Insecure-Requests", "1")
		req.Headers.Add("Accept-Encoding", "gzip, deflate")
		req.Headers.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		req.Headers.Add("Referer", "http://www.dianping.com/search/keyword/10/0_%E6%8A%98%E6%89%A3/p5")

		req.Headers.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
		req.Headers.Add("Cookie", "_lxsdk_cuid=177fbd285f3c8-02abce0bfbf566-193e6d50-1fa400-177fbd285f3c8; _lxsdk=177fbd285f3c8-02abce0bfbf566-193e6d50-1fa400-177fbd285f3c8; _hc.v=2ebd84dc-a82b-dc2d-3378-70a15e68ac03.1624107443; s_ViewType=10; fspop=test; _lx_utm=utm_source%3DBaidu%26utm_medium%3Dorganic; Hm_lvt_602b80cf8079ae6591966cc70a3940e7=1628149400; ll=7fd06e815b796be3df069dec7836c3df; ua=dpuser_0645025773; ctu=67d68e72361ddcb8e020661fa18758669cfc441e67066dbcd5551e9cebd511a1; uamo=13614205855; dper=52a8b9a1f5440b35c79f08fb1f9d4fca2bd9b3e130ef655b759148c822427a75b0198b19dbb76b79a428e00101ac193702206026f634bf413fd4945edf32ea9699bc4e0b58b70ff51edd1e163c540754f1dfdab746d04721665afd205684de53; dplet=e7fe65c536eae3f7450f6f79f5d5ef94; cy=10; cye=tianjin; Hm_lvt_dbeeb675516927da776beeb1d9802bd4=1628216061; Hm_lpvt_dbeeb675516927da776beeb1d9802bd4=1628216061; Hm_lpvt_602b80cf8079ae6591966cc70a3940e7=1628216124; _lxsdk_s=17b1937e164-c5b-473-cfe%7C%7C261")
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
	c.Visit("http://www.dianping.com/shop/H1pSh7In5XgwwmtQ")
	c.Wait()
}
