package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
)

type Settings struct {
	UserId        uint   `orm:"pk"`        // ID пользователя, первичный ключ
	SettingsName  string `orm:"size(255)"` // Название комплекта настроек
	BrowserWidth  uint   // Ширина окна браузера
	BrowserHeight uint   // Высота окна браузера
	HostPageUrl   string `orm:"size(255)"` // URL страницы хоста для скрапинга
	FlashQuantity uint   // Количество миганий
	FlashPeriod   uint   // Период мигания в мс
	FlashAllowed  bool   // True - включить подсветку, False - выключить
	DebugLevel    string `orm:"size(25)"` // Уровень отладочных сообщений

	MinArea          int    // Минимальная площадь объекта
	MaxArea          int    // Максимальная площадь объекта
	MinRentalPeriod  int    // Минимальный срок аренды, лет
	PropertyType     string `orm:"size(100)"` // Тип имущества (Zb: "Помещение")
	ContractType     string `orm:"size(100)"` // Вид договора (Zb: "Договор аренды")
	Country          string `orm:"size(100)"` // Страна (Zb: "РОССИЯ")
	PropertyLocation string `orm:"size(100)"` // Местоположение имущества (город) (Zb: "Москва (г)")
	SortFieldName    string `orm:"size(100)"` // Столбец, по которому сортировать ("Коэффициент доходности")

	AverageRental        int // Средняя стоимость аренды, рублей за кв.м. в месяц
	ProfitMonths         int // Количество доходных месяцев в году
	PriorRepair          int // Стоимость предварительного ремонта, рублей за кв.м.
	ContractRegistration int // Стоимость регистрации договора, рублей
	RunningCost          int // Мелкие расходы на запуск объекта
	YearlyInsurance      int // Стоимость годовой страховки в Альфа-Страховании, рубли
	// зависит от площади в метрах - до 100 кв.м = 4000 рублей
	MonthlyHeating           int // Отопление, рублей за кв.м. в месяц
	HousingOfficeMaintenance int // Обслуживание ЖЭКом, рублей за кв.м. в месяц
	AccountingService        int // Бухгалтерское обслуживание, рублей в месяц
	RequiredProfitMargin     int // Требуемый, приемлемый коэффициент доходности
}

func Init() {

	baseName, baseUserName, baseUserPassword := getDbAccount()

	//// Регистрация модели
	//orm.RegisterModel(new(Settings))

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

/* Получить настройки для "torgi.gov.ru" из БД */
func GetTorgiGovRuSettings(userId int) *Settings {

	o := orm.NewOrm()
	orm.Debug = false
	settings := new(Settings)
	var maps []orm.Params

	// Получить все настройки из базы
	num, err := o.QueryTable("settings").Filter("user_id", userId).Values(&maps)

	// TODO: Сделать обработку варианта, когда для данного пользователя ещё нет настроек в БД
	if num == 0 {
		beego.Info(fmt.Sprintf("Нет данных настроек в БД для пользователя с ID '%d'", userId))
	}

	if err != nil {
		beego.Error(fmt.Sprintf("Ошибка при получении данных настроек из БД: '%s'", err))
	}

	// Заполнить структуру настроек
	settings.UserId = uint(userId)
	SettingsFill(settings, maps)

	return settings

}

/* Заполнить структуру настроек */
func SettingsFill(settings *Settings, params []orm.Params) {

	for _, m := range params {

		settings.SettingsName = fmt.Sprintf("%v", m["SettingsName"])

		browserWidth, _ := strconv.Atoi(fmt.Sprintf("%v", m["BrowserWidth"]))
		settings.BrowserWidth = uint(browserWidth)

		browserHeight, _ := strconv.Atoi(fmt.Sprintf("%v", m["BrowserHeight"]))
		settings.BrowserHeight = uint(browserHeight)

		settings.HostPageUrl = fmt.Sprintf("%v", m["HostPageUrl"])

		flashQuantity, _ := strconv.Atoi(fmt.Sprintf("%v", m["FlashQuantity"]))
		settings.FlashQuantity = uint(flashQuantity)

		flashPeriod, _ := strconv.Atoi(fmt.Sprintf("%v", m["FlashPeriod"]))
		settings.FlashPeriod = uint(flashPeriod)

		settings.FlashAllowed, _ = strconv.ParseBool(fmt.Sprintf("%v", m["FlashAllowed"]))

		settings.DebugLevel = fmt.Sprintf("%v", m["DebugLevel"])

		minArea, _ := strconv.Atoi(fmt.Sprintf("%v", m["MinArea"]))
		settings.MinArea = int(minArea)

		maxArea, _ := strconv.Atoi(fmt.Sprintf("%v", m["MaxArea"]))
		settings.MaxArea = int(maxArea)

		minRentalPeriod, _ := strconv.Atoi(fmt.Sprintf("%v", m["MinRentalPeriod"]))
		settings.MinRentalPeriod = int(minRentalPeriod)

		settings.PropertyType = fmt.Sprintf("%v", m["PropertyType"])

		settings.ContractType = fmt.Sprintf("%v", m["ContractType"])

		settings.Country = fmt.Sprintf("%v", m["Country"])

		settings.PropertyLocation = fmt.Sprintf("%v", m["PropertyLocation"])

		settings.SortFieldName = fmt.Sprintf("%v", m["SortFieldName"])

		averageRental, _ := strconv.Atoi(fmt.Sprintf("%v", m["AverageRental"]))
		settings.AverageRental = int(averageRental)

		profitMonths, _ := strconv.Atoi(fmt.Sprintf("%v", m["ProfitMonths"]))
		settings.ProfitMonths = int(profitMonths)

		priorRepair, _ := strconv.Atoi(fmt.Sprintf("%v", m["PriorRepair"]))
		settings.PriorRepair = int(priorRepair)

		contractRegistration, _ := strconv.Atoi(fmt.Sprintf("%v", m["ContractRegistration"]))
		settings.ContractRegistration = int(contractRegistration)

		runningCost, _ := strconv.Atoi(fmt.Sprintf("%v", m["RunningCost"]))
		settings.RunningCost = int(runningCost)

		yearlyInsurance, _ := strconv.Atoi(fmt.Sprintf("%v", m["YearlyInsurance"]))
		settings.YearlyInsurance = int(yearlyInsurance)

		monthlyHeating, _ := strconv.Atoi(fmt.Sprintf("%v", m["MonthlyHeating"]))
		settings.MonthlyHeating = int(monthlyHeating)

		housingOfficeMaintenance, _ := strconv.Atoi(fmt.Sprintf("%v", m["HousingOfficeMaintenance"]))
		settings.HousingOfficeMaintenance = int(housingOfficeMaintenance)

		accountingService, _ := strconv.Atoi(fmt.Sprintf("%v", m["AccountingService"]))
		settings.AccountingService = int(accountingService)

		requiredProfitMargin, _ := strconv.Atoi(fmt.Sprintf("%v", m["RequiredProfitMargin"]))
		settings.RequiredProfitMargin = int(requiredProfitMargin)
	}

	//beego.Info(fmt.Sprintf("Настройки: '%v'", settings))
}
