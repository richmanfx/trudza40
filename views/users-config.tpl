{{ template "realty_header.tpl" .}}

<div class="columns">
    <div class="column ">
        <div class="content has-text-centered">

            {{/*<a type="button" class="btn btn-success btn-lg" data-toggle="modal" data-target="#createUser" href="#">*/}}
                {{/*Создать нового пользователя*/}}
            {{/*</a>*/}}

            <a class="button is-link" title="Создать нового пользователя" href="/create-user">
                <strong>Создать нового пользователя</strong>
            </a>

            <a class="button is-link" title="Удалить пользователя">
                <strong>Удалить пользователя</strong>
            </a>

            <a class="button is-link" title="Редактировать пользователя">
                <strong>Редактировать пользователя</strong>
            </a>

        </div>
    </div>
</div>

{{ template "footer.tpl" .}}
