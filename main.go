package main

import (
	"github.com/astaxie/beego"
	_ "trudza40/routers"
)

func main() {
	//controllers.SessionInit()
	beego.Run()
}
