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
	if err != nil {
		beego.Error("Не закрылся Аллерт")
		panic(err)
	}

	extSearchButtonXpath := "//ins[@id='ext_search']"
	btn, err := webDriver.FindElement(selenium.ByXPATH, extSearchButtonXpath)

	msg := "кнопка 'Расширенный поиск'"
	if err != nil {
		beego.Error("Не нашлась " + msg)
		panic(err)
	}

	err = btn.Click()
	if err != nil {
		beego.Error("Не кликнулась " + msg)
		panic(err)
	}

	beego.Info("Кликнута " + msg)
	time.Sleep(2 * time.Second)
}

/* Выбрать тип торгов */
func SetTradesType(webDriver selenium.WebDriver) {

	actionTypeLinkXpath := "//li[text()='В процессе подачи заявок']"
	link, err := webDriver.FindElement(selenium.ByXPATH, actionTypeLinkXpath)

	msg := "ссылка 'В процессе подачи заявок'"
	if err != nil {
		beego.Error("Не нашлась " + msg)
		panic(err)
	}

	err = link.Click()
	if err != nil {
		beego.Error("Не кликнулась " + msg)
		panic(err)
	}

	beego.Info("Кликнута " + msg)
	time.Sleep(2 * time.Second)

}

/* Указать тип имущества */
func SetAuctionType(webDriver selenium.WebDriver, settings *models.Settings) {

	// Кнопка с картинкой книги
	propertyTypeImgXpath :=
		"//td/label[text()='Тип имущества:']/../following-sibling::td[1]//table//tr/td/a[@title='Выбрать']/img"
	imgButton, err := webDriver.FindElement(selenium.ByXPATH, propertyTypeImgXpath)

	msg := "кнопка с изображением 'Тип имущества'"
	if err != nil {
		beego.Error("Не нашлась " + msg)
		panic(err)
	}

	err = imgButton.Click()
	if err != nil {
		beego.Error("Не кликнулась " + msg)
		panic(err)
	}

	beego.Info("Кликнута " + msg)
	time.Sleep(2 * time.Second)

	// Чекбокс
	checkBoxXpath := fmt.Sprintf("//td/span[text()='%s']/preceding-sibling::input", settings.PropertyType)
	checkBoxElement, err := webDriver.FindElement(selenium.ByXPATH, checkBoxXpath)

	if err != nil {
		beego.Error(fmt.Sprintf("Не нашёлся чекбокс '%s'", settings.PropertyType))
		panic(err)
	}

	err = checkBoxElement.Click()

	if err != nil {
		beego.Error(fmt.Sprintf("Не кликнулся чекбокс '%s' ", settings.PropertyType))
		panic(err)
	}

	beego.Info("Выбран чекбокс " + settings.PropertyType)
	time.Sleep(2 * time.Second)

	// Кнопка "Выбрать"
	buttonXpath := "//ins[text()='Выбрать']"
	buttonElement, err := webDriver.FindElement(selenium.ByXPATH, buttonXpath)

	msg = "кнопка 'Выбрать'"
	if err != nil {
		beego.Error("Не нашлась " + msg)
		panic(err)
	}

	err = buttonElement.Click()
	if err != nil {
		beego.Error("Не кликнулась " + msg)
		panic(err)
	}

	beego.Info("Кликнута " + msg)
	time.Sleep(2 * time.Second)

}

/* Указать вид договора */
func SetContractType(webDriver selenium.WebDriver, settings *models.Settings) {

	contractTypeImgXpath := "//td/label[text()='Вид договора:']/../" +
		"following-sibling::td[1]//table//tr/td/a[@title='Выбрать']/img"

	imgButton, err := webDriver.FindElement(selenium.ByXPATH, contractTypeImgXpath)

	msg := "кнопка с изображением 'Вид договора'"
	if err != nil {
		beego.Error("Не нашлась " + msg)
		panic(err)
	}

	err = imgButton.Click()
	if err != nil {
		beego.Error("Не кликнулась " + msg)
		panic(err)
	}

	beego.Info("Кликнута " + msg)
	time.Sleep(2 * time.Second)

	// Чекбокс
	checkBoxXpath := fmt.Sprintf("//td/span[text()='%s']/preceding-sibling::input", settings.ContractType)
	checkBoxElement, err := webDriver.FindElement(selenium.ByXPATH, checkBoxXpath)

	if err != nil {
		beego.Error(fmt.Sprintf("Не нашёлся чекбокс '%s'", settings.ContractType))
		panic(err)
	}

	err = checkBoxElement.Click()

	if err != nil {
		beego.Error(fmt.Sprintf("Не кликнулся чекбокс '%s' ", settings.ContractType))
		panic(err)
	}

	beego.Info("Выбран чекбокс " + settings.ContractType)
	time.Sleep(2 * time.Second)

	// Кнопка "Выбрать"
	buttonXpath := "//ins[text()='Выбрать']"
	buttonElement, err := webDriver.FindElement(selenium.ByXPATH, buttonXpath)

	msg = "кнопка 'Выбрать'"
	if err != nil {
		beego.Error("Не нашлась " + msg)
		panic(err)
	}

	err = buttonElement.Click()
	if err != nil {
		beego.Error("Не кликнулась " + msg)
		panic(err)
	}

	beego.Info("Кликнута " + msg)
	time.Sleep(2 * time.Second)

}

// Указать страну
func SetCountry(webDriver selenium.WebDriver) {

	labelSelectCountryXpath := "//label[text()='Страна размещения:']"
	selectCountryXpath := "//option[@title='РОССИЯ']"

	labelSelectCountryElement, err := webDriver.FindElement(selenium.ByXPATH, labelSelectCountryXpath)
	msg := "селектор выбора страны"
	if err != nil {
		beego.Error("Не нашёлся " + msg)
		panic(err)
	}

	err = labelSelectCountryElement.Click()
	beego.Info("Кликнули " + msg)
	if err != nil {
		beego.Error("Не кликнулся " + msg)
		panic(err)
	}

	selectCountryElement, err := webDriver.FindElement(selenium.ByXPATH, selectCountryXpath)
	msg = "пункт выбора страны"
	if err != nil {
		beego.Error("Не нашёлся " + msg)
		panic(err)
	}

	err = selectCountryElement.Click()
	beego.Info("Кликнули " + msg)
	if err != nil {
		beego.Error("Не кликнулся " + msg)
		panic(err)
	}

	beego.Info("Указана страна")
	time.Sleep(2 * time.Second)

}

// Указать местоположение имущества
func SetPropertyLocation(webDriver selenium.WebDriver, settings *models.Settings) {

	locationImgXpath := "//td/label[text()='Местоположение:']/.." +
		"/following-sibling::td[1]//table//tr/td/a[@title='Выбрать']/img"

	imgButton, err := webDriver.FindElement(selenium.ByXPATH, locationImgXpath)
	msg := "кнопка с изображением 'Местоположение'"
	if err != nil {
		beego.Error("Не нашлась " + msg)
		panic(err)
	}

	err = imgButton.Click()
	if err != nil {
		beego.Error("Не кликнулась " + msg)
		panic(err)
	}

	beego.Info("Кликнута " + msg)
	time.Sleep(2 * time.Second)

	// Субъект РФ
	fieldXpath := "//input[@name='container1:level1']"

	fieldElement, err := webDriver.FindElement(selenium.ByXPATH, fieldXpath)
	msg = "поле 'Субъект РФ'"
	if err != nil {
		beego.Error("Не нашлось " + msg)
		panic(err)
	}

	err = fieldElement.SendKeys(settings.PropertyLocation)
	if err != nil {
		beego.Error("Не введено значение в " + msg)
		panic(err)
	}

	beego.Info(fmt.Sprintf("Введено значение '%s' в %s", settings.PropertyLocation, msg))
	time.Sleep(2 * time.Second)

	// Кнопка "Выбрать"
	buttonXpath := "//ins[text()='Выбрать']"
	buttonElement, err := webDriver.FindElement(selenium.ByXPATH, buttonXpath)

	msg = "кнопка 'Выбрать'"
	if err != nil {
		beego.Error("Не нашлась " + msg)
		panic(err)
	}

	err = buttonElement.Click()
	if err != nil {
		beego.Error("Не кликнулась " + msg)
		panic(err)
	}

	beego.Info("Кликнута " + msg)
	time.Sleep(2 * time.Second)
}

// Указать диапазон площади объекта		TODO: Два раза вызывать метод ввода в поле !!!!
func SetObjectAreaRange(webDriver selenium.WebDriver, settings *models.Settings) {

	minFieldXpath := "//input[@name='extended:areaMeters:stringAreaMetersFrom']"
	minField, err := webDriver.FindElement(selenium.ByXPATH, minFieldXpath)
	msg := "поле 'Площадь (м²) с'"
	if err != nil {
		beego.Error("Не нашлось " + msg)
		panic(err)
	}

	err = minField.SendKeys(strconv.Itoa(int(settings.MinArea)))
	if err != nil {
		beego.Error("Не введено значение в " + msg)
		panic(err)
	}

	beego.Info(fmt.Sprintf("Введено значение '%d' в %s", settings.MinArea, msg))
	time.Sleep(2 * time.Second)

}
