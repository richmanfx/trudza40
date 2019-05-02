<!doctype html>
<html lang="ru">
<head>
    <meta charset="utf-8">
    <title>Realty Objects</title>
    <link rel="stylesheet" type="text/css" href="/static/css/real.css">
</head>


<body>

    <a href="http://localhost:8081">Сначала!</a>

    <h1>Объекты недвижимости</h1>

    <h2>Статистические:</h2>
        <ul>
            <li>Средняя стоимость аренды: {{ .settings.AverageRental }} руб/кв.м. в месяц </li>
            <li>Количество доходных месяцев в году: {{ .settings.ProfitMonths }}</li>
        </ul>

    <h2>Разовые затраты:</h2>
        <ul>
            <li>Стоимость регистрации договора: {{ .settings.ContractRegistration }} рублей </li>
            <li>Расходы на запуск объекта: {{ .settings.RunningCost }} рублей </li>
        </ul>

    <h2>Расценки:</h2>
        <ul>
            <li>Отопление: {{ .settings.MonthlyHeating }} руб/кв.м. в месяц </li>
            <li>Обслуживание ЖЭКом: {{ .settings.HousingOfficeMaintenance }} руб/кв.м. в месяц </li>
            <li>Бухгалтерское обслуживание: {{ .settings.AccountingService }} руб/мес. </li>
            <li>Предварительный ремонт: {{ .settings.PriorRepair }} руб/кв.м. </li>
        </ul>

    <br><br>
    <h2>Основные параметры</h2>

    <table class="real-table">
        <caption>{{ .settings.PropertyLocation }}, {{ .settings.ContractType }}</caption>

        {{/* Заголовки столбцов */}}

            {{ range .result}}
                <tr>
                    <td>{{ .NotificationNumber }}</td>
                    <td>{{ .ProfitMargin }}</td>
                    <td>{{ .Area }}</td>
                </tr>
            {{ end }}


    </table>

</body>
</html>
