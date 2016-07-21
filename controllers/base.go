package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller

	layoutFile  string
	headerFile  string
	sidebarFile string
	footerFile  string
}

func (this *BaseController) Construct() {
	fmt.Println("--base construct--")
	this.sidebarFile = "include/sidebar/classic_sidebar.html"
	this.headerFile = "include/header.html"
	this.layoutFile = "include/layout/classic.html"
	this.footerFile = "include/footer.html"
}

func (this *BaseController) MyRender(viewFile string) {
	this.Construct()

	this.Layout = this.layoutFile
	this.TplName = viewFile

	this.LayoutSections = make(map[string]string)
	this.LayoutSections["headerFile"] = this.headerFile
	this.LayoutSections["footerFile"] = this.footerFile
	this.LayoutSections["sidebarFile"] = this.sidebarFile

	this.PrepareViewData()
	this.Render()
}

func (this *BaseController) PrepareViewData() {
	staticUrl := beego.AppConfig.String("static_url")
	siteUrl := beego.AppConfig.String("siteUrl")

	this.Data["staticUrl"] = staticUrl
	this.Data["siteUrl"] = siteUrl
}

func (this *BaseController) CheckLogin() bool {
	id := this.GetSession("id")
	name := this.GetSession("name")
	if id != nil && name != nil {
		return true
	} else {
		this.Redirect("home/index", 200)
		return false
	}
}
