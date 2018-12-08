package models

import (
	"fmt"
	_ "github.com/astaxie/beego/migration"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
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
		log.Panicf("Ошибка регистрации драйвера PostgreSQL: %v", err)
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

	log.Info("Работает функция 'CheckUserInDB'")

	var err error
	//var loginFromDB string

	// Считать из БД
	//requestResult := db.QueryRow("SELECT login FROM user WHERE login=?", login)

	var o orm.Ormer
	o = orm.NewOrm()

	user := User{Login: login}
	err = o.Read(&user)
	fmt.Printf("ERR: %v\n", err)

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

// Вернуть имя БД, имя пользователя в БД и его пароль
func getDbAccount() (baseName, baseUserName, baseUserPassword string) {

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
