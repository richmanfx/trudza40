package models

type ObjectInfo struct {
	NotificationNumber string  // Номер извещения
	Area               float64 // Площадь, кв.м
	MonthlyRental      float64 // Стоимость аренды в месяц, рублей
	RentalPeriod       int     // Срок аренды, месяцев TODO: на странице и 'лет', и 'мес.'
	WebLink            string  // Ссылка для просмотра
}
