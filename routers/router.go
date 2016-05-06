package routers

import (
	"beego_code/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//自定义方法及 RESTful 规则
	beego.Router("/", &controllers.MainController{})
	beego.Router("user/getuser", &controllers.MainController{}, "get:GetUser")
	beego.Router("user/updateuser", &controllers.MainController{}, "*:UpdateUser")
	beego.Router("user/deleteuser", &controllers.MainController{}, "get,post:DeleteUser")
	beego.Router("user/adduser", &controllers.MainController{}, "get:AddUser;post:UpdateUser")
	beego.Router("user/get_config_data", &controllers.MainController{}, "get:GetConfigData")

	beego.Router("/msg/get_msg_by_id", &controllers.MsgController{}, "get:GetMsgById")

	//注解路由 ????
	//beego.Include(&controllers.MsgController{})
	//beego.Include(&controllers.CMSController{})

	ns := beego.NewNamespace("/v2",
		beego.NSNamespace("/log",
			beego.NSInclude(
				&controllers.LogController{},
			),
		),

		// beego.NSNamespace("/user",
		//     beego.NSInclude(
		//         &controllers.UserController{},
		//     ),
		// ),
	)
	beego.AddNamespace(ns)
}
