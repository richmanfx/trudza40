package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tebeka/selenium"
)

type ScrapController struct {
	beego.Controller
}

/* Настроить параметры для скрапинга сайта "torgi.gov.ru" */
func (context *ScrapController) TorgiGovRuSettings() {
	context.TplName = "settings-torgi-gov-ru.tpl"
	context.Data["title"] = "Settings"

}

/* Скрапить сайт "torgi.gov.ru" */
func (context *ScrapController) TorgiGovRuScraping() {
	context.TplName = "scrap-result.tpl"
	context.Data["title"] = "Scraping"

	// Считать данные для скрапинга
	// TODO: Считываться будет из сохранённых пользователем данных в БД
	// TODO: Временно захардкожено

	var webDriver selenium.WebDriver
	var capabilities = make(selenium.Capabilities, 1)
	var browser = "chrome"
	var remoteDriverPort = 4444
	var err error

	capabilities["browserName"] = browser
	//capabilities["phantomjs.binary.path"] = "/usr/local/bin/phantomjs"

	webDriver, err = selenium.NewRemote(capabilities, fmt.Sprintf("http://localhost:%d/wd/hub", remoteDriverPort))
	if err != nil {
		panic(err)
	}
	defer webDriver.Quit()

	err = webDriver.ResizeWindow("", 1920, 1080)
	if err != nil {
		beego.Error("Браузер не смог выставить размер окна")
	}

	err = webDriver.Get("https://torgi.gov.ru/index.html")
	if err != nil {
		panic(err)
	}

	title, _ := webDriver.Title()

	beego.Info(fmt.Sprintf("Title: %s", title))

}
