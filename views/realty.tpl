<!-- Шаблон хедера -->
{{ template "header.tpl" .}}


<div class="columns">
    <div class="column is-one-fifth">Недвижимость</div>
    <div class="column">
        Основная инфо про недвижимость
        <p>
            {{.Website}} {{.Email}}
        </p>

        <a href="/login">
            Залогиниться
        </a>

    </div>
</div>


<!-- Шаблон футера -->
{{ template "footer.tpl" .}}