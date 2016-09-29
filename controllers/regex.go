package controllers

import (
	"beego_action/helpers"
	"fmt"
)

type RegexController struct {
	BaseController
}

func (this *RegexController) Check() {
	p := fmt.Println
	p(helpers.CheckEmail("14sdd8fdf4qq.com"))
	p(helpers.CheckIp("12.32.324d.4"))
	p(helpers.CheckNum("23432dar324"))
	p(helpers.CheckDate("20s15-01-05"))
	p(helpers.CheckUrl("http:/s/www.baidu.com"))
	p(helpers.CheckPhone("1379d7592668"))
}
