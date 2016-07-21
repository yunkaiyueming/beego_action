package routers

import (
	"beego_action/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//自定义方法及 RESTful 规则
	beego.Router("user/welcome", &controllers.MainController{})
	beego.Router("user/get_all_user", &controllers.MainController{}, "get:GetAllUser")
	beego.Router("user/update_user", &controllers.MainController{}, "*:UpdateUser")
	beego.Router("user/delete_user", &controllers.MainController{}, "get,post:DeleteUser")
	beego.Router("user/add_user", &controllers.MainController{}, "get:AddUser;post:UpdateUser")
	beego.Router("user/get_config_data", &controllers.MainController{}, "get:GetConfigData")

	beego.Router("game/get_game_json", &controllers.GameController{}, "get:GetGameJson")

	//机器管理
	beego.Router("/", &controllers.HomeController{}, "get:Index")
	beego.Router("home/index", &controllers.HomeController{}, "get:Index")
	beego.Router("performance/index", &controllers.PerformanceController{}, "get:Index")
	beego.Router("performance/get_go_env", &controllers.PerformanceController{}, "get:GetGoEnv")

	beego.Router("json/", &controllers.JsonTestController{}, "get:Get")
	beego.Router("json/claw_url_msg", &controllers.JsonTestController{}, "get:ClawUrlMsg")
	beego.Router("json/claw_resonse_header", &controllers.JsonTestController{}, "get:ClawResponseHeader")

	beego.Router("qiniu/get_files_msg", &controllers.QiniuController{}, "get:GetFilesMsg")
	beego.Router("qiniu/get_down_file_url", &controllers.QiniuController{}, "get:GetDownFileUrl")
	beego.Router("qiniu/simple_upload_file", &controllers.QiniuController{}, "get:SimpleUploadFile")

	//登录,权限相关
	beego.Router("login/login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("login/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("login/seting", &controllers.LoginController{}, "get:Seting")

	beego.Router("article/test", &controllers.ArticleController{}, "get:TestFunc")

	//注解路由
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
