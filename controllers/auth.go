package controllers

import (
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

/* Форма для авторизации */
func (controller *AuthController) Login() {
	controller.TplName = "login.tpl"
	controller.Data["title"] = "Realty Login"
}

/* Обработка логина\пароля из формы авторизации */
func (controller *AuthController) LoginProcessing() {

	var err error
	var userName string
	var userPassword string

	userName = controller.GetString("user_name")
	userPassword = controller.GetString("user_password")

	// Проверить существование пользователя в базе
	GlobalUserId, err = models.CheckUserInDB(userName)

	// Проверить пароль по Хешу из БД
	if err == nil {
		// Можно закомментировать для возможности залогиниться без Хешей из БД
		err = models.CheckPasswordInDB(userName, userPassword)

		if err == nil {

			beego.Info(fmt.Sprintf("Пользователь '%s' вошёл в приложение.", userName))

			// Логин пользователя для заголовка
			GlobalUserLogin = userName

			//////// Добавить куку с ID пользователя
			////userSession = controller.StartSession()
			////username := userSession.Get("username")
			//
			//// Создаёт или читает "session id" из HTTP запроса
			//// Если "session id" существует, то возвращает SessionStore с этим ID
			//userSession, _ = globalSessions.SessionStart(w, r)
			//defer userSession.SessionRelease(w)
			////username := userSession.Get("username")
			//if r.Method == "GET" {
			//	controller.Redirect("/realty", http.StatusSeeOther)
			//} else {
			//	err = userSession.Set("username", r.Form["username"])
			//	//err = userSession.Set("username", username)
			//	//err = userSession.Set("username", userName)
			//
			//	if err != nil {
			//		beego.Error("Не добавился username в сессию")
			//	} else {
			//		beego.Info("Добавлен username в сессию")
			//	}
			//}

			//err = userSession.Set("username", r.Form["username"])
			//if err != nil {
			//	beego.Error("Не добавился username в сессию")
			//}

			//err = userSession.Set("UserID", GlobalUserId) // TODO: Здесь валится
			//time.Sleep(1 * time.Second)
			//if err != nil {
			//	beego.Error("Не добавился UserID в сессию")
			//}
			//beego.Info("Добавлен UserID")

			// Новая сессия с новой кукой
			//userID := controller.StartSession().Get("UserID") // Новая сессия
			//
			//beego.Info(fmt.Sprintf("UserID: %v", GlobalUserId))
			//if userID == nil {
			//	beego.Info("userID = nil!")
			//	return
			//}

			//// 3. Получить существующую или новую сессию
			session, err := store.Get(controller.Ctx.Request, "session-name")
			if err != nil {
				http.Error(controller.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
				return
			}

			// Установить некоторые значения для сессии
			session.Values["UserID"] = GlobalUserId

			// Сохранить сессию в ответе сервера
			err = session.Save(controller.Ctx.Request, controller.Ctx.ResponseWriter)
			if err != nil {
				http.Error(controller.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
				return
			}

			//userSession, _ = globalSessions.SessionStart(controller.Ctx.ResponseWriter, controller.Ctx.Request)
			//beego.Info("userSession: ", userSession)
			//defer userSession.SessionRelease(controller.Ctx.ResponseWriter)
			//
			//store, err := globalSessions.GetSessionStore(userSession.SessionID())
			//if err != nil {
			//	beego.Info(fmt.Sprintf("Ошибка globalSessions.GetSessionStore: '%v'", err))
			//}
			//
			//userId := userSession.Get("UserID")
			//if controller.Ctx.Request.Method == "GET" {
			//	t, _ := template.ParseFiles("realty.tpl")
			//	err = t.Execute(controller.Ctx.ResponseWriter, nil)
			//	if err != nil {
			//		beego.Info(fmt.Sprintf("Ошибка t.Execute: '%v'", err))
			//	}
			//} else {
			//	err = store.Set("UserID", userId)
			//	if err != nil {
			//		beego.Info(fmt.Sprintf("Ошибка store.Set: '%v'", err))
			//	}
			//	store.SessionRelease(controller.Ctx.ResponseWriter)
			//	//userSession.SessionRelease(controller.Ctx.ResponseWriter)
			//}
			//

			// 4.
			// Получить из сессии значение UserID
			//userId := controller.GetSession("UserID")
			//beego.Info(fmt.Sprintf("Из сессии userId='%v'", userId))
			//if userId == nil {	// Если нет UserID в сессии, то выставить
			//	controller.SetSession("UserID", GlobalUserId)
			//	beego.Info(fmt.Sprintf("Выставили в сессии userId='%v'", GlobalUserId))
			//}

			//beego.Info("Редирект на '/realty'")
			//controller.Redirect("/realty", http.StatusSeeOther)
			controller.TplName = "realty.tpl"
			controller.Data["title"] = "Realty"
			controller.Data["GlobalUserLogin"] = GlobalUserLogin
		}

	}

	if err != nil {
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

	session, err := store.Get(controller.Ctx.Request, "session-name")
	if err != nil {
		http.Error(controller.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	// Проверить залогинен ли юзер
	userID := session.Values["UserID"]
	beego.Info(fmt.Sprintf("Logout. UserID: %v", userID))

	if userID != nil {
		// UserID присутствует - удаляем
		//session.Flashes()
		session.Values["UserID"] = ""
	}

	controller.TplName = "login.tpl"
}
