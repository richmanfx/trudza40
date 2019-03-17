package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"net/http"
)

var globalSessions *session.Manager
var userSession session.Store

/* Инициализировать механизм сессий */
func SessionInit() {

	var (
		err                   error
		providerName          = "memory" // Сессии будем хранить в памяти, не в базе
		sessionProviderConfig session.ManagerConfig
	)

	//sessionProviderConfig.CookieName = "begoosessionID"
	sessionProviderConfig.Gclifetime = 3600

	globalSessions, err = session.NewManager(providerName, &sessionProviderConfig)
	if err != nil {
		beego.Error("Error create session manager")
	}

	go globalSessions.GC() // Почистить старые сессии
}

/* Приверить наличие сессии */
func CheckSession(controller *MainController) {

	userSession = controller.StartSession()
	userID := userSession.Get("UserID")
	sid := userSession.SessionID()

	beego.Info(fmt.Sprintf("В функции 'CheckSession' sid='%v' и userId='%v'", sid, userID))

	if userID == nil {
		// Пользователь ранее не логинился
		beego.Debug("Пользователь ранее не логинился")
		controller.Redirect("/realty/login", http.StatusSeeOther)
	}

}
