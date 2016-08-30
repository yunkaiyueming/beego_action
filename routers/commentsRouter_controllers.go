package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["beego_action/controllers:LogController"] = append(beego.GlobalControllerRouter["beego_action/controllers:LogController"],
		beego.ControllerComments{
			Method: "GetLogById",
			Router: `/get_log_by_id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["beego_action/controllers:LogController"] = append(beego.GlobalControllerRouter["beego_action/controllers:LogController"],
		beego.ControllerComments{
			Method: "GetConfig",
			Router: `/my_config`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
