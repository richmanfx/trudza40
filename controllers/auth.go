package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
	"trudza40/models"
)

type AuthController struct {
	beego.Controller
}

func (context *AuthController) Login() {
	context.TplName = "login.tpl"
	context.Data["title"] = "Realty Login"
}

func (context *AuthController) LoginProcessing() {

	err := errors.New(fmt.Sprintln("err: 'Ошибка'"))
	userName := context.GetString("user_name")
	userPassword := context.GetString("user_password")

	//if userName == "Вася" {
	//	err = nil
	//}

	// Проверить существование пользователя в базе
	err = models.CheckUserInDB(userName)

	// Проверить пароль по Хешу из БД	TODO
	if err == nil {
		// Можно закомментировать для возможности залогиниться без Хешей из БД
		err = models.CheckPasswordInDB(userName, userPassword)
	}

	if err == nil {
		beego.Info(fmt.Sprintf("Пользователь '%s' вошёл в приложение.", userName))
		//helpers.UserLogin = userName	// Логин пользователя в заголовок

		// Изменить сессию

		// Сгенерировать sessid

		// Сохранить сессию в БД

		// Выставить в браузере Куки

		// Направить на индексную страницу
		//context.Abort(string(http.StatusOK))
		context.Redirect("/realty", http.StatusSeeOther)

	} else {
		beego.Error("Ошибка авторизации - неверный логин/пароль.")

		// Вывод сообщения об ошибке в окне модальном
		context.TplName = "message-modal.tpl"
		context.Data["title"] = "Ошибка"
		context.Data["message1"] = "Ошибка"
		context.Data["message2"] = "Ошибка авторизации - неверный логин/пароль."
		//context.Data["message3"] = err

	}
}
