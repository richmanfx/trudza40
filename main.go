package main

import (
	"github.com/astaxie/beego"
	"trudza40/controllers"
	_ "trudza40/routers"
)

func main() {
	controllers.SessionInit()
	beego.Run()
}
