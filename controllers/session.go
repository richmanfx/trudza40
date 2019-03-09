package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"net/http"
)

var globalSessions *session.Manager

/* Инициализировать мешаизм сессий */
func SessionInit() {

	var (
		err                   error
		providerName          = "memory" // Сессии будем хранить в памяти, не в базе
		sessionProviderConfig session.ManagerConfig
	)

	sessionProviderConfig.CookieName = "gosessionid"
	sessionProviderConfig.Gclifetime = 3600

	globalSessions, err = session.NewManager(providerName, &sessionProviderConfig)
	if err != nil {
		beego.Error("Error create session manager")
	}

	go globalSessions.GC() // Почистить старые сессии
}

/* Приверить наличие сессии */
func CheckSession(controller *MainController) {

	userSession := controller.GetSession("userSession")
	if userSession == nil {
		//controller.SetSession("userSession", int(1))
		//controller.Data["sessid"] = 0

		// Если нет сессии, то редирект на страницу авторизации
		//controller.Abort("303")
		controller.Redirect("/realty/login", http.StatusSeeOther)

	} else {
		// Если сессия есть, то инкрементировать её номер
		controller.SetSession("userSession", userSession.(int)+1)
		controller.Data["sessid"] = userSession.(int)
	}
}
