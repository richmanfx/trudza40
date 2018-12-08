package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
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

	if userName == "Вася" {
		err = nil
	}

	//userPassword := context.GetString("user_password")

	// Проверить существование пользователя в базе	TODO
	err = models.CheckUserInDB(userName)

	// Проверить пароль по Хешу из БД	TODO

	if err == nil {
		log.Infof("Пользователь '%s' вошёл в приложение.", userName)
		//helpers.UserLogin = userName	// Логин пользователя в заголовок

		// Изменить сессию

		// Сгенерировать sessid

		// Сохранить сессию в БД

		// Выставить в браузере Куки

		// Направить на индексную страницу
		//context.Abort(string(http.StatusOK))
		context.Redirect("/realty", http.StatusSeeOther)

	} else {
		// Вывод сообщения об ошибке в окне (модальном?)
		log.Errorln("Ошибка авторизации - неверный логин/пароль.")
		context.TplName = "message-modal.tpl"
		context.Data["title"] = "Ошибка"
		context.Data["message1"] = "Ошибка"
		context.Data["message2"] = "Ошибка авторизации - неверный логин/пароль."
		//context.Data["message3"] = err

	}
}
