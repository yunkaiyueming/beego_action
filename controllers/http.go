package controllers

import (
	"beego_action/helpers"
	"fmt"
	"io/ioutil"
	_ "strings"

	"github.com/astaxie/beego/httplib"
)

type HttpController struct {
	BaseController
}

func (this *HttpController) Construct() {
	this.headerFile = "include/header.html"
	this.footerFile = "include/footer.html"
	this.layoutFile = "include/layout/classic.html"
	this.sidebarFile = "include/sidebar/classic_sidebar.html"
}

func (this *HttpController) MyRender(viewFile string) {
	this.Construct()
	this.BaseController.MyRender(viewFile)
	fmt.Println("http render")
}

func (this *HttpController) Get() {
	request := this.Ctx.Request
	url := "http://www.baidu.com"
	response := helpers.HttpGetDo(url, request)

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	//请求信息
	this.Data["req_method"] = request.Method
	this.Data["req_proto"] = request.Proto
	this.Data["req_headers"] = request.Header
	this.Data["url"] = url
	//响应信息
	this.Data["proto"] = response.Proto
	this.Data["code"] = response.Status
	this.Data["rep_headers"] = response.Header
	this.Data["ret"] = string(body)
	this.MyRender("http/view_get.html")
}

func (this *HttpController) UserHttpLib() {
	//httplib.Get("http://www.kancloud.cn/hello123/beego/126134").Debug(true).Response()
	req := httplib.Get("http://www.baidu.com")
	str, _ := req.String()
	this.Ctx.WriteString(str)
}

func (this *HttpController) UseContext() {
	fmt.Println(this.Ctx.Input.Domain())
	fmt.Println(this.Ctx.Input.Site())
	fmt.Println(this.Ctx.Input.Host())
	fmt.Println(this.Ctx.Input.URI()) //带查询参数
	fmt.Println(this.Ctx.Input.URL())

}

func (this *HttpController) TestNet() {
	helpers.GetHost()
}
