package models

import (
	_ "github.com/astaxie/beego/migration"
	_ "github.com/lib/pq"
)

// Строка результата скрапинга для одного объекта недвижимости
type ObjectScrapResult struct {
	OrderNumber              int     // Порядковый номер объекта
	NotificationNumber       string  // Номер извещения
	ProfitMargin             int     // Коэффициент доходности
	Address                  string  // Адрес объекта
	Area                     float64 // Площадь
	AuctionData              string  // Дата торгов
	GuaranteeAmount          float32 // Сумма залога
	LossFreeRental           int     // Безубыточная сдача, руб/кв.м. в месяц
	YearRental               int     // Выплаты ренты в год
	MonthlyRental            int     // Выплаты ренты в месяц
	YearInsurance            int     // Страховка за год
	MonthlyCost              int     // Расходы в месяц
	MonthlyHeating           int     // Стоимость отопления в месяц
	HousingOfficeMaintenance int     // Обслуживание ЖЭКом в месяц
	MonthlyProfit            int     // Доход в месяц
	YearProfit               int     // Доход в год
	PriorRepair              int     // Предварительный ремонт
	WebLink                  string  // Ссылка для просмотра
}
