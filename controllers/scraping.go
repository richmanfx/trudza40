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

var settings *models.Settings

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
	settings.AverageRental = uint(averageRental)

	profitMonths, _ := strconv.Atoi(controller.GetString("profit_months"))
	settings.ProfitMonths = uint(profitMonths)

	priorRepair, _ := strconv.Atoi(controller.GetString("prior_repair"))
	settings.PriorRepair = uint(priorRepair)

	contractRegistration, _ := strconv.Atoi(controller.GetString("contract_registration"))
	settings.ContractRegistration = uint(contractRegistration)

	runningCost, _ := strconv.Atoi(controller.GetString("running_cost"))
	settings.RunningCost = uint(runningCost)

	yearlyInsurance, _ := strconv.Atoi(controller.GetString("yearly_insurance"))
	settings.YearlyInsurance = uint(yearlyInsurance)

	monthlyHeating, _ := strconv.Atoi(controller.GetString("monthly_heating"))
	settings.MonthlyHeating = uint(monthlyHeating)

	housingOfficeMaintenance, _ := strconv.Atoi(controller.GetString("housing_office_maintenance"))
	settings.HousingOfficeMaintenance = uint(housingOfficeMaintenance)

	accountingService, _ := strconv.Atoi(controller.GetString("accounting_service"))
	settings.AccountingService = uint(accountingService)

	requiredProfitMargin, _ := strconv.Atoi(controller.GetString("required_profit_margin"))
	settings.RequiredProfitMargin = uint(requiredProfitMargin)

	o := orm.NewOrm() // Использовать ORM "Ormer"
	orm.Debug = true  // Логирование ORM запросов

	num, err := o.QueryTable("settings").Filter("user_id", GlobalUserId).Update(orm.Params{
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
		beego.Info(fmt.Sprintf("Настройки сохранены в БД, записей '%d'", num))
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

	// Собрать информацию по объектам на всех страницах
	ObjectInfoCollect(webDriver)

}

func ObjectInfoCollect(webDriver selenium.WebDriver) {

	// Определить количество найденных объектов
	quantity := pageobjects.GetObjectsQuantity(webDriver)
	beego.Info("Количество найденных объектов: ", quantity)

	// Собрать иформацию об объектах на текущей странице в коллекцию
	objectInfo := onePageObjectInfoCollect(webDriver)
	beego.Info(objectInfo)

	// Добавить к основной коллекци

	// Есть ли следующая страница - рекурсия
}

/* Собрать иформацию об объектах на текущей странице в коллекцию */
func onePageObjectInfoCollect(webDriver selenium.WebDriver) []models.ObjectInfo {

	var objects []models.ObjectInfo

	// Количество объектов на данной странице
	realObjectXpath := "//div[@class='scrollx']/table//tr[contains(@class,'datarow')]"
	realObjects, err := webDriver.FindElements(selenium.ByXPATH, realObjectXpath)
	pageobjects.SeleniumError(err, "Не нашлось количество объектов недвижимости на странице")
	realObjectsQuantity := len(realObjects)
	beego.Info("количество объектов недвижимости на странице:", realObjectsQuantity)

	// Номера извещений объектов
	noticeNumbersXpath := realObjectXpath + "/td[3]/span/span[1]"
	objectsNoticeNumbers, err := webDriver.FindElements(selenium.ByXPATH, noticeNumbersXpath)
	pageobjects.SeleniumError(err, "Не нашлись номера извещений объектов")
	beego.Info(objectsNoticeNumbers)

	// Площадь объектов
	areaXpath := realObjectXpath + "/td[3]/span/span[4]"
	objectsAreas, err := webDriver.FindElements(selenium.ByXPATH, areaXpath)
	pageobjects.SeleniumError(err, "Не нашлись площади объектов")
	beego.Info(objectsAreas)

	// Стоимость аренды в месяц
	rentXpath := realObjectXpath + "/td[7]/span"
	objectsRent, err := webDriver.FindElements(selenium.ByXPATH, rentXpath)
	pageobjects.SeleniumError(err, "Не нашлись стоимости аренды объектов")
	beego.Info(objectsRent)

	// Срок аренды
	rentPeriodsXpath := realObjectXpath + "/td[6]/span/span[2]"
	objectsRentPeriods, err := webDriver.FindElements(selenium.ByXPATH, rentPeriodsXpath)
	pageobjects.SeleniumError(err, "Не нашлись сроки аренды объектов")
	beego.Info(objectsRentPeriods)

	// Ссылка для просмотра
	linkXpath := realObjectXpath + "/td[1]//a[@title='Просмотр']"
	objectsLink, err := webDriver.FindElements(selenium.ByXPATH, linkXpath)
	pageobjects.SeleniumError(err, "Не нашлись ссылки для просмотра объектов")
	beego.Info(objectsLink)

	// Информацию в коллекцию
	for index := 0; index < realObjectsQuantity; index++ {

		objectsNoticeNumber, _ := objectsNoticeNumbers[index].Text()
		beego.Info("Номер извещения", objectsNoticeNumber)

		objectsArea, _ := objectsAreas[index].Text()
		beego.Info("Площадь", strings.Replace(objectsArea, " м²", "", 1))

		objectRent, _ := objectsRent[index].Text()
		str1 := strings.Replace(objectRent, ",", ".", -1)
		str2 := strings.Replace(str1, " ", "", -1)
		rent := strings.Replace(str2, "руб.", "", -1)
		beego.Info("Аренда в месяц", rent)
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
