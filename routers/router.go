package routers

import (
	"beego_action/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//信息管理
	beego.Router("user/index", &controllers.UserController{}, "get:GetAllUser")
	beego.Router("user/update_user", &controllers.UserController{}, "*:UpdateUser")
	beego.Router("user/delete_user", &controllers.UserController{}, "get,post:DeleteUser")
	beego.Router("user/add_user", &controllers.UserController{}, "get:AddUser")
	beego.Router("user/nginx", &controllers.UserController{}, "get:Nginx")

	//机器管理
	beego.Router("/", &controllers.HomeController{}, "get:Index")
	beego.Router("home/index", &controllers.HomeController{}, "get:Index")
	beego.Router("performance/index", &controllers.PerformanceController{}, "get:Index")
	beego.Router("performance/get_go_env", &controllers.PerformanceController{}, "get:GetGoEnv")

	//接口管理
	beego.Router("http/get", &controllers.HttpController{}, "get:Get")
	beego.Router("http/user_http_lib", &controllers.HttpController{}, "get:UserHttpLib")
	beego.Router("http/use_context", &controllers.HttpController{}, "get:UseContext")

	beego.Router("json/", &controllers.JsonTestController{}, "get:Get")
	beego.Router("json/claw_url_msg", &controllers.JsonTestController{}, "get:ClawUrlMsg")
	beego.Router("json/claw_resonse_header", &controllers.JsonTestController{}, "get:ClawResponseHeader")
	beego.Router("json/test_xml", &controllers.JsonTestController{}, "get:TestXml")

	//文件模块
	beego.Router("qiniu/get_files_msg", &controllers.QiniuController{}, "get:GetFilesMsg")
	beego.Router("qiniu/get_down_file_url", &controllers.QiniuController{}, "get:GetDownFileUrl")
	beego.Router("qiniu/simple_upload_file", &controllers.QiniuController{}, "get:SimpleUploadFile")
	beego.Router("qiniu/index", &controllers.QiniuController{}, "get:GetBucketFiles")
	beego.Router("file/get_file_con", &controllers.FileController{}, "get:GetFileCon")
	beego.Router("file/list_dir", &controllers.FileController{}, "get:ListDir")

	//登录,权限相关
	beego.Router("login/login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("login/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("login/seting", &controllers.LoginController{}, "get:Seting")

	//国际化翻译模块
	beego.Router("translate/", &controllers.TranslateController{}, "get:GetLang")

	//RPC
	beego.Router("rpc/http", &controllers.RpcController{}, "get:Client_http")
	beego.Router("rpc/json", &controllers.RpcController{}, "get:Client_json")
	beego.Router("rpc/tcp", &controllers.RpcController{}, "get:Client_tcp")

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
