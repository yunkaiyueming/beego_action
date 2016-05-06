package controllers

import (
	"github.com/astaxie/beego"
)

type LogController struct {
	beego.Controller
}

func (m *LogController) GetLogById() {
	m.Ctx.WriteString("get log controller")
}
