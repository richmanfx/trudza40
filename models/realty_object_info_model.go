package models

type ObjectInfo struct {
	NotificationNumber      string  // Номер извещения
	Address                 string  // Адрес
	Area                    float64 // Площадь, кв.м
	MonthlyRental           int     // Стоимость аренды в месяц, рублей
	RentalPeriod            int     // Срок аренды, месяцев
	WebLink                 string  // Ссылка для просмотра
	AuctionData             string  // Дата проведения аукциона
	ClosingApplicationsDate string  // Дата окончания подачи заявок
	GuaranteeAmount         string  // Информация про залог
}
