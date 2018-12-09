package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"trudza40/models"
)

type ConfigController struct {
	beego.Controller
}

/* Редактирование основной конфигурации */
func (context *ConfigController) EditConfig() {
	context.TplName = "config.tpl"
	context.Data["title"] = "Realty Config"
}

/* Конфигурирование пользователей */
func (context *ConfigController) UsersConfig() {

	// Считать из БД пользователей
	users, err := models.GetUsers()
	beego.Info(fmt.Sprintf("Пользователи из БД: '%v'", users))

	if err == nil {
		context.TplName = "users-config.tpl"
		context.Data["title"] = "Users Config"
		context.Data["users"] = users
	} else {
		beego.Error("Ошибка при получении данных о пользователях из БД")
		// Вывод сообщения об ошибке в модальном окне
		context.TplName = "message-modal.tpl"
		context.Data["title"] = "Ошибка"
		context.Data["message1"] = "Ошибка"
		context.Data["message2"] = "Ошибка при получении данных о пользователях из БД"
		context.Data["message3"] = err
	}
}

/* Ввести данные нового пользователя в модальной форме */
func (context *ConfigController) CreateUser() {
	context.TplName = "create-user-modal.tpl"
	context.Data["title"] = "Create User"
}

/* Создать нового пользователя в БД  */
func (context *ConfigController) CreateUserInDb() {

	// Пользователь
	user := new(models.User)

	// Данные из формы
	user.Login = context.GetString("login")
	user.FullName = context.GetString("full_name")
	openPassword := context.GetString("password")

	// Получить Соль и Хеш пароля
	user.Salt = models.CreateSalt()
	user.Password = models.CreateHash(openPassword, user.Salt)

	err := models.CreateUserInDbProcessing(*user)

	if err == nil {
		beego.Info("Пользователь удачно создан")
		// Вывод сообщения об удачном создании пользователя
		context.TplName = "message-modal.tpl"
		context.Data["title"] = "Info"
		context.Data["message1"] = "Информация"
		context.Data["message2"] = fmt.Sprintf("Пользователь '%s' удачно создан", user.Login)

	} else {
		beego.Error("Создать пользователя не удалось")
		// Вывод сообщения об ошибке в модальном окне
		context.TplName = "message-modal.tpl"
		context.Data["title"] = "Ошибка"
		context.Data["message1"] = "Ошибка"
		context.Data["message2"] = "Создать пользователя не удалось"
		context.Data["message3"] = err
	}
}
