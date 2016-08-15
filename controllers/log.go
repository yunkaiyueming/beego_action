package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	_ "github.com/astaxie/beego/logs"
)

type LogController struct {
	beego.Controller
}

// @Description 注解路由形式测试
// @router /get_log_by_id [get]
func (this *LogController) GetLogById() {
	this.Ctx.WriteString("get log controller")
}

func (this *LogController) Record_log(msg_level, msg string) {
	beego.SetLogger("file", `{"filename":"E:/GO_PATH/src/beego_action/test.log"}`)
	//beego.BeeLogger.DelLogger("console")
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(true)

	beego.Emergency(msg)
}

// @Description 不同方式获取配置信息
// @router /my_config [get]
func (this *LogController) GetConfig() {
	qiniu_ak := beego.AppConfig.String("qiniu_ak")
	fmt.Println(qiniu_ak)

	machine_id := beego.AppConfig.String("machine_id")
	prodAppId := beego.AppConfig.String(machine_id + "::appid") //支持runmode::key
	fmt.Println(prodAppId)

	iniconf, _ := config.NewConfig("ini", "conf/app.conf")
	http := iniconf.String("mysqlurls")
	fmt.Println(http)
	this.StopRun()
}

func (this *LogController) LogTest() {

}
