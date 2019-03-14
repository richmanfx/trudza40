package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"net/http"
	"trudza40/models"
)

var GlobalUserId = 0
var GlobalUserLogin string

type AuthController struct {
	beego.Controller
}

func (controller *AuthController) Login() {
	controller.TplName = "login.tpl"
	controller.Data["title"] = "Realty Login"
}

func (controller *AuthController) LoginProcessing() {

	err := errors.New(fmt.Sprintln("err: 'Ошибка'"))
	userName := controller.GetString("user_name")
	userPassword := controller.GetString("user_password")
	var userId int

	// Проверить существование пользователя в базе
	userId, err = models.CheckUserInDB(userName)
	GlobalUserId = userId

	// Проверить пароль по Хешу из БД
	if err == nil {
		// Можно закомментировать для возможности залогиниться без Хешей из БД
		err = models.CheckPasswordInDB(userName, userPassword)
	}

	if err == nil {
		beego.Info(fmt.Sprintf("Пользователь '%s' вошёл в приложение.", userName))

		// Логин пользователя для заголовка
		GlobalUserLogin = userName // TODO: Пока заголовок не реализован

		// Добавить куку с ID полльзователя
		err = userSession.Set("UserID", userId)
		if err != nil {
			beego.Error("Не добавился UserID")
		}

		// Новая сессия с новой кукой
		userID := controller.StartSession().Get("UserID") // Новая сессия
		beego.Info(fmt.Sprintf("UserID: %v", userID))
		if userID == nil {
			return
		}

		controller.Redirect("/realty", http.StatusSeeOther)

	} else {
		beego.Error("Ошибка авторизации - неверный логин/пароль.")

		// Вывод сообщения об ошибке в модальном окне
		controller.TplName = "message-modal.tpl"
		controller.Data["title"] = "Ошибка"
		controller.Data["message1"] = "Ошибка"
		controller.Data["message2"] = "Ошибка авторизации - неверный логин/пароль"
	}
}

/* Разлогинивание */
func (controller *AuthController) Logout() {
	// Check if user is logged in
	session := controller.StartSession()
	userID := session.Get("UserID")
	beego.Info(fmt.Sprintf("UserID: %v", userID))

	if userID != nil {
		// UserID is set and can be deleted
		err := session.Delete("UserID")
		if err != nil {
			beego.Error(fmt.Sprintf("Не удалилась кука UserID: %v", err))

			// Сообщение об ошибке в браузер
			// Вывод сообщения об ошибке в модальном окне
			controller.TplName = "message-modal.tpl"
			controller.Data["title"] = "Ошибка"
			controller.Data["message1"] = "Ошибка"
			controller.Data["message2"] = "Ошибка разлогинивания - не удалилась кука UserID"
			controller.Data["message3"] = err
		} else {
			controller.Redirect("/realty/login", http.StatusSeeOther)
		}
	}
}
