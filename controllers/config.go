package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"trudza40/models"
)

type ConfigController struct {
	beego.Controller
}

/* Редактирование основной конфигурации */
func (controller *ConfigController) EditConfig() {
	controller.TplName = "config.tpl"
	controller.Data["title"] = "Realty Config"
}

/* Конфигурирование пользователей */
func (controller *ConfigController) UsersConfig() {

	// Считать из БД пользователей
	users, err := models.GetUsers()
	beego.Info(fmt.Sprintf("Пользователи из БД: '%v'", users))

	if err == nil {
		controller.TplName = "users-config.tpl"
		controller.Data["title"] = "Users Config"
		controller.Data["GlobalUserLogin"] = GlobalUserLogin
		controller.Data["users"] = users
	} else {
		beego.Error("Ошибка при получении данных о пользователях из БД")
		// Вывод сообщения об ошибке в модальном окне
		controller.TplName = "message-modal.tpl"
		controller.Data["title"] = "Ошибка"
		controller.Data["message1"] = "Ошибка"
		controller.Data["message2"] = "Ошибка при получении данных о пользователях из БД"
		controller.Data["message3"] = err
	}
}

/* Ввести данные нового пользователя в модальной форме */
func (controller *ConfigController) CreateUser() {
	controller.TplName = "create-user-modal.tpl"
	controller.Data["title"] = "Create User"
}

/* Создать нового пользователя в БД */
func (controller *ConfigController) CreateUserInDb() {

	// Пользователь
	user := new(models.User)

	// Данные из формы
	user.Login = controller.GetString("login")
	user.FullName = controller.GetString("full_name")
	openPassword := controller.GetString("password")

	// Получить Соль и Хеш пароля
	user.Salt = models.CreateSalt()
	user.Password = models.CreateHash(openPassword, user.Salt)

	err := models.CreateUserInDbProcessing(*user)

	if err == nil {
		beego.Info("Пользователь удачно создан")
		// Вывод сообщения об удачном создании пользователя
		controller.TplName = "message-modal.tpl"
		controller.Data["title"] = "Info"
		controller.Data["message1"] = "Информация"
		controller.Data["message2"] = fmt.Sprintf("Пользователь '%s' удачно создан", user.Login)

	} else {
		beego.Error("Создать пользователя не удалось")
		// Вывод сообщения об ошибке в модальном окне
		controller.TplName = "message-modal.tpl"
		controller.Data["title"] = "Ошибка"
		controller.Data["message1"] = "Ошибка"
		controller.Data["message2"] = "Создать пользователя не удалось"
		controller.Data["message3"] = err
	}
}

/* Удалить пользователя */
func (controller *ConfigController) DeleteUser() {

	// Пользователь
	user := new(models.User)

	// Данные из формы
	user.Login = controller.GetString("login")
	user.FullName = controller.GetString("full_name")

	err := models.DeleteUserInDb(*user)

	if err == nil {
		beego.Info("Пользователь удачно удалён")

		controller.TplName = "message-modal.tpl"
		controller.Data["title"] = "Info"
		controller.Data["message1"] = "Информация"
		controller.Data["message2"] = fmt.Sprintf("Пользователь '%s' удачно удалён", user.Login)

	} else {
		beego.Error("Удалить пользователя не удалось")

		controller.TplName = "message-modal.tpl"
		controller.Data["title"] = "Ошибка"
		controller.Data["message1"] = "Ошибка"
		controller.Data["message2"] = "Удалить пользователя не удалось"
		controller.Data["message3"] = err
	}
}

/* Изменить пароль пользователя */
func (controller *ConfigController) ChangePassword() {

	// Пользователь
	user := new(models.User)

	// Данные из формы
	user.Login = controller.GetString("login")
	user.FullName = controller.GetString("full_name")
	user.Id, _ = strconv.Atoi(controller.GetString("id"))
	user.Password = controller.GetString("new_password")

	// Записать в БД новый пароль
	err := models.SavePassword(*user)

	if err == nil {
		beego.Info("Пароль пользователя удачно изменён")
		// Вывод сообщения об удачном создании пользователя
		controller.TplName = "message-modal.tpl"
		controller.Data["title"] = "Info"
		controller.Data["message1"] = "Информация"
		controller.Data["message2"] = fmt.Sprintf("Пароль пользователя '%s' удачно изменён", user.Login)

	} else {
		beego.Error("Изменить пароль пользователя не удалось")
		// Вывод сообщения об ошибке в модальном окне
		controller.TplName = "message-modal.tpl"
		controller.Data["title"] = "Ошибка"
		controller.Data["message1"] = "Ошибка"
		controller.Data["message2"] = "Изменить пароль пользователя не удалось"
		controller.Data["message3"] = err
	}

}
