package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	// Регистрация модели
	orm.RegisterModel(new(User))

	// Регистрация драйвера БД
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		log.Panic("Ошибка регистрации драйвера PostgreSQL.")
	}

	// БД по умолчанию
	err = orm.RegisterDataBase(
		"default",
		"postgres",
		"user=investor password=realty2018 host=127.0.0.1 port=5432 dbname=object_valuation sslmode=disable")
	if err != nil {
		log.Panic("Ошибка регистрации базы данных PostgreSQL.")
	}

	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Panic("Ошибка синхронизации базы данных PostgreSQL.")
	}
}
