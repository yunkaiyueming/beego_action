package controllers

import (
	_ "fmt"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Index() {
	viewFile := "home/view_machine.html"
	this.MyRender(viewFile)
}

func (this *HomeController) MyRender(viewFile string) {
	this.Layout = "include/layout/classic.html"
	this.TplName = viewFile
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["headerFile"] = "include/header.html"
	this.LayoutSections["footerFile"] = "include/footer.html"
	this.LayoutSections["sidebarFile"] = "include/sidebar/classic_sidebar.html"

	this.Data["staticUrl"] = this.PrepareData()
	this.Render()
}

func (this *HomeController) PrepareData() string {
	staticUrl := beego.AppConfig.String("static_url")
	return staticUrl
}

func init() {

}
