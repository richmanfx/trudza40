<!-- Шаблон хедера -->
{{ template "header.tpl" .}}


    <div class="columns">
        <div class="column is-one-fifth">Статистика</div>
        <div class="column">
            Основной контент
            <br>
            <a href="/realty/login">
                Недвижимость
            </a>
        </div>
    </div>

    {{/*<script src="/static/js/reload.min.js"></script>*/}}


<!-- Шаблон футера -->
{{ template "footer.tpl" .}}
