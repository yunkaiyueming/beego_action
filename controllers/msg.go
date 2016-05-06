package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
)

type MsgController struct{
	beego.Controller
}

func (m *MsgController) GetMsgById(){
	m.Ctx.WriteString("GetMsgById")
}

