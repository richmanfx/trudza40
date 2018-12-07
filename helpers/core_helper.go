package helpers

import (
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var db *sql.DB

// Соединение с БД и проверка соединения
func dbConnect() error {
	var err error
	//SetLogFormat()

	// Проверка соединения с БД
	log.Infof("Состояние соединения с БД перед подключением => db: '%v'", db)

	// Соединение с БД
	log.Infof("Подключение к БД")
	db, err = sql.Open(
		"postgres",
		"investor:realty2018@tcp(localhost:5432)/object_valuation?charset=utf8&parseTime=true&sslmode=disable")
	if err == nil {
		// Проверка соединения с БД
		log.Infof("Проверка соединения с БД после ")
		err = db.Ping()
	}
	if err != nil {
		log.Errorf("Ошибка подключения к БД: '%v'", err)
	}
	return err
}
