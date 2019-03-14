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

	MinArea          uint   // Минимальная площадь объекта
	MaxArea          uint   // Максимальная площадь объекта
	MinRentalPeriod  uint   // Минимальный срок аренды, лет
	PropertyType     string `orm:"size(100)"` // Тип имущества (Zb: "Помещение")
	ContractType     string `orm:"size(100)"` // Вид договора (Zb: "Договор аренды")
	Country          string `orm:"size(100)"` // Страна (Zb: "РОССИЯ")
	PropertyLocation string `orm:"size(100)"` // Местоположение имущества (город) (Zb: "Москва (г)")
	SortFieldName    string `orm:"size(100)"` // Столбец, по которому сортировать ("Коэффициент доходности")

	AverageRental        uint // Средняя стоимость аренды, копеек за кв.м. в месяц
	ProfitMonths         uint // Количество доходных месяцев в году
	PriorRepair          uint // Предварительный ремонт, копеек за кв.м.
	ContractRegistration uint // Стоимость регистрации договора, копеек
	RunningCost          uint // Мелкие расходы на запуск объекта
	YearlyInsurance      uint // Стоимость годовой страховки (копейки) в Альфа-Страховании,
	// зависит от площади в метрах - до 100 кв.м = 4000 копеек
	MonthlyHeating           uint // Отопление, копеек за кв.м. в месяц
	HousingOfficeMaintenance uint // Обслуживание ЖЭКом, копеек за кв.м. в месяц
	AccountingService        uint // Бухгалтерское обслуживание, копеек в месяц
	RequiredProfitMargin     uint // Требуемый, приемлемый коэффициент доходности
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
	orm.Debug = true
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

// Заполнить структуру настроек
func SettingsFill(settings *Settings, params []orm.Params) {

	for _, m := range params {

		settings.SettingsName = fmt.Sprintf("%v", m["SettingsName"])

		browserWidth, _ := strconv.Atoi(fmt.Sprintf("%v", m["BrowserWidth"]))
		settings.BrowserWidth = uint(browserWidth)

		browserHeight, _ := strconv.Atoi(fmt.Sprintf("%v", m["BrowserHeight"]))
		settings.BrowserHeight = uint(browserHeight)

		settings.DebugLevel = fmt.Sprintf("%v", m["DebugLevel"])

		settings.HostPageUrl = fmt.Sprintf("%v", m["HostPageUrl"])

		flashQuantity, _ := strconv.Atoi(fmt.Sprintf("%v", m["FlashQuantity"]))
		settings.FlashQuantity = uint(flashQuantity)

		flashPeriod, _ := strconv.Atoi(fmt.Sprintf("%v", m["FlashPeriod"]))
		settings.FlashPeriod = uint(flashPeriod)

		settings.FlashAllowed, _ = strconv.ParseBool(fmt.Sprintf("%v", m["FlashAllowed"]))

		minArea, _ := strconv.Atoi(fmt.Sprintf("%v", m["MinArea"]))
		settings.MinArea = uint(minArea)

		maxArea, _ := strconv.Atoi(fmt.Sprintf("%v", m["MaxArea"]))
		settings.MaxArea = uint(maxArea)

		minRentalPeriod, _ := strconv.Atoi(fmt.Sprintf("%v", m["MinRentalPeriod"]))
		settings.MinRentalPeriod = uint(minRentalPeriod)

		settings.PropertyType = fmt.Sprintf("%v", m["PropertyType"])

		settings.ContractType = fmt.Sprintf("%v", m["ContractType"])

		settings.Country = fmt.Sprintf("%v", m["Country"])

		settings.PropertyLocation = fmt.Sprintf("%v", m["PropertyLocation"])

		settings.SortFieldName = fmt.Sprintf("%v", m["SortFieldName"])

		averageRental, _ := strconv.Atoi(fmt.Sprintf("%v", m["AverageRental"]))
		settings.AverageRental = uint(averageRental)

		profitMonths, _ := strconv.Atoi(fmt.Sprintf("%v", m["ProfitMonths"]))
		settings.ProfitMonths = uint(profitMonths)

	}

	print()
}
