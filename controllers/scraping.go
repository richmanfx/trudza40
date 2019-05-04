package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tebeka/selenium"
	"math"
	"strconv"
	"strings"
	"time"
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

	//controller.ProgressBarShow()		// TODO: не работает пока

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

	// Собрать дополнительную информацию по каждому объекту
	AdditionalObjectInfoCollect(webDriver)

	// Рассчитать все параметры для каждого объекта
	scrapResult := PaybackCalculation()

	// Подготовить заголовки столбцов
	titles := getTableTitles()

	// Вывести результаты расчётов
	controller.HtmlReportCreate(titles, scrapResult)

}

/* Собрать дополнительную информацию по каждому объекту */
func AdditionalObjectInfoCollect(webDriver selenium.WebDriver) {

	for index, object := range allObjectsInfo {

		// Перейти на страницу объекта недвижимости
		err := webDriver.Get(object.WebLink)
		time.Sleep(3 * time.Second)
		pageobjects.SeleniumError(err, "Не открыласть страница объекта")

		// Сумма залога
		depositXpath := "//label[contains(text(),'Описание обременения')]/../../td/span"
		deposit, err := webDriver.FindElement(selenium.ByXPATH, depositXpath)
		pageobjects.SeleniumError(err, "Не нашлась информация про залог/депозит")
		allObjectsInfo[index].GuaranteeAmount, _ = deposit.Text()
		if allObjectsInfo[index].GuaranteeAmount == "" {
			depositXpath = "//label[contains(text(),'Размер задатка')]/../../td//table//span"
			deposit, err := webDriver.FindElement(selenium.ByXPATH, depositXpath)
			pageobjects.SeleniumError(err, "Не нашлась информация про задаток")
			allObjectsInfo[index].GuaranteeAmount, _ = deposit.Text()
		}

		//// На закладку "Общие"
		tabXpath := "//span[text()='Общие']"
		tabLink, err := webDriver.FindElement(selenium.ByXPATH, tabXpath)
		pageobjects.SeleniumError(err, "Не нашлась закладка 'Общие'")
		err = tabLink.Click()
		pageobjects.SeleniumError(err, "Не кликнулась закладка 'Общие'")
		time.Sleep(3 * time.Second)

		// Адрес
		addressXpath := "//label[contains(text(),'Адрес')]/../../td/span"
		fullAddress, err := webDriver.FindElement(selenium.ByXPATH, addressXpath)
		pageobjects.SeleniumError(err, "Не нашёлся адрес")
		allObjectsInfo[index].Address, _ = fullAddress.Text()

		// Дата торгов
		auctionDateXpath := "//label[contains(text(),'Дата и время проведения аукциона')]/../../td/span"
		auctionDate, err := webDriver.FindElement(selenium.ByXPATH, auctionDateXpath)
		pageobjects.SeleniumError(err, "Не нашласть дата проведения аукциона")
		allObjectsInfo[index].AuctionData, _ = auctionDate.Text()

		// Дата окончания подачи заявок
		closingApplicationsDateXpath := "//label[contains(text(),'Дата окончания подачи заявок')]/../../td/span"
		closingApplicationsDate, err := webDriver.FindElement(selenium.ByXPATH, closingApplicationsDateXpath)
		pageobjects.SeleniumError(err, "Не нашласть дата окончания подачи заявок")
		allObjectsInfo[index].ClosingApplicationsDate, _ = closingApplicationsDate.Text()

	}

}

func (controller *ScrapController) ProgressBarShow() {
	controller.TplName = "message-modal.tpl"
	controller.Data["message2"] = "<progress class=\"progress is-large is-info\" max=\"100\">60%</progress>"
}

/* Подготовить заголовки столбцов таблицы */
func getTableTitles() []string {

	var titles []string

	titles = append(titles, "N")
	titles = append(titles, "Номер извещения")
	titles = append(titles, "Коэффициент доходности")
	titles = append(titles, "Адрес")
	titles = append(titles, "Площадь, кв.м")
	titles = append(titles, "Дата торгов")
	titles = append(titles, "Дата окончания подачи заявок")
	titles = append(titles, "Сумма залога")
	titles = append(titles, "Безубыточная сдача, руб/кв.м. в месяц")

	titles = append(titles, "Доход в месяц, рублей")
	titles = append(titles, "Расходы в месяц, рублей")
	titles = append(titles, "Выплата ренты в месяц, рублей")
	titles = append(titles, "Стоимость отопления в месяц, рублей")
	titles = append(titles, "Обслуживание ЖЭКом в месяц, рублей")

	titles = append(titles, "Доход в год, рублей")
	titles = append(titles, "Выплата ренты в год, рублей")
	titles = append(titles, "Страховка за год, рублей")
	titles = append(titles, "Предварительный ремонт, рублей")

	return titles
}

/* Вывести результаты расчётов */
func (controller *ScrapController) HtmlReportCreate(titles []string, scrapResult []models.ObjectScrapResult) {
	controller.TplName = "result-torgi-gov-ru.tpl"
	controller.Data["titles"] = titles
	controller.Data["result"] = scrapResult
	//beego.Info("scrapResult:", scrapResult)
	controller.Data["settings"] = settings
}

/* Рассчитать коэффициент окупаемости для каждого объекта */
func PaybackCalculation() []models.ObjectScrapResult {

	var scrapResult []models.ObjectScrapResult

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
		//beego.Info("\n================================================================================================")
		//beego.Info("Страховка всей площади за год:", yearAllAreaInsurance)

		// Стоимость предварительного ремонта всей площади
		allAreaRepair := settings.PriorRepair * int(objectInfo.Area)
		//beego.Info("Стоимость предварительного ремонта всей площади:", allAreaRepair)

		// Стоимость отопления в месяц
		monthHeating := settings.MonthlyHeating * int(objectInfo.Area)
		//beego.Info("Стоимость отопления в месяц:", monthHeating)

		// Обслуживание ЖЭКом в месяц
		monthHousingOffice := settings.HousingOfficeMaintenance * int(objectInfo.Area)
		//beego.Info("Стоимость обслуживания ЖЭКом в месяц:", monthHousingOffice)

		// Доход от аренды в месяц
		monthRentalIncome := settings.AverageRental * int(objectInfo.Area)
		//beego.Info("Доход от аренды в месяц:", monthRentalIncome)

		// Расходы в месяц
		monthPayout :=
			objectInfo.MonthlyRental + // Стоимость аренды в месяц
				monthHeating + // Стоимость отопления в месяц
				monthHousingOffice + // Обслуживание ЖЭКом в месяц
				settings.AccountingService + // Бухгалтерское обслуживание в месяц
				int((settings.ContractRegistration+settings.RunningCost)/objectInfo.RentalPeriod/12)
		//beego.Info("Расходы в месяц:", monthPayout)

		// Доход в год с учётом несдаваемых месяцев
		yearRentalIncome := monthRentalIncome * settings.ProfitMonths
		//beego.Info("Доход в год с учётом несдаваемых месяцев:", yearRentalIncome)

		// Коэффициент доходности
		profitMargin := (yearRentalIncome - (monthPayout*12 + yearAllAreaInsurance)) /
			(settings.ContractRegistration + settings.RunningCost)
		//beego.Info("Коэффициент доходности:", profitMargin)

		// Безубыточность сдачи, руб/кв.м. в месяц
		lossFreeRent := ((monthPayout * 12) / settings.ProfitMonths) / int(objectInfo.Area)
		//beego.Info("Безубыточность сдачи, руб/кв.м. в месяц:", lossFreeRent)

		//// Собрать большой словарь с параметрами объектов для отчёта
		var oneObjectScrapResult models.ObjectScrapResult

		// Номер извещения
		oneObjectScrapResult.NotificationNumber = objectInfo.NotificationNumber

		// Ссылка на объект на сайте
		oneObjectScrapResult.WebLink = objectInfo.WebLink

		// Коэффициент доходности
		oneObjectScrapResult.ProfitMargin = profitMargin

		// Адрес объекта
		oneObjectScrapResult.Address = objectInfo.Address

		// Площадь объекта
		oneObjectScrapResult.Area = objectInfo.Area

		// Дата торгов
		oneObjectScrapResult.AuctionData = objectInfo.AuctionData

		// Дата окончания подачи заявок
		oneObjectScrapResult.ClosingApplicationsDate = objectInfo.ClosingApplicationsDate

		// Сумма залога
		oneObjectScrapResult.GuaranteeAmount = objectInfo.GuaranteeAmount

		// Безубыточная сдача
		oneObjectScrapResult.LossFreeRental = lossFreeRent

		// Выплаты ренты в год
		oneObjectScrapResult.YearRental = objectInfo.MonthlyRental * 12

		// Выплаты ренты в месяц
		oneObjectScrapResult.MonthlyRental = objectInfo.MonthlyRental

		// Страховка за год
		oneObjectScrapResult.YearInsurance = yearAllAreaInsurance

		// Расходы в месяц
		oneObjectScrapResult.MonthlyCost = monthPayout

		// Стоимость отопления в месяц
		oneObjectScrapResult.MonthlyHeating = monthHeating

		// Обслуживание ЖЭКом в месяц
		oneObjectScrapResult.HousingOfficeMaintenance = monthHousingOffice

		// Доход в месяц
		oneObjectScrapResult.MonthlyProfit = monthRentalIncome

		// Доход в год
		oneObjectScrapResult.YearProfit = yearRentalIncome

		// Предварительный ремонт
		oneObjectScrapResult.PriorRepair = allAreaRepair

		// Добавить объект в общий результат
		scrapResult = append(scrapResult, oneObjectScrapResult)

	}

	// TODO:
	// Отсортировать большой словарь по значению, указанному в конфиге

	// Проставить порядковый номер в первом столбце
	scrapResult = SetObjectSerialNumber(scrapResult) // TODO: сделать с указателем

	return scrapResult
}

/* Проставить порядковый номер объекта в первом столбце таблицы */
func SetObjectSerialNumber(scrapResult []models.ObjectScrapResult) []models.ObjectScrapResult {
	for index := range scrapResult {
		scrapResult[index].OrderNumber = index + 1
	}
	return scrapResult
}

/* Собрать информацию по объектам на всех страницах */
func ObjectInfoCollect(webDriver selenium.WebDriver) {

	// Собрать иформацию об объектах на текущей странице
	objectsInfo := onePageObjectInfoCollect(webDriver)
	//beego.Debug(fmt.Sprintf("Информация об объектах на одной странице: %v", objectsInfo))

	// Добавить к основной коллекци
	allObjectsInfo = append(allObjectsInfo, objectsInfo...)

	// **********************************************************
	//return // Для отладки - скрапить только первую страницу
	// **********************************************************

	// Есть ли следующая страница
	nextPageXpath := "//a[@title='Перейти на одну страницу вперед']"
	nextPage, err := webDriver.FindElements(selenium.ByXPATH, nextPageXpath)
	pageobjects.SeleniumError(err, "Ошибка при определении наличия следующей страницы")

	if len(nextPage) < 1 { // Условие выхода из рекурсии - нет следующей страницы
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

	//// Запомнить ссылку
	//objectsListStartUrl, err := webDriver.CurrentURL()
	//pageobjects.SeleniumError(err, "Не удалось получить текущий URL")
	//beego.Info("Запомнили ссылку на страницу со списком объектов: ", objectsListStartUrl)

	// Количество объектов на данной странице
	realObjectXpath := "//div[@class='scrollx']/table//tr[contains(@class,'datarow')]"
	realObjects, err := webDriver.FindElements(selenium.ByXPATH, realObjectXpath)
	pageobjects.SeleniumError(err, "Не нашлось количество объектов недвижимости на странице")
	realObjectsQuantity := len(realObjects)
	//beego.Debug("количество объектов недвижимости на странице:", realObjectsQuantity)

	// Объекты на странице
	objects := make([]models.ObjectInfo, realObjectsQuantity)

	// Номер извещения
	noticeNumbersXpath := realObjectXpath + "/td[3]/span/span[1]"
	objectsNoticeNumbers, err := webDriver.FindElements(selenium.ByXPATH, noticeNumbersXpath)
	pageobjects.SeleniumError(err, "Не нашлись номера извещений объектов")
	//beego.Debug(objectsNoticeNumbers)
	for index := range objects {
		objects[index].NotificationNumber, err = objectsNoticeNumbers[index].Text()
		if err != nil {
			beego.Error("Ошибка при записи номеров извещений: ", err)
		}
	}

	// Площадь
	areaXpath := realObjectXpath + "/td[3]/span/span[4]"
	objectsAreas, err := webDriver.FindElements(selenium.ByXPATH, areaXpath)
	pageobjects.SeleniumError(err, "Не нашлись площади объектов")
	//beego.Debug(objectsAreas)
	for index := range objects {
		areaString, _ := objectsAreas[index].Text()
		objects[index].Area, err = strconv.ParseFloat(strings.Replace(areaString, " м²", "", 1), 64)
		if err != nil {
			beego.Error("Ошибка преобразования площади объекта из строки в float: ", err)
		}
	}

	// Стоимость аренды в месяц
	rentXpath := realObjectXpath + "/td[7]/span"
	objectsRent, err := webDriver.FindElements(selenium.ByXPATH, rentXpath)
	pageobjects.SeleniumError(err, "Не нашлись стоимости аренды объектов")
	//beego.Debug(objectsRent)
	for index := range objects {
		tempString1, err := objectsRent[index].Text()
		if err != nil {
			beego.Error("Ошибка преобразования стоимости аренды объекта в строку: ", err)
		}
		tempString2 := strings.Replace(tempString1, ",", ".", -1)
		tempString3 := strings.Replace(tempString2, " ", "", -1)
		rent := strings.Replace(tempString3, "руб.", "", -1)
		tmpFloat, err := strconv.ParseFloat(rent, 64)
		if err != nil {
			beego.Error("Ошибка преобразования стоимости аренды объекта из строки в float: ", err)
		}
		objects[index].MonthlyRental = int(math.Ceil(tmpFloat)) // Ближайшее большее целое
	}

	// Срок аренды
	// TODO: На странице и "лет", и "мес"???
	rentPeriodsXpath := realObjectXpath + "/td[6]/span/span[2]"
	objectsRentPeriods, err := webDriver.FindElements(selenium.ByXPATH, rentPeriodsXpath)
	pageobjects.SeleniumError(err, "Не нашлись сроки аренды объектов")
	//beego.Debug(objectsRentPeriods)
	for index := range objects {
		objectRentPeriodString, err := objectsRentPeriods[index].Text()
		if err != nil {
			beego.Error("Ошибка преобразования сроков аренды объекта в строку: ", err)
		}
		objects[index].RentalPeriod, err = strconv.Atoi(strings.Replace(objectRentPeriodString, " лет", "", -1))
		if err != nil {
			beego.Error("Ошибка преобразования сроков аренды объекта из строки в int: ", err)
		}
	}

	// Ссылка для просмотра - должна быть последней при скрапинге, так как уходит с основной страницы с объектами
	linkXpath := realObjectXpath + "/td[1]//a[@title='Просмотр']"
	objectsLinks, err := webDriver.FindElements(selenium.ByXPATH, linkXpath)
	pageobjects.SeleniumError(err, "Не нашлись ссылки для просмотра объектов")
	//beego.Debug("Ссылки на объекты: ", objectsLinks)
	for index := range objects {
		objects[index].WebLink, err = objectsLinks[index].GetAttribute("href")
		if err != nil {
			beego.Error("Ошибка при записи ссылки для просмотра объектов: ", err)
		}
		//beego.Debug("Ссылка на объект: ", objects[index].WebLink)
	}

	//// Вернуться на страницу со списком объектов
	//beego.Info("Возврат на страницу со списком объектов: ", objectsListStartUrl)
	//err = webDriver.Get(objectsListStartUrl)
	//pageobjects.SeleniumError(err, "Не удалось вернуться на страницу со списком объектов")
	//time.Sleep(3 * time.Second)

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
