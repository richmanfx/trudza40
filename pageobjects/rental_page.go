package pageobjects

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tebeka/selenium"
	"strconv"
	"strings"
	"time"
	"trudza40/models"
)

/* Войти в расширенный поиск */
func ComeInExtSearch(webDriver selenium.WebDriver) {

	//// Закрыть алерт
	//err := webDriver.AcceptAlert()
	//SeleniumError(err, "Не закрылся Аллерт")

	extSearchButtonXpath := "//ins[@id='ext_search']"
	btn, err := webDriver.FindElement(selenium.ByXPATH, extSearchButtonXpath)

	msg := "кнопка 'Расширенный поиск'"
	SeleniumError(err, "Не нашлась "+msg)

	err = btn.Click()
	SeleniumError(err, "Не кликнулась "+msg)

	//beego.Debug("Кликнута " + msg)
	time.Sleep(2 * time.Second)
}

/* Выбрать тип торгов */
func SetTradesType(webDriver selenium.WebDriver) {

	actionTypeLinkXpath := "//li[text()='В процессе подачи заявок']"
	link, err := webDriver.FindElement(selenium.ByXPATH, actionTypeLinkXpath)

	msg := "ссылка 'В процессе подачи заявок'"
	SeleniumError(err, "Не нашлась "+msg)

	err = link.Click()
	SeleniumError(err, "Не кликнулась "+msg)

	//beego.Debug("Кликнута " + msg)
	time.Sleep(2 * time.Second)

}

/* Указать тип имущества */
func SetAuctionType(webDriver selenium.WebDriver, settings *models.Settings) {

	// Кнопка с картинкой книги
	msg := "кнопка с изображением 'Тип имущества'"
	propertyTypeImgXpath :=
		"//td/label[text()='Тип имущества:']/../following-sibling::td[1]//table//tr/td/a[@title='Выбрать']/img"
	imgButton, err := webDriver.FindElement(selenium.ByXPATH, propertyTypeImgXpath)
	SeleniumError(err, "Не нашлась "+msg)

	err = imgButton.Click()
	SeleniumError(err, "Не кликнулась "+msg)

	//beego.Debug("Кликнута " + msg)
	time.Sleep(2 * time.Second)

	// Чекбокс
	checkBoxXpath := fmt.Sprintf("//td/span[text()='%s']/preceding-sibling::input", settings.PropertyType)
	checkBoxElement, err := webDriver.FindElement(selenium.ByXPATH, checkBoxXpath)
	SeleniumError(err, fmt.Sprintf("Не нашёлся чекбокс '%s'", settings.PropertyType))

	err = checkBoxElement.Click()
	SeleniumError(err, fmt.Sprintf("Не кликнулся чекбокс '%s' ", settings.PropertyType))

	//beego.Debug("Выбран чекбокс " + settings.PropertyType)

	// Кнопка "Выбрать"
	msg = "кнопка 'Выбрать'"
	buttonXpath := "//ins[text()='Выбрать']"
	buttonElement, err := webDriver.FindElement(selenium.ByXPATH, buttonXpath)
	SeleniumError(err, "Не нашлась "+msg)

	err = buttonElement.Click()
	SeleniumError(err, "Не кликнулась "+msg)

	//beego.Debug("Кликнута " + msg)
	time.Sleep(2 * time.Second)

}

/* Указать вид договора */
func SetContractType(webDriver selenium.WebDriver, settings *models.Settings) {

	contractTypeImgXpath := "//td/label[text()='Вид договора:']/../" +
		"following-sibling::td[1]//table//tr/td/a[@title='Выбрать']/img"

	imgButton, err := webDriver.FindElement(selenium.ByXPATH, contractTypeImgXpath)

	msg := "кнопка с изображением 'Вид договора'"
	SeleniumError(err, "Не нашлась "+msg)

	err = imgButton.Click()
	SeleniumError(err, "Не кликнулась "+msg)

	//beego.Debug("Кликнута " + msg)
	time.Sleep(2 * time.Second)

	// Чекбокс
	checkBoxXpath := fmt.Sprintf("//td/span[text()='%s']/preceding-sibling::input", settings.ContractType)
	checkBoxElement, err := webDriver.FindElement(selenium.ByXPATH, checkBoxXpath)
	SeleniumError(err, fmt.Sprintf("Не нашёлся чекбокс '%s'", settings.ContractType))

	err = checkBoxElement.Click()
	SeleniumError(err, fmt.Sprintf("Не кликнулся чекбокс '%s' ", settings.ContractType))

	//beego.Debug("Выбран чекбокс " + settings.ContractType)

	// Кнопка "Выбрать"
	msg = "кнопка 'Выбрать'"

	buttonXpath := "//ins[text()='Выбрать']"
	buttonElement, err := webDriver.FindElement(selenium.ByXPATH, buttonXpath)
	SeleniumError(err, "Не нашлась "+msg)

	err = buttonElement.Click()
	SeleniumError(err, "Не кликнулась "+msg)

	//beego.Debug("Кликнута " + msg)
	time.Sleep(2 * time.Second)

}

/* Указать страну */
func SetCountry(webDriver selenium.WebDriver) {

	labelSelectCountryXpath := "//label[text()='Страна размещения:']"
	selectCountryXpath := "//option[@title='РОССИЯ']"

	msg := "селектор выбора страны"
	labelSelectCountryElement, err := webDriver.FindElement(selenium.ByXPATH, labelSelectCountryXpath)
	SeleniumError(err, "Не нашёлся "+msg)

	err = labelSelectCountryElement.Click()
	//beego.Debug("Кликнули " + msg)
	SeleniumError(err, "Не кликнулся "+msg)

	selectCountryElement, err := webDriver.FindElement(selenium.ByXPATH, selectCountryXpath)
	msg = "пункт выбора страны"
	SeleniumError(err, "Не нашёлся "+msg)

	err = selectCountryElement.Click()
	//beego.Debug("Кликнули " + msg)
	SeleniumError(err, "Не кликнулся "+msg)

	//beego.Debug("Указана страна")
	time.Sleep(2 * time.Second)

}

/* Указать местоположение имущества */
func SetPropertyLocation(webDriver selenium.WebDriver, settings *models.Settings) {

	locationImgXpath := "//td/label[text()='Местоположение:']/.." +
		"/following-sibling::td[1]//table//tr/td/a[@title='Выбрать']/img"

	msg := "кнопка с изображением 'Местоположение'"
	imgButton, err := webDriver.FindElement(selenium.ByXPATH, locationImgXpath)
	SeleniumError(err, "Не нашлась "+msg)

	err = imgButton.Click()
	SeleniumError(err, "Не кликнулась "+msg)

	//beego.Debug("Кликнута " + msg)
	time.Sleep(2 * time.Second)

	// Субъект РФ
	fieldXpath := "//input[@name='container1:level1']"

	msg = "поле 'Субъект РФ'"
	fieldElement, err := webDriver.FindElement(selenium.ByXPATH, fieldXpath)
	SeleniumError(err, "Не нашлось "+msg)

	err = fieldElement.SendKeys(settings.PropertyLocation)
	SeleniumError(err, "Не введено значение в "+msg)

	//beego.Debug(fmt.Sprintf("Введено значение '%s' в %s", settings.PropertyLocation, msg))

	// Кнопка "Выбрать"
	buttonXpath := "//ins[text()='Выбрать']"
	buttonElement, err := webDriver.FindElement(selenium.ByXPATH, buttonXpath)

	msg = "кнопка 'Выбрать'"
	SeleniumError(err, "Не нашлась "+msg)

	err = buttonElement.Click()
	SeleniumError(err, "Не кликнулась "+msg)

	//beego.Debug("Кликнута " + msg)
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

	SeleniumError(err, "Не нашлось "+msg)
	err = fieldElement.SendKeys(strconv.Itoa(int(area)))
	SeleniumError(err, "Не введено значение в "+msg)
	//beego.Debug(fmt.Sprintf("Введено значение '%d' в %s", area, msg))
}

/* Обработка селениумных ошибок */
func SeleniumError(err error, msg string) {
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

/* Кликнуть на кнопке поиска */
func SearchButtonClick(webDriver selenium.WebDriver) {

	msg := "кнопка 'Поиск'"
	buttonXpath := "//ins[@id='lot_search']"
	buttonElement, err := webDriver.FindElement(selenium.ByXPATH, buttonXpath)
	SeleniumError(err, "Не нашлась "+msg)

	err = buttonElement.Click()
	SeleniumError(err, "Не нашлась "+msg)

	//beego.Debug("Кликнута " + msg)
	time.Sleep(2 * time.Second)
}

/* Дождаться отображения объектов */
func ObjectsWait(webDriver selenium.WebDriver) {

	checkXpath := "//h2/span[contains(text(),'найдено лотов')]"
	labelElement, err := webDriver.FindElement(selenium.ByXPATH, checkXpath)
	SeleniumError(err, "Поиск не отработал, лоты не нашлись")
	time.Sleep(4 * time.Second)
	labelIsDisplayed, err := labelElement.IsDisplayed()
	SeleniumError(err, "Ошибка при проверке отображения лотов")

	if !labelIsDisplayed {
		beego.Error("Лоты не отобразились")
	} else {
		beego.Debug("Лоты отобразились - удачный поиск")
		//_, err = labelElement.LocationInView()		// Вроде бы так и не скролит
		err = labelElement.MoveTo(0, 0)
	}

}

/* Определить количество найденных объектов */
func GetObjectsQuantity(webDriver selenium.WebDriver) int {

	xpath := "//h2/span[contains(text(),'найдено лотов')]"
	labelElement, err := webDriver.FindElement(selenium.ByXPATH, xpath)
	SeleniumError(err, "Количество лотов не нашлось")
	labelText, err := labelElement.Text()
	splitLabelText := strings.Split(labelText, " ")
	objectsQuantity, _ := strconv.Atoi(splitLabelText[len(splitLabelText)-1])

	return objectsQuantity
}

/* Перейти на следующую страницу - пагинация */
func GoToNextPage(webDriver selenium.WebDriver) {

	nextPageXpath := "//a[@title='Перейти на одну страницу вперед']"
	nextPageLink, err := webDriver.FindElement(selenium.ByXPATH, nextPageXpath)
	SeleniumError(err, "Не обнаружен элемент для пагинации")

	err = nextPageLink.Click()
	SeleniumError(err, "Ошибка при клике элемента пагинации")

	time.Sleep(2 * time.Second)

}
