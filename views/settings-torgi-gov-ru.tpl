{{ template "realty_header.tpl" .}}

<div class="columns">
    <div class="column"></div>
    <div class="column ">
        <div class="content">
            <div class="field boxes">
            <form method="post" class="" action="/realty/save-settings">

                <label class="label">Название комплекта настроек</label>
                <div class="control">
                    <input class="input" id="settings_name" type="text" placeholder="Текст"
                           value="{{ .settings.SettingsName }}" name="settings_name">
                </div>
                <br>

                <label class="label">Ширина окна браузера, px</label>
                <div class="control">
                    <input class="input" id="browser_width" type="text" placeholder="Пиксели"
                           value="{{ .settings.BrowserWidth }}" name="browser_width">
                </div>
                <br>

                <label class="label">Высота окна браузера, px</label>
                <div class="control">
                    <input class="input" id="browser_height" type="text" placeholder="Пиксели"
                           value="{{ .settings.BrowserHeight }}" name="browser_height">
                </div>
                <br>

                <label class="label">URL страницы хоста для скрапинга</label>
                <div class="control">
                    <input class="input" id="host_page_url" type="text" placeholder="URL"
                           value="{{ .settings.HostPageUrl }}" name="host_page_url">
                </div>
                <br><br>

                {{ if eq .settings.FlashAllowed true }}
                    <input class="checkbox" id="flash_allowed" type="checkbox" name="flash_allowed" checked>
                {{ else }}
                    <input class="checkbox" id="flash_allowed" type="checkbox" name="flash_allowed">
                {{ end }}
                <label class="label" for="flash_allowed">Включить подсветку элементов</label>
                <br>

                <label class="label">Количество миганий при подсветке элементов</label>
                <div class="control">
                    <input class="input" id="flash_quantity" type="text" placeholder="Разы"
                           value="{{ .settings.FlashQuantity }}" name="flash_quantity">
                </div>
                <br>

                <label class="label">Период мигания при подсветке элементов, мс</label>
                <div class="control">
                    <input class="input" id="flash_period" type="text" placeholder="Миллисекунды"
                           value="{{ .settings.FlashPeriod }}" name="flash_period">
                </div>
                <br>

                <label class="label">Уровень отладочных сообщений</label>
                <div class="control">
                    <input class="input" id="debug_level" type="text" placeholder="'Info' или 'Debug'"
                           value="{{ .settings.DebugLevel }}" name="debug_level">
                </div>
                <br>

                <label class="label">Минимальная площадь объекта, кв.м.</label>
                <div class="control">
                    <input class="input" id="min_area" type="text" placeholder="Метры квадратные"
                           value="{{ .settings.MinArea }}" name="min_area">
                </div>
                <br>

                <label class="label">Максимальная площадь объекта, кв.м.</label>
                <div class="control">
                    <input class="input" id="max_area" type="text" placeholder="Метры квадратные"
                           value="{{ .settings.MaxArea }}" name="max_area">
                </div>
                <br>

                <label class="label">Минимальный срок аренды, лет</label>
                <div class="control">
                    <input class="input" id="min_rental_period" type="text" placeholder="Лет"
                           value="{{ .settings.MinRentalPeriod }}" name="min_rental_period">
                </div>
                <br>

                <label class="label">Тип имущества</label>
                <div class="control">
                    <input class="input" id="property_type" type="text" placeholder="Например, 'Помещение'"
                           value="{{ .settings.PropertyType }}" name="property_type">
                </div>
                <br>

                <label class="label">Вид договора</label>
                <div class="control">
                    <input class="input" id="contract_type" type="text" placeholder="Например, 'Договор аренды'"
                           value="{{ .settings.ContractType }}" name="contract_type">
                </div>
                <br>

                <label class="label">Страна</label>
                <div class="control">
                    <input class="input" id="country" type="text" placeholder="Например, 'РОССИЯ'"
                           value="{{ .settings.Country }}" name="country">
                </div>
                <br>

                <label class="label">Местоположение имущества (город)</label>
                <div class="control">
                    <input class="input" id="property_location" type="text" placeholder="Например, 'Москва (г)'"
                           value="{{ .settings.PropertyLocation }}" name="property_location">
                </div>
                <br>

                <label class="label">Столбец, по которому сортировать</label>
                <div class="control">
                    <input class="input" id="sort_field_name" type="text"
                           placeholder="Например, 'Коэффициент доходности'"
                           value="{{ .settings.SortFieldName }}" name="sort_field_name">
                </div>
                <br>

                <label class="label">Средняя стоимость аренды, рублей за кв.м. в месяц</label>
                <div class="control">
                    <input class="input" id="average_rental" type="text" placeholder="Рубли"
                           value="{{ .settings.AverageRental }}" name="average_rental">
                </div>
                <br>

                <label class="label">Количество доходных месяцев в году</label>
                <div class="control">
                    <input class="input" id="profit_months" type="text" placeholder="Количество месяцев"
                           value="{{ .settings.ProfitMonths }}" name="profit_months">
                </div>
                <br>

                <label class="label">Стоимость предварительного ремонта, рублей за кв.м.</label>
                <div class="control">
                    <input class="input" id="prior_repair" type="text" placeholder="Рубли"
                           value="{{ .settings.PriorRepair }}" name="prior_repair">
                </div>
                <br>

                <label class="label">Стоимость регистрации договора, рублей</label>
                <div class="control">
                    <input class="input" id="contract_registration" type="text" placeholder="Рубли"
                           value="{{ .settings.ContractRegistration }}" name="contract_registration">
                </div>
                <br>

                <label class="label">Мелкие расходы на запуск объекта, рублей</label>
                <div class="control">
                    <input class="input" id="running_cost" type="text" placeholder="Рубли"
                           value="{{ .settings.RunningCost }}" name="running_cost">
                </div>
                <br>

                <label class="label">Стоимость годовой страховки, рублей</label>
                <div class="control">
                    <input class="input" id="yearly_insurance" type="text" placeholder="Рубли"
                           value="{{ .settings.YearlyInsurance }}" name="yearly_insurance">
                </div>
                <br>

                <label class="label">Стоимость отопления, рублей за кв.м. в месяц</label>
                <div class="control">
                    <input class="input" id="monthly_heating" type="text" placeholder="Рубли"
                           value="{{ .settings.MonthlyHeating }}" name="monthly_heating">
                </div>
                <br>

                <label class="label">Обслуживание ЖЭКом в месяц, рублей за кв.м. в месяц</label>
                <div class="control">
                    <input class="input" id="housing_office_maintenance" type="text" placeholder="Рубли"
                           value="{{ .settings.HousingOfficeMaintenance }}" name="housing_office_maintenance">
                </div>
                <br>

                <label class="label">Бухгалтерское обслуживание, рублей в месяц</label>
                <div class="control">
                    <input class="input" id="accounting_service" type="text" placeholder="Рубли"
                           value="{{ .settings.AccountingService }}" name="accounting_service">
                </div>
                <br>

                <label class="label">Требуемый, приемлемый коэффициент доходности</label>
                <div class="control">
                    <input class="input" id="required_profit_margin" type="text" placeholder="По умолчанию - 25"
                           value="{{ .settings.RequiredProfitMargin }}" name="required_profit_margin">
                </div>
                <br>

                <div class="field">
                    <div class="control">
                        <button class="button is-success" type="submit">
                            Сохранить настройки
                        </button>
                    </div>
                </div>
            </form>
            </div>
        </div>
    </div>
    <div class="column"></div>
    <div class="column"></div>
</div>

{{ template "footer.tpl" .}}
