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
	beego.Router("/edit-config", &controllers.ConfigController{}, "get:EditConfig")
	beego.Router("/users-config", &controllers.ConfigController{}, "get:UsersConfig")
	beego.Router("/create-user", &controllers.ConfigController{}, "get:CreateUser")
	beego.Router("/create-user-in-db", &controllers.ConfigController{}, "post:CreateUserInDb")
}
