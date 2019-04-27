package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session"
)

type MainController struct {
	beego.Controller
}

func (controller *MainController) Get() {
	CheckSession(controller)
	controller.TplName = "index.tpl"
	controller.Data["title"] = "Труд за 40" // TODO: Отказаться от "Труд за 40"
}

func (controller *MainController) Realty() {
	//CheckSession(controller)
	controller.TplName = "realty.tpl"
	controller.Data["title"] = "Realty"
}
