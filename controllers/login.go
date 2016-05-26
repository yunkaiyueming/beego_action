package controllers

import (
	"fmt"

	"beego_action/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Login() {
	name := l.GetString("name")
	pwd := l.GetString("pwd")
	action := l.GetString("action")
	if action == "" {
		fmt.Println("action empty")
		l.TplName = "user/login.tpl"
		l.Render()
		return
	}

	user_model := &models.UserModel{}
	user_info := user_model.CheckGetUser(name, pwd)

	if user_info.Id == 0 {
		fmt.Println("login false")
		l.Data["error_msg"] = "user login error"
		l.TplName = "user/login.tpl"
		l.Render()
		return
	} else {
		l.SetSession("name", user_info.Name)
		l.SetSession("id", user_info.Id)
		fmt.Println("recording session")

		//????跳转不成功
		l.Redirect("user/welcome", 200)
		return
	}

	fmt.Println("bad log")
}
