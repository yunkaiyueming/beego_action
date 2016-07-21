package controllers

import (
	"fmt"
)

type HomeController struct {
	BaseController
}

type MachineConfig struct {
	Name   string
	Ip     string
	User   string
	Port   int
	Dbinfo string
}

func (this *HomeController) Construct() {
	fmt.Println("--home construct--")
}

func (this *HomeController) Index() {
	this.CheckLogin()
	this.getMachineConfig()
	this.MyRender("home/view_machine.html")
}

func (this *HomeController) getMachineConfig() {
	machineConfigData := map[string]MachineConfig{
		"bi":        {Name: "bi", Ip: "s119.00.25.59", User: "00", Port: 10220},
		"oa":        {Name: "oa", Ip: "s119.29.00.59", User: "00", Port: 10220},
		"rsdk-set":  {Name: "rsdk-set", Ip: "s119.00.25.59", User: "00", Port: 10220},
		"bi2-agent": {Name: "bi2-agent", Ip: "s119.00.25.59", User: "00", Port: 10220},
	}

	this.Data["machineConfigData"] = machineConfigData
}
