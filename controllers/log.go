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

func Record_log(msg_level, msg string) {
	beego.SetLogger("file", `{"filename":"E:/GO_PATH/src/beego_action/test.log"}`)
	//beego.BeeLogger.DelLogger("console")
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(true)

	beego.Emergency(msg)
}
