package main

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
)

func testBase64() {
	bs := `/9j/4AAQSkZJRgABAQAAAQABAAD//gA+Q1JFQVRPUjogZ2QtanBlZyB2MS4wICh1c2luZyBJSkcg
	SlBFRyB2NjIpLCBkZWZhdWx0IHF1YWxpdHkK/9sAQwAIBgYHBgUIBwcHCQkICgwUDQwLCwwZEhMP
	FB0aHx4dGhwcICQuJyAiLCMcHCg3KSwwMTQ0NB8nOT04MjwuMzQy/9sAQwEJCQkMCwwYDQ0YMiEc
	ITIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIy/8AAEQgB
	aQEnAwEiAAIRAQMRAf/EAB8AAAEFAQEBAQEBAAAAAAAAAAABAgMEBQYHCAkKC//EALUQAAIBAwMC
	BAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYn
	KCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeY
	mZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5
	+v/EAB8BAAMBAQEBAQEBAQEAAAAAAAABAgMEBQYHCAkKC//EALURAAIBAgQEAwQHBQQEAAECdwAB
	AgMRBAUhMQYSQVEHYXETIjKBCBRCkaGxwQkjM1LwFWJy0QoWJDThJfEXGBkaJicoKSo1Njc4OTpD
	REVGR0hJSlNUVVZXWFlaY2RlZmdoaWpzdHV2d3h5eoKDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ip
	qrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uLj5OXm5+jp6vLz9PX29/j5+v/aAAwDAQACEQMR
	AD8A9/ooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAK
	KKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKKKACiiigAooooAKKimnjgTdI2BX
	G+JPiHpWhJ5TtIbpx8qdNo/vGgDo7vWrG1crNI6sP7qFv5VLbatZXabobhH9cfw/WvBL/wCKd9PN
	/o1tFHbfd2/eP5t/hWGfGOso8nkXbwh/lO3rilzD5T6Fv/EdvYXPkyPCvO3zHk2qG54rB1D4h6dY
	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`
	bs = strings.Replace(bs, " ", "", -1)
	uDec, err := base64.URLEncoding.DecodeString(bs)
	if err != nil {
		fmt.Printf("Error decoding string: %s ", err.Error())
		return
	}

	fmt.Println(string(uDec))
}

func testReg() {
	s := `"=?gb18030?B?wLXV5Mjj?=" <rigphiojfs@@vip.sohu.com>`
	reg := regexp.MustCompile(`\w+([-+.]*\w+@)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	from := reg.FindAllString(s, -1)
	fmt.Println(from)
}

func testTrim() {
	ss := "----=_Part_11916_1472422870.1593446519845"
	ss = strings.Trim(ss, "-")
	fmt.Println(ss)
}

func testContains() {
	str := `<!doctype html>
	<html>
	<head>
	<meta charset="UTF-8">
	<title>中高端人才招聘-猎聘</title>
	<!--  -->
	<meta name="renderer" content="webkit"><!-- 使360浏览器 默认启用极速内核 -->
	<link rel="icon" href="//concat.lietou-static.com/fe-nlpt-pc/v5/static/favicon.b16a8905.ico" type="image/x-icon" />
	<link rel="dns-prefetch" href="//concat.lietou-static.com" />
	<!-- AppAdhoc start -->
	<script src='https://sdk.gua.com/ab.plus.js'></script>
	<script>
	adhoc('init', {
		appKey: 'ADHOC_8a5bdd68-c964-4a5d-8c23-297b7ba0ea58',
		clientIdDomain:'liepin.com'
	})
	</script>
	<!-- AppAdhoc end -->

	<!-- 浏览器升级提醒 -->
	<!--[if IE]>
	<script crossorigin="anonymous" src="//concat.lietou-static.com/fe-nlpt-pc/v5/static/js/ie-prompt.ec60e546.js"></script>
	<![endif]-->
	<script crossorigin="anonymous" src="//concat.lietou-static.com/fe-nlpt-pc/v5/static/js/femonitor.min.ee371522.js"></script>
	<script crossorigin="anonymous" src="//concat.lietou-static.com/fe-nlpt-pc/v5/static/js/loader.f5cc4300.js"></script>
	<script crossorigin="anonymous" src="//concat.lietou-static.com/fe-nlpt-pc/v5/static/js/jquery-1.7.1.min.c7e0488b.js"></script>
	<script crossorigin="anonymous" src="//concat.lietou-static.com/fe-nlpt-pc/v5/static/js/polyfill.min.7d2ef4bb.js"></script>
	<script crossorigin="anonymous" src="//concat.lietou-static.com/fe-nlpt-pc/v5/static/js/common/lib.0f2577a7.js"></script>
	<link rel="stylesheet" href="//concat.lietou-static.com/fe-nlpt-pc/v5/css/common/common.900ffa6e.css">
	<script crossorigin="anonymous" src="//concat.lietou-static.com/fe-nlpt-pc/v5/js/common/common.4027d324.js"></script>
	<link rel="stylesheet" href="//concat.lietou-static.com/fe-nlpt-pc/v5/css/common/vendors.d5636985.css">
	<script crossorigin="anonymous" src="//concat.lietou-static.com/fe-nlpt-pc/v5/js/common/vendors.8d96b3e4.js"></script>
	<link rel="stylesheet" href="//concat.lietou-static.com/fe-nlpt-pc/v5/static/css/swiper.min.6c1ec3a0.css">
	<link rel="stylesheet" href="//concat.lietou-static.com/fe-nlpt-pc/v5/css/pages/user.login.a7907b95.css">
	<script id="CaptchaScript" src="https://captcha.myqcloud.com/Captcha.js"></script>
	</head>
	<body>
	<div id="root"></div>

	<script src="//concat.lietou-static.com/fe-nlpt-pc/v5/static/js/swiper.min.bc508491.js"></script>
	<script src="//concat.lietou-static.com/fe-nlpt-pc/v5/js/pages/user.login.dd26c9a8.js"></script>
	<!--  -->
	<script src="//concat.lietou-static.com/fe-nlpt-pc/v5/static/js/tlog.min.4a361aeb.js"></script>
	<script>

	(function () {
	var dlog_js = document.createElement("script");
	dlog_js.src = "//static3.lietou-static.com/dlog.js?v=3&q=" + parseInt(''+new Date()/3E5);
	var s = document.getElementsByTagName("script")[0];
	s.parentNode.insertBefore(dlog_js, s);
	})();

	</script>

	</body>
	</html>
		`
	fmt.Println(strings.Contains(str, `<div id="root"></div>`))
}
func main() {
	// testTrim()
	testContains()
}
