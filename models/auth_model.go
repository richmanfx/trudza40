package models

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/migration"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/ini.v1"
	"io"
	"log"
	"strconv"
	"time"
)

type User struct {
	Id       int    `orm:"pk"`
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

/* Проверить наличие пользователя в БД */
func CheckUserInDB(login string) (int, error) {
	beego.Info("Работает функция 'CheckUserInDB'")

	var err error
	o := orm.NewOrm() // Использовать ORM "Ormer"
	orm.Debug = true  // Логирование ORM запросов

	user := User{Login: login}
	exist := o.QueryTable(user).Filter("login", login).Exist() // Существует ли в базе?

	if exist {
		beego.Info(fmt.Sprintf("Пользователь '%s' существует", user.Login))
		// Получить ID пользователя из база
		err := o.QueryTable("user").Filter("login", login).One(&user, "id") // Только ID интересует
		if err != nil {
			beego.Error("Ошибка при запросе получения из базы ID пользователя")
		}
		beego.Info(fmt.Sprintf("ID пользователя: %d", user.Id))
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
	return user.Id, err
}

/* Вернуть имя БД, имя пользователя в БД и его пароль */
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

/* Получить параметры из конфигурационного INI файла */
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

/* Проверить пароль по Хешу из БД */
func CheckPasswordInDB(login, password string) error {

	// Получить Соль из БД
	salt, err := GetSaltFromDb(login)
	beego.Info(fmt.Sprintf("Соль из БД: '%s'", salt))

	if err == nil {

		// Сгенерить Хеш пароля с Солью
		newHash := CreateHash(password, salt)
		beego.Info(fmt.Sprintf("Хеш с Солью: '%s'", newHash))

		// Считать Хеш из БД
		var oldHash string
		oldHash, err = GetHashFromDb(login)
		if err == nil {
			err = bcrypt.CompareHashAndPassword([]byte(oldHash), []byte(password))
		}

		if err == nil {
			beego.Info("Хеш пароля совпадает с Хешем из БД")
		} else {
			beego.Info("Хеш пароля не совпадает с Хешем из БД")
			err = errors.New(fmt.Sprintln("Неверный логин/пароль"))
		}
	}
	if err != nil {
		beego.Info(fmt.Sprintf("Ошибка при проверке пароля по Хешу из БД: '%v'", err))
	}
	return err
}

/* Получить "соль" из БД для заданного пользователя */
func GetSaltFromDb(userLogin string) (string, error) {

	o := orm.NewOrm() // Использовать ORM "Ormer"
	orm.Debug = true  // Логирование ORM запросов

	// Получить "соль"
	user := User{Login: userLogin}
	err := o.QueryTable("user").Filter("login", userLogin).One(&user, "salt") // Только Salt интересует

	if err == orm.ErrMultiRows {
		fmt.Printf("Returned Multi Rows Not One") // Have multiple records
	}
	if err == orm.ErrNoRows {
		fmt.Printf("Not row found") // No result
	}

	salt := user.Salt

	//defer db.Close()		// TODO: Пока не знаю закрывать ли...
	if err != nil {
		beego.Info(fmt.Sprintf("Ошибка получения 'соли' для пользователя с логином '%s': %v", userLogin, err))
	}
	return salt, err
}

/* Получить Хеш пароля с заданной солью */
func CreateHash(password string, salt string) string {
	intSalt, _ := strconv.Atoi(salt)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), intSalt)
	return string(hashedPassword)
}

/* Получить хеш из БД для заданного пользователя */
func GetHashFromDb(userLogin string) (string, error) {

	o := orm.NewOrm() // Использовать ORM "Ormer"
	orm.Debug = true  // Логирование ORM запросов

	// Получить "Хеш"
	user := User{Login: userLogin}
	err := o.QueryTable("user").Filter("login", userLogin).One(&user, "password") // Только Хеш пароля интересует

	if err == orm.ErrMultiRows {
		fmt.Printf("Returned Multi Rows Not One") // Have multiple records
	}
	if err == orm.ErrNoRows {
		fmt.Printf("Not row found") // No result
	}

	hash := user.Password

	//defer db.Close()		// TODO:
	if err != nil {
		beego.Info(
			fmt.Sprintf("Ошибка получения из базы Хеша пароля для пользователя с логином '%s': %v", userLogin, err))
	}
	return hash, err
}

/* Создать пользователя в БД */
func CreateUserInDbProcessing(user User) error {

	o := orm.NewOrm() // Использовать ORM "Ormer"
	orm.Debug = true  // Логирование ORM запросов

	id, err := o.Insert(&user)
	if err == nil {
		beego.Info(fmt.Sprintf("Record Id: '%d'", id))
	}

	// 	defer db.Close()		TODO

	if err != nil {
		beego.Error(fmt.Sprintf("Ошибка при создании пользователя в БД: '%v'", err))
	}
	return err
}

/* Сгенерировать "соль" */
func CreateSalt() string {
	hash := sha512.New()
	_, err := io.WriteString(hash, time.Now().String())

	if err != nil {
		beego.Error("Ошибка при генерации соли")
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}

/* Считать из БД всех пользователей */
func GetUsers() ([]User, error) {

	o := orm.NewOrm() // Использовать ORM "Ormer"
	orm.Debug = true  // Логирование ORM запросов

	var usersList []User

	// Считать из БД
	_, err := o.QueryTable("user").OrderBy("Id").All(&usersList)

	//defer db.Close()		// TODO:
	if err != nil {
		beego.Error(fmt.Sprintf("Ошибка при считывании из БД всех пользователей: '%v'", err))
	}
	return usersList, err
}

/* Удалить пользователя из БД */
func DeleteUserInDb(user User) error {

	o := orm.NewOrm() // Использовать ORM "Ormer"
	orm.Debug = true  // Логирование ORM запросов

	id, err := o.QueryTable("user").Filter("login", user.Login).Filter("full_name", user.FullName).Delete()
	if err == nil {
		beego.Info(fmt.Sprintf("Deleted record ID: '%d'", id))
	}

	//defer db.Close()		// TODO
	if err != nil {
		beego.Error(fmt.Sprintf("Ошибка при удалении пользователя '%s' в БД: '%v'", user.Login, err))
	}
	return err
}

/* Записать в БД новый пароль заданного пользователя */
func SavePassword(user User) error {

	o := orm.NewOrm() // Использовать ORM "Ormer"
	orm.Debug = true  // Логирование ORM запросов

	// Сгенерировать новую соль
	salt := CreateSalt()
	beego.Info(fmt.Sprintf("Новая Соль: '%s'", salt))

	// Сгенерить Хеш пароля с Солью
	newHash := CreateHash(user.Password, salt)
	beego.Info(fmt.Sprintf("Хеш с Солью: '%s'", newHash))

	// Занести новый хеш пароля и новую соль в БД
	id, err :=
		o.QueryTable("user").
			Filter("id", user.Id).
			Update(orm.Params{"password": newHash, "salt": salt})
	if err == nil {
		beego.Info(fmt.Sprintf("Updated record ID: '%d'", id))
	}

	if err != nil {
		beego.Error(fmt.Sprintf("Ошибка записи в БД нового пароля пользователя: '%v'", err))
	}
	return err
}
