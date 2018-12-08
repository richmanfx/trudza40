package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/migration"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"gopkg.in/ini.v1"
	"log"
)

type User struct {
	Id       int
	Login    string `orm:"size(255)"`
	FullName string `orm:"size(255)"`
	Password string `orm:"size(511)"`
	Salt     string `orm:"size(511)"`
}

func init() {

	baseName, baseUserName, baseUserPassword := getDbAccount()

	// Регистрация модели
	orm.RegisterModel(new(User))

	// Регистрация драйвера БД
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		beego.Error(fmt.Sprintf("Ошибка регистрации драйвера PostgreSQL: %v", err))
	}

	// БД по умолчанию
	dataSourceString := fmt.Sprintf("user=%s password=%s host=127.0.0.1 port=5432 dbname=%s sslmode=disable",
		baseUserName, baseUserPassword, baseName)
	err = orm.RegisterDataBase("default", "postgres", dataSourceString)
	if err != nil {
		log.Panicf("Ошибка регистрации базы данных PostgreSQL: %v", err)
	}

	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Panicf("Ошибка синхронизации базы данных PostgreSQL: %v", err)
	}
}

// Проверить наличие пользователя в БД
func CheckUserInDB(login string) error {
	beego.Info("Работает функция 'CheckUserInDB'")

	var err error
	var o orm.Ormer

	o = orm.NewOrm() // Использовать ORM "Ormer"
	orm.Debug = true // Логирование ORM запросов

	user := User{Login: login}
	exist := o.QueryTable(user).Filter("login", login).Exist() // Существует ли в базе?

	if exist {
		beego.Info(fmt.Sprintf("Пользователь существует: %d / %s / %s", user.Id, user.Login, user.FullName))
	} else {
		err = errors.New(fmt.Sprintf("Пользователь '%s' в БД не существует", login))
	}

	if err == orm.ErrNoRows {
		beego.Error("Отсутствует результат запроса")
	} else if err == orm.ErrMissPK {
		beego.Error("Отсутствует 'primary key'")
	}

	//defer db.Close()		// TODO: Пока не закрываем - не ясно как

	if err != nil {
		beego.Error(fmt.Sprintf("Ошибка при проверке наличия пользователя '%s' в БД: '%v'", login, err))
	}
	return err
}

// Вернуть имя БД, имя пользователя в БД и его пароль
func getDbAccount() (baseName, baseUserName, baseUserPassword string) {

	// Файл аккаунтов
	const (
		accountDirName  = "/usr/local/etc"
		accountFileName = "accounts.ini"
	)

	// Полное имя файла аккаунтов
	fullAccountFileName := accountDirName + "/" + accountFileName
	beego.Info(fmt.Sprintf("Full config file name: '%s'", fullAccountFileName))

	// Чтение параметров из файла аккаунтов
	getConfigParameters(fullAccountFileName, &baseName, &baseUserName, &baseUserPassword)

	return baseName, baseUserName, baseUserPassword
}

// Получить параметры из конфигурационного INI файла
func getConfigParameters(fullConfigFileName string, baseName, baseUserName, baseUserPassword *string) {

	config, err := ini.Load(fullConfigFileName)
	if err != nil {
		log.Panicf("Fail to read config file: %v", err)
	}

	*baseName = config.Section("").Key("DATABASENAME").String()
	beego.Info(fmt.Sprintf("Используемая база данных: '%s'", *baseName))

	*baseUserName = config.Section("").Key("BASEUSERNAME").String()
	*baseUserPassword = config.Section("").Key("BASEUSERPASSWORD").String()
}
