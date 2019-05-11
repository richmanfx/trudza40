package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/sessions"
	"net/http"
)

//var globalSessions *session.Manager
//var userSession session.Store

// Инициализация хранилища сессий секретным ключом
var store = sessions.NewCookieStore([]byte("Здесь-секретный-ключ_вынести_его"))

/* Инициализировать механизм сессий */
func SessionInit() {

	//var (
	//	err                   error
	//	providerName          = "memory" // Сессии будем хранить в памяти, не в базе
	//	sessionProviderConfig session.ManagerConfig
	//)
	//
	////sessionProviderConfig.CookieName = "begoosessionID"
	//sessionProviderConfig.Gclifetime = 3600
	//
	//// Инициализировать данные
	//globalSessions, err = session.NewManager(providerName, &sessionProviderConfig)
	//if err != nil {
	//	beego.Error("Error create session manager")
	//}
	//
	//// Почистить старые сессии
	//go globalSessions.GC()

}

/* Приверить наличие сессии */
func CheckSession(controller *MainController) {

	session, err := store.Get(controller.Ctx.Request, "session-name")
	if err != nil {
		http.Error(controller.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	if session.ID == "" {
		// Пользователь ранее не логинился
		beego.Info("Пользователь ранее не логинился")

		//controller.TplName = "/realty/login"
		controller.Redirect("/realty/login", http.StatusSeeOther)
	}

	beego.Info(fmt.Sprintf("В функции 'CheckSession' sid='%v'", session.ID))

}
