package controllers

import (
	"beego_action/helpers"
	"fmt"
	"io/ioutil"
	_ "strings"
)

type HttpController struct {
	BaseController
}

func (this *HttpController) Construct() {
	fmt.Println("--http construct--")
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
