package controllers

import (
	"beego_action/models"
	_ "fmt"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (a *ArticleController) TestFunc() {
	article_info := &models.Article{}
	article_info.Read()
}
