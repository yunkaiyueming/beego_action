package controllers

import (
	"beego_action/models"
	"fmt"
)

type TranslateController struct {
	BaseController

	lang string
}

func (this *TranslateController) GetLang() {
	if this.Ctx.GetCookie("lang") == "" {
		this.Ctx.SetCookie("lang", "zn")
		this.lang = "zn"
	} else {
		this.lang = this.Ctx.GetCookie("lang")
	}

	setction := "Tittle"
	key := "7"
	val := models.GetTranslateKey(this.lang, setction, key)
	fmt.Println(val)
}
