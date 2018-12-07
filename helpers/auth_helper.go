package helpers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

// Вернуть имя БД, имя пользователя в БД и его пароль
func GetDbAccount() (baseName, baseUserName, baseUserPassword string) {

	// Файл аккаунтов
	const (
		accountDirName  = "/usr/local/etc"
		accountFileName = "accounts.ini"
	)

	// Полное имя файла аккаунтов
	fullAccountFileName := accountDirName + "/" + accountFileName
	log.Debugf("Full config file name: %s", fullAccountFileName)

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
	log.Infof("Используемая база данных: %s", *baseName)

	*baseUserName = config.Section("").Key("BASEUSERNAME").String()
	*baseUserPassword = config.Section("").Key("BASEUSERPASSWORD").String()
}

// Проверить наличие пользователя в БД
func CheckUserInDB(login string) error {

	log.Info("Работает функция 'CheckUserInDB'")

	var err error
	//var loginFromDB string

	baseName, baseUserName, baseUserPassword := GetDbAccount()

	// Регистрация драйвера БД
	err = orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		log.Panic("Ошибка регистрации драйвера PostgreSQL.")
	}

	// БД по умолчанию
	dataSourceString := fmt.Sprintf("user=%s password=%s host=127.0.0.1 port=5432 dbname=%s sslmode=disable",
		baseUserName, baseUserPassword, baseName)
	err = orm.RegisterDataBase("default", "postgres", dataSourceString)
	if err != nil {
		log.Panic("Ошибка регистрации базы данных PostgreSQL.")
	}

	// TODO: Работаем здесь!!!

	//// Подключиться к БД
	//err = dbConnect()
	//if err == nil {
	//
	//	// Считать из БД
	//	requestResult := db.QueryRow("SELECT login FROM user WHERE login=?", login)
	//
	//	err = requestResult.Scan(&loginFromDB)
	//
	//	if err == nil {
	//		log.Debugf("Пользователь '%s' существует", login)
	//	} else {
	//		err = errors.New(fmt.Sprintf("Пользователь '%s' в БД не существует", login))
	//	}
	//}
	//defer db.Close()

	if err != nil {
		log.Errorf("Ошибка при проверке наличия пользователя '%s' в БД: '%v'", login, err)
	}
	return err
}
