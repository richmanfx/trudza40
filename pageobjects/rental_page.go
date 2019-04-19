package pageobjects

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tebeka/selenium"
	"strconv"
	"time"
	"trudza40/models"
)

/* Войти в расширенный поиск */
func ComeInExtSearch(webDriver selenium.WebDriver) {

	// Закрыть алерт
	err := webDriver.AcceptAlert()
	seleniumError(err, "Не закрылся Аллерт")

	extSearchButtonXpath := "//ins[@id='ext_search']"
	btn, err := webDriver.FindElement(selenium.ByXPATH, extSearchButtonXpath)

	msg := "кнопка 'Расширенный поиск'"
	seleniumError(err, "Не нашлась "+msg)

	err = btn.Click()
	seleniumError(err, "Не кликнулась "+msg)

	beego.Info("Кликнута " + msg)
	time.Sleep(2 * time.Second)
}

/* Выбрать тип торгов */
func SetTradesType(webDriver selenium.WebDriver) {

	actionTypeLinkXpath := "//li[text()='В процессе подачи заявок']"
	link, err := webDriver.FindElement(selenium.ByXPATH, actionTypeLinkXpath)

	msg := "ссылка 'В процессе подачи заявок'"
	seleniumError(err, "Не нашлась "+msg)

	err = link.Click()
	seleniumError(err, "Не кликнулась "+msg)

	beego.Info("Кликнута " + msg)
	time.Sleep(2 * time.Second)

}

/* Указать тип имущества */
func SetAuctionType(webDriver selenium.WebDriver, settings *models.Settings) {

	// Кнопка с картинкой книги
	msg := "кнопка с изображением 'Тип имущества'"
	propertyTypeImgXpath :=
		"//td/label[text()='Тип имущества:']/../following-sibling::td[1]//table//tr/td/a[@title='Выбрать']/img"
	imgButton, err := webDriver.FindElement(selenium.ByXPATH, propertyTypeImgXpath)
	seleniumError(err, "Не нашлась "+msg)

	err = imgButton.Click()
	seleniumError(err, "Не кликнулась "+msg)

	beego.Info("Кликнута " + msg)
	time.Sleep(2 * time.Second)

	// Чекбокс
	checkBoxXpath := fmt.Sprintf("//td/span[text()='%s']/preceding-sibling::input", settings.PropertyType)
	checkBoxElement, err := webDriver.FindElement(selenium.ByXPATH, checkBoxXpath)
	seleniumError(err, fmt.Sprintf("Не нашёлся чекбокс '%s'", settings.PropertyType))

	err = checkBoxElement.Click()
	seleniumError(err, fmt.Sprintf("Не кликнулся чекбокс '%s' ", settings.PropertyType))

	beego.Info("Выбран чекбокс " + settings.PropertyType)

	// Кнопка "Выбрать"
	msg = "кнопка 'Выбрать'"
	buttonXpath := "//ins[text()='Выбрать']"
	buttonElement, err := webDriver.FindElement(selenium.ByXPATH, buttonXpath)
	seleniumError(err, "Не нашлась "+msg)

	err = buttonElement.Click()
	seleniumError(err, "Не кликнулась "+msg)

	beego.Info("Кликнута " + msg)
	time.Sleep(2 * time.Second)

}

/* Указать вид договора */
func SetContractType(webDriver selenium.WebDriver, settings *models.Settings) {

	contractTypeImgXpath := "//td/label[text()='Вид договора:']/../" +
		"following-sibling::td[1]//table//tr/td/a[@title='Выбрать']/img"

	imgButton, err := webDriver.FindElement(selenium.ByXPATH, contractTypeImgXpath)

	msg := "кнопка с изображением 'Вид договора'"
	seleniumError(err, "Не нашлась "+msg)

	err = imgButton.Click()
	seleniumError(err, "Не кликнулась "+msg)

	beego.Info("Кликнута " + msg)
	time.Sleep(2 * time.Second)

	// Чекбокс
	checkBoxXpath := fmt.Sprintf("//td/span[text()='%s']/preceding-sibling::input", settings.ContractType)
	checkBoxElement, err := webDriver.FindElement(selenium.ByXPATH, checkBoxXpath)
	seleniumError(err, fmt.Sprintf("Не нашёлся чекбокс '%s'", settings.ContractType))

	err = checkBoxElement.Click()
	seleniumError(err, fmt.Sprintf("Не кликнулся чекбокс '%s' ", settings.ContractType))

	beego.Info("Выбран чекбокс " + settings.ContractType)

	// Кнопка "Выбрать"
	msg = "кнопка 'Выбрать'"

	buttonXpath := "//ins[text()='Выбрать']"
	buttonElement, err := webDriver.FindElement(selenium.ByXPATH, buttonXpath)
	seleniumError(err, "Не нашлась "+msg)

	err = buttonElement.Click()
	seleniumError(err, "Не кликнулась "+msg)

	beego.Info("Кликнута " + msg)
	time.Sleep(2 * time.Second)

}

/* Указать страну */
func SetCountry(webDriver selenium.WebDriver) {

	labelSelectCountryXpath := "//label[text()='Страна размещения:']"
	selectCountryXpath := "//option[@title='РОССИЯ']"

	msg := "селектор выбора страны"
	labelSelectCountryElement, err := webDriver.FindElement(selenium.ByXPATH, labelSelectCountryXpath)
	seleniumError(err, "Не нашёлся "+msg)

	err = labelSelectCountryElement.Click()
	beego.Info("Кликнули " + msg)
	seleniumError(err, "Не кликнулся "+msg)

	selectCountryElement, err := webDriver.FindElement(selenium.ByXPATH, selectCountryXpath)
	msg = "пункт выбора страны"
	seleniumError(err, "Не нашёлся "+msg)

	err = selectCountryElement.Click()
	beego.Info("Кликнули " + msg)
	seleniumError(err, "Не кликнулся "+msg)

	beego.Info("Указана страна")
	time.Sleep(2 * time.Second)

}

/* Указать местоположение имущества */
func SetPropertyLocation(webDriver selenium.WebDriver, settings *models.Settings) {

	locationImgXpath := "//td/label[text()='Местоположение:']/.." +
		"/following-sibling::td[1]//table//tr/td/a[@title='Выбрать']/img"

	msg := "кнопка с изображением 'Местоположение'"
	imgButton, err := webDriver.FindElement(selenium.ByXPATH, locationImgXpath)
	seleniumError(err, "Не нашлась "+msg)

	err = imgButton.Click()
	seleniumError(err, "Не кликнулась "+msg)

	beego.Info("Кликнута " + msg)
	time.Sleep(2 * time.Second)

	// Субъект РФ
	fieldXpath := "//input[@name='container1:level1']"

	msg = "поле 'Субъект РФ'"
	fieldElement, err := webDriver.FindElement(selenium.ByXPATH, fieldXpath)
	seleniumError(err, "Не нашлось "+msg)

	err = fieldElement.SendKeys(settings.PropertyLocation)
	seleniumError(err, "Не введено значение в "+msg)

	beego.Info(fmt.Sprintf("Введено значение '%s' в %s", settings.PropertyLocation, msg))

	// Кнопка "Выбрать"
	buttonXpath := "//ins[text()='Выбрать']"
	buttonElement, err := webDriver.FindElement(selenium.ByXPATH, buttonXpath)

	msg = "кнопка 'Выбрать'"
	seleniumError(err, "Не нашлась "+msg)

	err = buttonElement.Click()
	seleniumError(err, "Не кликнулась "+msg)

	beego.Info("Кликнута " + msg)
	time.Sleep(2 * time.Second)
}

/* Указать диапазон площади объекта */
func SetObjectAreaRange(webDriver selenium.WebDriver, settings *models.Settings) {

	msg := "поле 'Площадь (м²) с'"

	minFieldXpath := "//input[@name='extended:areaMeters:stringAreaMetersFrom']"
	inputValueInField(webDriver, minFieldXpath, settings.MinArea, msg)

	maxFieldXpath := "//input[@name='extended:areaMeters:stringAreaMetersTo']"
	inputValueInField(webDriver, maxFieldXpath, settings.MaxArea, msg)

}

/* Ввод значения в поле ввода */
func inputValueInField(webDriver selenium.WebDriver, fieldXpath string, area uint, msg string) {
	fieldElement, err := webDriver.FindElement(selenium.ByXPATH, fieldXpath)

	seleniumError(err, "Не нашлось "+msg)
	err = fieldElement.SendKeys(strconv.Itoa(int(area)))
	seleniumError(err, "Не введено значение в "+msg)
	beego.Info(fmt.Sprintf("Введено значение '%d' в %s", area, msg))
}

/* Обработка селениумных ошибок */
func seleniumError(err error, msg string) {
	if err != nil {
		beego.Error(fmt.Sprintf("%s: '%s'", msg, err))
		panic(err)
	}
}

/* Указать минимальный срок аренды */
func SetRentalPeriod(webDriver selenium.WebDriver, settings *models.Settings) {

	msg := "поле 'Срок аренды (мес.) с:'"
	fieldXpath := "//input[@name='extended:propertyExtended:stringRentFrom']"
	inputValueInField(webDriver, fieldXpath, settings.MinRentalPeriod*12, msg)

}
