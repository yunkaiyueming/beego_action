package controllers

import (
	"fmt"

	"beego_action/helpers"
	"beego_action/models"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Construct() {
	this.sidebarFile = "include/sidebar/classic_sidebar.html"
	this.headerFile = "include/header.html"
	this.layoutFile = "include/layout/main.html"
	this.footerFile = "include/footer.html"
}

func (this *LoginController) Login() {
	name := this.GetString("name")
	pwd := this.GetString("pwd")
	action := this.GetString("action")
	if action == "" {
		this.Construct()
		this.MyRender("login/view_login.html")
		return
	}

	user_model := &models.UserModel{}
	user_info := user_model.CheckGetUser(name, pwd)

	if user_info.Id == 0 {
		fmt.Println("login false")
		this.Data["error_msg"] = "user login error"
		this.MyRender("login/view_login.html")
		return
	} else {
		this.SetSession("name", user_info.Name)
		this.SetSession("id", user_info.Id)
		this.Redirect(helpers.SiteUrl("home/index"), 302)
		return
	}

	fmt.Println("bad log")
}

func (this *LoginController) Logout() {
	name := this.GetSession("name")
	if name != nil {
		this.DestroySession()
	}
	this.Redirect(helpers.SiteUrl("login/login"), 302)
}

func (this *LoginController) Seting() {
	this.GetSessionUser()
}

func (this *LoginController) GetSessionUser() models.UserModel {
	id := this.GetSession("id")
	userInfo := models.UserModel{}
	if id != nil {
		userModel := &models.UserModel{}
		userInfo = userModel.GetUserById(id.(int))
	}

	fmt.Println(userInfo)
	return userInfo
}
