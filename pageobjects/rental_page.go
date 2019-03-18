package pageobjects

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tebeka/selenium"
	"time"
)

/* Войти в расширенный поиск */
func ComeInExtSearch(webDriver selenium.WebDriver) {

	extSearchButtonXpath := "//ins[@id='ext_search']"
	btn, err := webDriver.FindElement(selenium.ByXPATH, extSearchButtonXpath)
	if err != nil {
		beego.Error(fmt.Sprintf("Не нашлась кнопка 'Расширенный поиск'"))
		panic(err)
	}

	err = btn.Click()
	if err != nil {
		beego.Error(fmt.Sprintf("Не кликнулась кнопка 'Расширенный поиск'"))
		panic(err)
	}
	time.Sleep(3 * time.Second)
}

/* Указать тип имущества */
func SetAuctionType(webDriver selenium.WebDriver) {

	actionTypeLinkXpath := "//li[text()='В процессе подачи заявок']"
	link, err := webDriver.FindElement(selenium.ByXPATH, actionTypeLinkXpath)
	if err != nil {
		beego.Error(fmt.Sprintf("Не нашлась ссылка 'В процессе подачи заявок'"))
		panic(err)
	}

	err = link.Click()
	if err != nil {
		beego.Error(fmt.Sprintf("Не кликнулась ссылка 'В процессе подачи заявок'"))
		panic(err)
	}
	time.Sleep(10 * time.Second)
}
