package routers

import (
	"github.com/astaxie/beego"
	"trudza40/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/realty", &controllers.MainController{}, "get:Realty")
	beego.Router("/login", &controllers.AuthController{}, "get:Login")
	beego.Router("/login-processing", &controllers.AuthController{}, "post:LoginProcessing")
}
