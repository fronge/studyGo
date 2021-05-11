package main
import (
	"fmt"
	"regexp"
)

var body = `<li><label>
座&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;机：</label><span style="font-size:13px;"><span class="secret">12312313-123132</span></span>
</li>`
func main() {
	phone := regexp.MustCompile(`<span class="secret">(.*?)<`).FindStringSubmatch(body)
	fmt.Println(phone)
}