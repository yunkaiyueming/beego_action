package controllers

import (
	"beego_action/models"
	_ "encoding/json"
	_ "fmt"

	"github.com/astaxie/beego"
)

type GameController struct {
	beego.Controller
}

//API JSON形式
func (g *GameController) GetGameJson() {
	GameModel := &models.GameModel{}
	GameInfoJson := GameModel.GetGameJson()
	g.Data["json"] = GameInfoJson
	g.ServeJSON()
}
