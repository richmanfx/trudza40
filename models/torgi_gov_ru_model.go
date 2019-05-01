package models

import (
	_ "github.com/astaxie/beego/migration"
	_ "github.com/lib/pq"
	"time"
)

// Строка результата скрапинга для одного объекта недвижимости
type ObjectScrapResult struct {
	OrderNumber              int       // Порядковый номер объекта
	NotificationNumber       string    // Номер извещения
	ProfitMargin             int       // Коэффициент доходности
	Address                  string    // Адрес объекта
	Area                     float32   // Площадь
	TradingDate              time.Time // Дата торгов
	GuaranteeAmount          float32   // Сумма залога
	LossFreeRental           float32   // Безубыточная сдача, руб/кв.м. в месяц
	YearRental               float32   // Выплаты ренты в год
	MonthlyRental            float32   // Выплаты ренты в месяц
	YearInsurance            int       // Страховка за год
	MonthlyCost              float32   // Расходы в месяц
	MonthlyHeating           int       // Стоимость отопления в месяц
	HousingOfficeMaintenance int       // Обслуживание ЖЭКом в месяц
	MonthlyProfit            float32   // Доход в месяц
	YearProfit               float32   // Доход в год
	PriorRepair              int       // Предварительный ремонт
}
