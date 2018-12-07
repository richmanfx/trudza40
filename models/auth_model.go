package main

import (
	"fmt"
	_ "github.com/astaxie/beego/migration"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"trudza40/helpers"
)

type User struct {
	Id       int
	Login    string `orm:"size(255)"`
	FullName string `orm:"size(255)"`
	Password string `orm:"size(511)"`
	Salt     string `orm:"size(511)"`
}

func init() {

	baseName, baseUserName, baseUserPassword := helpers.GetDbAccount()

	// Регистрация модели
	orm.RegisterModel(new(User))

	// Регистрация драйвера БД
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
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

	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Panic("Ошибка синхронизации базы данных PostgreSQL.")
	}
}
