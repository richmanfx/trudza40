package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
	c.Data["title"] = "Труд за 40"
}

func (c *MainController) Realty() {
	c.TplName = "realty.tpl"
	c.Data["title"] = "Realty"
}
