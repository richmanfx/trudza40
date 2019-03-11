package routers

import (
	"github.com/astaxie/beego"
	"trudza40/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.Router("/realty", &controllers.MainController{}, "get:Realty")
	beego.Router("/realty/login", &controllers.AuthController{}, "get:Login")
	beego.Router("/realty/login-processing", &controllers.AuthController{}, "post:LoginProcessing")
	beego.Router("/realty/logout", &controllers.AuthController{}, "get:Logout")

	beego.Router("/realty/edit-config", &controllers.ConfigController{}, "get:EditConfig")
	beego.Router("/realty/users-config", &controllers.ConfigController{}, "get:UsersConfig")
	beego.Router("/realty/create-user", &controllers.ConfigController{}, "get:CreateUser")
	beego.Router("/realty/create-user-in-db", &controllers.ConfigController{}, "post:CreateUserInDb")
	beego.Router("/realty/delete-user", &controllers.ConfigController{}, "post:DeleteUser")
	beego.Router("/realty/change-password", &controllers.ConfigController{}, "post:ChangePassword")

	beego.Router("/realty/settings_torgi_gov_ru", &controllers.ScrapController{}, "get:TorgiGovRuSettings")
	beego.Router("/realty/scraping_torgi_gov_ru", &controllers.ScrapController{}, "get:TorgiGovRuScraping")

}
