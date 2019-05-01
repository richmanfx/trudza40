package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tebeka/selenium"
	"strconv"
	"strings"
	"trudza40/models"
	"trudza40/pageobjects"
)

type ScrapController struct {
	beego.Controller
}

/* Глобальные переменные */
var settings *models.Settings
var allObjectsInfo []models.ObjectInfo

//var log = logs.NewLogger(10000)

/* Настроить параметры для скрапинга сайта "torgi.gov.ru" */
func (controller *ScrapController) TorgiGovRuSettings() {
	controller.TplName = "settings-torgi-gov-ru.tpl"
	controller.Data["title"] = "Settings"

	// Получить значения настроек из БД для залогиненного пользователя
	if GlobalUserId == 0 {
		beego.Error("Пользователь не авторизован, UserId = 0")
	}
	settings = models.GetTorgiGovRuSettings(GlobalUserId)

	// Вывести параметры из настроек на форму с настройками
	controller.Data["settings"] = settings
}

/* Сохранить настройки в БД */
func (controller *ScrapController) SaveSettings() {

	//var settings models.Settings

	// Данные из формы
	settings.SettingsName = controller.GetString("settings_name")

	browserWidth, _ := strconv.Atoi(controller.GetString("browser_width"))
	settings.BrowserWidth = uint(browserWidth)

	browserHeight, _ := strconv.Atoi(controller.GetString("browser_height"))
	settings.BrowserHeight = uint(browserHeight)

	settings.HostPageUrl = controller.GetString("host_page_url")

	flashAllowed := controller.GetString("flash_allowed")
	if flashAllowed == "on" {
		settings.FlashAllowed = true
	} else {
		settings.FlashAllowed = false
	}

	flashQuantity, _ := strconv.Atoi(controller.GetString("flash_quantity"))
	settings.FlashQuantity = uint(flashQuantity)

	flashPeriod, _ := strconv.Atoi(controller.GetString("flash_period"))
	settings.FlashPeriod = uint(flashPeriod)

	settings.DebugLevel = controller.GetString("debug_level")

	minArea, _ := strconv.Atoi(controller.GetString("min_area"))
	settings.MinArea = uint(minArea)

	maxArea, _ := strconv.Atoi(controller.GetString("max_area"))
	settings.MaxArea = uint(maxArea)

	minRentalPeriod, _ := strconv.Atoi(controller.GetString("min_rental_period"))
	settings.MinRentalPeriod = uint(minRentalPeriod)

	settings.PropertyType = controller.GetString("property_type")
	settings.ContractType = controller.GetString("contract_type")
	settings.Country = controller.GetString("country")
	settings.PropertyLocation = controller.GetString("property_location")
	settings.SortFieldName = controller.GetString("sort_field_name")

	averageRental, _ := strconv.Atoi(controller.GetString("average_rental"))
	settings.AverageRental = int(averageRental)

	profitMonths, _ := strconv.Atoi(controller.GetString("profit_months"))
	settings.ProfitMonths = int(profitMonths)

	priorRepair, _ := strconv.Atoi(controller.GetString("prior_repair"))
	settings.PriorRepair = int(priorRepair)

	contractRegistration, _ := strconv.Atoi(controller.GetString("contract_registration"))
	settings.ContractRegistration = int(contractRegistration)

	runningCost, _ := strconv.Atoi(controller.GetString("running_cost"))
	settings.RunningCost = int(runningCost)

	yearlyInsurance, _ := strconv.Atoi(controller.GetString("yearly_insurance"))
	settings.YearlyInsurance = int(yearlyInsurance)

	monthlyHeating, _ := strconv.Atoi(controller.GetString("monthly_heating"))
	settings.MonthlyHeating = int(monthlyHeating)

	housingOfficeMaintenance, _ := strconv.Atoi(controller.GetString("housing_office_maintenance"))
	settings.HousingOfficeMaintenance = int(housingOfficeMaintenance)

	accountingService, _ := strconv.Atoi(controller.GetString("accounting_service"))
	settings.AccountingService = int(accountingService)

	requiredProfitMargin, _ := strconv.Atoi(controller.GetString("required_profit_margin"))
	settings.RequiredProfitMargin = int(requiredProfitMargin)

	o := orm.NewOrm() // Использовать ORM "Ormer"
	orm.Debug = false // Логирование ORM запросов

	_, err := o.QueryTable("settings").Filter("user_id", GlobalUserId).Update(orm.Params{
		"user_id":                    GlobalUserId,
		"settings_name":              settings.SettingsName,
		"browser_width":              settings.BrowserWidth,
		"browser_height":             settings.BrowserHeight,
		"host_page_url":              settings.HostPageUrl,
		"flash_allowed":              settings.FlashAllowed,
		"flash_quantity":             settings.FlashQuantity,
		"flash_period":               settings.FlashPeriod,
		"debug_level":                settings.DebugLevel,
		"min_area":                   settings.MinArea,
		"max_area":                   settings.MaxArea,
		"min_rental_period":          settings.MinRentalPeriod,
		"property_type":              settings.PropertyType,
		"contract_type":              settings.ContractType,
		"country":                    settings.Country,
		"property_location":          settings.PropertyLocation,
		"sort_field_name":            settings.SortFieldName,
		"average_rental":             settings.AverageRental,
		"profit_months":              settings.ProfitMonths,
		"prior_repair":               settings.PriorRepair,
		"contract_registration":      settings.ContractRegistration,
		"running_cost":               settings.RunningCost,
		"yearly_insurance":           settings.YearlyInsurance,
		"monthly_heating":            settings.MonthlyHeating,
		"housing_office_maintenance": settings.HousingOfficeMaintenance,
		"accounting_service":         settings.AccountingService,
		"required_profit_margin":     settings.RequiredProfitMargin,
	})
	if err == nil {
		//beego.Debug(fmt.Sprintf("Настройки сохранены в БД, записей '%d'", num))
		controller.TplName = "message-modal.tpl"
		controller.Data["title"] = "Info"
		controller.Data["message1"] = "Информация"
		controller.Data["message2"] = "Настройки сохранены в БД"
	} else {
		beego.Error(fmt.Sprintf("Не удалилось сохранить настройки в БД: %v", err))
		controller.TplName = "message-modal.tpl"
		controller.Data["title"] = "Ошибка"
		controller.Data["message1"] = "Ошибка"
		controller.Data["message2"] = "Не удалилось сохранить настройки в БД"
		controller.Data["message3"] = err
	}
}

/* Скрапить сайт "torgi.gov.ru" */
func (controller *ScrapController) TorgiGovRuScraping() {
	controller.TplName = "scrap-result.tpl"
	controller.Data["title"] = "Scraping"

	// Считать данные для скрапинга
	settings = models.GetTorgiGovRuSettings(GlobalUserId)

	var webDriver selenium.WebDriver
	var capabilities = make(selenium.Capabilities, 1)
	var browser = "chrome"
	var remoteDriverPort = 4444
	var err error

	capabilities["browserName"] = browser

	// Удалённый ВебДрайвер
	webDriver, err = selenium.NewRemote(capabilities, fmt.Sprintf("http://localhost:%d/wd/hub", remoteDriverPort))
	if err != nil {
		beego.Error("Не смогли получить удалённый WebDriver")
		// TODO: Вывести сообщение об ошибке в браузер
		panic(err)
	}
	defer func() {
		if err := webDriver.Quit(); err != nil {
			beego.Error("Не смогли выйти из инстанса браузера (закрыть WebDriver)")
		}
	}()

	// Выставить размеры окна браузера
	err = webDriver.ResizeWindow("", int(settings.BrowserWidth), int(settings.BrowserHeight))
	if err != nil {
		beego.Error("Браузер не смог выставить размер окна")
	}

	// Открыть страницу
	err = webDriver.Get(settings.HostPageUrl)
	if err != nil {
		beego.Error("Не открылся сайт")
		panic(err)
	}

	// Выставить фильтры поиска
	SetSearchFilters(webDriver)

	// Искать
	ObjectsSearch(webDriver)

	// Определить количество найденных объектов
	quantity := pageobjects.GetObjectsQuantity(webDriver)
	//beego.Info(fmt.Sprintf("Количество найденных объектов: %v", quantity))

	// Слайс для информации обо всех объектах со всех страниц
	allObjectsInfo = make([]models.ObjectInfo, 0, quantity)

	// Собрать информацию по объектам на всех страницах
	ObjectInfoCollect(webDriver)
	//beego.Debug(fmt.Sprintf("Информация об объектах на всех страницах: %v", allObjectsInfo))

	// Сохранить данные в файле // TODO: А надо ли?

	// Рассчитать коэффициент окупаемости для каждого объекта
	PaybackCalculation()

	// Вывести результаты расчётов

}

/* Рассчитать коэффициент окупаемости для каждого объекта */
func PaybackCalculation() {

	for _, objectInfo := range allObjectsInfo {

		// Страховка всей площади за год, руб
		// Стоимость годовой страховки (рубли) в Альфа-Страховании, зависит от площади в метрах
		// 	  100: 4000,
		//    300: 6000,
		//    500: 8000,
		//    1000: 12000,
		//    99999999: 20000,
		var yearAllAreaInsurance int
		if objectInfo.Area < 100 {
			yearAllAreaInsurance = int(settings.YearlyInsurance)
		} else {
			beego.Error("Площадь объекта больше чем то, на что расчитана страховка")
			panic("Площадь объекта больше чем то, на что расчитана страховка")
		}
		beego.Info("Страховка всей площади за год:", yearAllAreaInsurance)

		// Стоимость предварительного ремонта всей площади
		allAreaRepair := settings.PriorRepair * int(objectInfo.Area)
		beego.Info("Стоимость предварительного ремонта всей площади:", allAreaRepair)

		// Стоимость отопления в месяц
		monthHeating := settings.MonthlyHeating * int(objectInfo.Area)
		beego.Info("Стоимость отопления в месяц:", monthHeating)

		// Обслуживание ЖЭКом в месяц
		monthHousingOffice := settings.HousingOfficeMaintenance * int(objectInfo.Area)
		beego.Info("Стоимость обслуживания ЖЭКом в месяц:", monthHousingOffice)

		// Доход от аренды в месяц
		monthRentalIncome := settings.AverageRental * int(objectInfo.Area)
		beego.Info("Доход от аренды в месяц:", monthRentalIncome)

		// Расходы в месяц
		monthPayout := int(objectInfo.MonthlyRental) + monthHeating + monthRentalIncome + settings.AccountingService +
			int((settings.ContractRegistration+settings.RunningCost)/objectInfo.RentalPeriod/12)
		beego.Info("Расходы в месяц:", monthPayout)

		// Доход в год с учётом несдаваемых месяцев
		yearRentalIncome := monthRentalIncome * settings.ProfitMonths
		beego.Info("Доход в год с учётом несдаваемых месяцев:", yearRentalIncome)

		// Коэффициент доходности
		profitMargin := (yearRentalIncome - (monthPayout * 12) - yearAllAreaInsurance) /
			(settings.ContractRegistration + settings.RunningCost)
		beego.Info("Коэффициент доходности:", profitMargin)

		// Безубыточность сдачи, руб/кв.м. в месяц
		lossFreeRent := ((monthPayout * 12) / settings.ProfitMonths) / int(objectInfo.Area)
		beego.Info("Безубыточность сдачи, руб/кв.м. в месяц:", lossFreeRent)

		// Собрать большой словарь с параметрами объектов для отчёта

		// Отсортировать большой словарь по коэффициенту доходности

	}

}

/* Собрать информацию по объектам на всех страницах */
func ObjectInfoCollect(webDriver selenium.WebDriver) {

	// Собрать иформацию об объектах на текущей странице
	objectsInfo := onePageObjectInfoCollect(webDriver)
	//beego.Debug(fmt.Sprintf("Информация об объектах на одной странице: %v", objectsInfo))

	// Добавить к основной коллекци
	allObjectsInfo = append(allObjectsInfo, objectsInfo...)

	// Есть ли следующая страница
	nextPageXpath := "//a[@title='Перейти на одну страницу вперед']"
	nextPage, err := webDriver.FindElements(selenium.ByXPATH, nextPageXpath)
	pageobjects.SeleniumError(err, "Ошибка при определении наличия следующей страницы")

	if len(nextPage) < 1 { // Условие выхода из рекурсии
		// Выходим
		return
	} else {
		// Перейти на следующую страницу
		pageobjects.GoToNextPage(webDriver)

		// Рекурсия
		ObjectInfoCollect(webDriver)
	}

}

/* Собрать иформацию об объектах на текущей странице в коллекцию */
func onePageObjectInfoCollect(webDriver selenium.WebDriver) []models.ObjectInfo {

	var objects []models.ObjectInfo

	// Количество объектов на данной странице
	realObjectXpath := "//div[@class='scrollx']/table//tr[contains(@class,'datarow')]"
	realObjects, err := webDriver.FindElements(selenium.ByXPATH, realObjectXpath)
	pageobjects.SeleniumError(err, "Не нашлось количество объектов недвижимости на странице")
	realObjectsQuantity := len(realObjects)
	//beego.Debug("количество объектов недвижимости на странице:", realObjectsQuantity)

	// Номера извещений объектов
	noticeNumbersXpath := realObjectXpath + "/td[3]/span/span[1]"
	objectsNoticeNumbers, err := webDriver.FindElements(selenium.ByXPATH, noticeNumbersXpath)
	pageobjects.SeleniumError(err, "Не нашлись номера извещений объектов")
	//beego.Debug(objectsNoticeNumbers)

	// Площадь объектов
	areaXpath := realObjectXpath + "/td[3]/span/span[4]"
	objectsAreas, err := webDriver.FindElements(selenium.ByXPATH, areaXpath)
	pageobjects.SeleniumError(err, "Не нашлись площади объектов")
	//beego.Debug(objectsAreas)

	// Стоимость аренды в месяц
	rentXpath := realObjectXpath + "/td[7]/span"
	objectsRent, err := webDriver.FindElements(selenium.ByXPATH, rentXpath)
	pageobjects.SeleniumError(err, "Не нашлись стоимости аренды объектов")
	//beego.Debug(objectsRent)

	// Срок аренды
	rentPeriodsXpath := realObjectXpath + "/td[6]/span/span[2]"
	objectsRentPeriods, err := webDriver.FindElements(selenium.ByXPATH, rentPeriodsXpath)
	pageobjects.SeleniumError(err, "Не нашлись сроки аренды объектов")
	//beego.Debug(objectsRentPeriods)

	// Ссылка для просмотра
	linkXpath := realObjectXpath + "/td[1]//a[@title='Просмотр']"
	objectsLinks, err := webDriver.FindElements(selenium.ByXPATH, linkXpath)
	pageobjects.SeleniumError(err, "Не нашлись ссылки для просмотра объектов")
	//beego.Debug(objectsLinks)

	// Информацию в коллекцию
	for index := 0; index < realObjectsQuantity; index++ {

		var object models.ObjectInfo

		// Номер извещения объекта
		object.NotificationNumber, _ = objectsNoticeNumbers[index].Text()
		//beego.Debug(fmt.Sprintf("Номер извещения: %s", object.NotificationNumber))

		// Площадь объекта
		objectsArea, _ := objectsAreas[index].Text()
		//beego.Debug(fmt.Sprintf("Площадь: %s", strings.Replace(objectsArea, " м²", "", 1)))
		object.Area, err = strconv.ParseFloat(strings.Replace(objectsArea, " м²", "", 1), 64)
		if err != nil {
			beego.Error("Ошибка преобразования площади объекта из строки в float: ", err)
		}

		// Стоимость аренды в месяц
		objectRent, _ := objectsRent[index].Text()
		tempString1 := strings.Replace(objectRent, ",", ".", -1)
		tempString2 := strings.Replace(tempString1, " ", "", -1)
		rent := strings.Replace(tempString2, "руб.", "", -1)
		//beego.Debug(fmt.Sprintf("Аренда в месяц: %s", rent))
		object.MonthlyRental, err = strconv.ParseFloat(rent, 64)
		if err != nil {
			beego.Error("Ошибка преобразования стоимости аренды объекта из строки в float: ", err)
		}

		// Срок аренды
		objectRentPeriod, _ := objectsRentPeriods[index].Text()
		//beego.Debug(fmt.Sprintf("Срок аренды: %s", strings.Replace(objectRentPeriod, " лет", "", -1))) // TODO: На странице и "лет", и "мес"???
		object.RentalPeriod, err = strconv.Atoi(strings.Replace(objectRentPeriod, " лет", "", -1))
		if err != nil {
			beego.Error("Ошибка преобразования срока аренды объекта из строки в int: ", err)
		}

		// Ссылка для просмотра
		object.WebLink, _ = objectsLinks[index].GetAttribute("href")
		//beego.Debug(fmt.Sprintf("Ссылка для просмотра: %s", object.WebLink))

		// Добавить в коллекцию информацию про один объект
		objects = append(objects, object)

	}

	return objects
}

/* Искать */
func ObjectsSearch(webDriver selenium.WebDriver) {

	// Искать
	pageobjects.SearchButtonClick(webDriver)

	// Дождаться отображения объектов
	pageobjects.ObjectsWait(webDriver)

}

/* Выставить фильтры поиска */
func SetSearchFilters(webDriver selenium.WebDriver) {

	// Войти в расширенный поиск
	pageobjects.ComeInExtSearch(webDriver)

	// Выбрать тип торгов
	pageobjects.SetTradesType(webDriver)

	// Указать тип имущества
	pageobjects.SetAuctionType(webDriver, settings)

	// Указать вид договора
	pageobjects.SetContractType(webDriver, settings)

	// Указать страну
	pageobjects.SetCountry(webDriver)

	// Указать местоположение имущества
	pageobjects.SetPropertyLocation(webDriver, settings)

	// Указать диапазон площади объекта
	pageobjects.SetObjectAreaRange(webDriver, settings)

	// Указать минимальный срок аренды
	pageobjects.SetRentalPeriod(webDriver, settings)

}
