{{ template "realty_header.tpl" .}}


<div class="columns">
    <div class="column"></div>
    <div class="column ">
        <div class="content has-text-centered">

            <table class="table">
                <tr>
                    <td> 1 </td>
                    <td><strong> torgi.gov.ru </strong></td>
                    <td>
                        <a class="button" title="Параметры для скрапинга" href="/realty/settings_torgi_gov_ru">
                            Параметры
                        </a>
                    </td>
                    <td>
                        <a class="button" title="Скрапинг 'torgi.gov.ru'" href="/realty/scraping_torgi_gov_ru">
                            Скрапинг
                        </a>
                    </td>
                    <td>
                        <a class="button" title="Последний отчёт скрапинга" href="/realty/last_report_torgi_gov_ru">
                            Отчёт
                        </a>
                    </td>
                </tr>

                <tr>
                    <td> 2 </td>
                    <td><strong> второй.сайт.ru </strong></td>
                    <td>
                        <a class="button">
                            Параметры
                        </a>
                    </td>
                    <td>
                        <a class="button" title="Скрапинг 'второй.сайт.ru'" href="#">
                            Скрапинг
                        </a>
                    </td>
                    <td>
                        <a class="button">
                            Отчёт
                        </a>
                    </td>
                </tr>

            </table>
        </div>
    </div>
    <div class="column"></div>
</div>


<!-- Шаблон футера -->
{{ template "footer.tpl" .}}
