package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["beego_action/controllers:LogController"] = append(beego.GlobalControllerRouter["beego_action/controllers:LogController"],
		beego.ControllerComments{
			"GetLogById",
			`/get_log_by_id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beego_action/controllers:LogController"] = append(beego.GlobalControllerRouter["beego_action/controllers:LogController"],
		beego.ControllerComments{
			"GetConfig",
			`/my_config`,
			[]string{"get"},
			nil})

}
