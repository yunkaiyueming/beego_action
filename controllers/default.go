package controllers

import (
	"beego_action/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["test"] = "hello, study go"

	c.Data["appname"] = beego.AppConfig.String("appname")
	c.Data["static_url"] = beego.AppConfig.String("static_url")
	c.Data["test_str"] = beego.AppConfig.String("prod::test_str")
	models.Record_log("eme", "get user msg")

	//session操作
	c.MySessionOp()

	//URL构建操作
	c.MyUrlFor()

	c.TplName = "index.tpl"
	c.Render()
}

func (c *MainController) GetAllUser() {
	u := &models.UserModel{}
	users, err := u.GetAllUser()
	models.CheckError(err)
	fmt.Println(users)

	c.Data["users"] = users
	c.TplName = "user/get_user.tpl"
	c.Render()
}

//UpdateUser is
func (c *MainController) UpdateUser() {
	c.Ctx.WriteString("hello updateuser")
}

//AddUser is
func (c *MainController) AddUser() {
	u := &models.UserModel{}
	id := u.AddUser()
	if id > 0 {
		msg := fmt.Sprintf("add user success %d", id)
		c.Ctx.WriteString(msg)
	} else {
		c.Ctx.WriteString("add user false")
	}
}

func (c *MainController) DeleteUser() {
	c.Ctx.WriteString("hello deleteuser")
}

func (c *MainController) GetConfigData() {
	iniconf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic("get app config is wrong")
	}

	c.Data["appname"] = iniconf.String("appname")
	c.Data["appid"], err = iniconf.Int64("Dev::appid")
	if err != nil {
		panic("appid is wrong")
	}
	c.TplName = "get_config_data.tpl"
}

func (c *MainController) MySessionOp() {
	v := c.GetSession("pv")
	if v != nil {
		v = v.(int) + 1
	} else {
		v = int(1)
	}
	c.SetSession("pv", v)

	c.SessionRegenerateID()
	fmt.Println(v)
}

func (c *MainController) MyUrlFor() {
	c.URLFor("TestController.List")
	// 输出 /api/list

	url2 := c.URLFor("TestController.Get", ":last", "xie", ":first", "asta")
	fmt.Println(url2)
	// 输出 /person/xie/asta

	c.URLFor("TestController.Myext")
	// 输出 /Test/Myext

	c.URLFor("TestController.GetUrl")
	// 输出 /Test/GetUrl

	//c.Abort()
}
