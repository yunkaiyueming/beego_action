package controllers

import (
	"github.com/astaxie/beego"
)

type LogController struct {
	beego.Controller
}

// @Description 注解路由形式测试
// @router /:get_log_by_id [get]
func (m *LogController) GetLogById() {
	m.Ctx.WriteString("get log controller")
}
