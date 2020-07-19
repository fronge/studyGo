package email

import (
	"bigdata/grab/spiders/tools"
)

var _ tools.Spider = &Spider{}

type Spider struct {
	*tools.Base
}

func NewSpider() tools.Spider {
	s := &Spider{Base: &tools.Base{}}
	return s
}
