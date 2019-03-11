package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
)

type TorgiGovRuSettings struct {
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

	AverageRental        float32 // Средняя стоимость аренды, рублей за кв.м. в месяц
	ProfitMonths         uint    // Количество доходных месяцев в году
	PriorRepair          float32 // Предварительный ремонт, рублей за кв.м.
	ContractRegistration float32 // Стоимость регистрации договора, рублей
	RunningCost          float32 // Мелкие расходы на запуск объекта
	YearlyInsurance      float32 // Стоимость годовой страховки (рубли) в Альфа-Страховании,
	// зависит от площади в метрах - до 100 кв.м = 4000 рублей
	MonthlyHeating           float32 // Отопление, рублей за кв.м. в месяц
	HousingOfficeMaintenance float32 // Обслуживание ЖЭКом, рублей за кв.м. в месяц
	AccountingService        float32 // Бухгалтерское обслуживание, рублей в месяц
	RequiredProfitMargin     uint    // Требуемый, приемлемый коэффициент доходности
}

func Init() {

	baseName, baseUserName, baseUserPassword := getDbAccount()

	// Регистрация модели
	orm.RegisterModel(new(TorgiGovRuSettings))

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
