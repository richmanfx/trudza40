<!-- Шаблон хедера -->
{{ template "realty_header.tpl" .}}

<div class="modal is-active" >
    <div class="modal-background">
    </div>
    <div class="modal-content">
        Создать нового пользователя

        <div class="box">
            <form id="id_create_user_form" name="create_user_form" method="post" action="/create-user-in-db">

                <label class="label is-medium" for="id_login"> Логин </label>
                <input id="id_login" class="input is-medium is-primary" name="login"
                       placeholder=" Логин" required>
                <br>
                <br>

                <label class="label is-medium" for="id_full_name"> Полное имя (ФИО) </label>
                <input id="id_full_name" class="input is-medium is-primary" name="full_name"
                       placeholder=" ФИО" required>
                <br>
                <br>

                <label class="label is-medium" for="id_password"> Пароль </label>
                <input id="id_password" class="input is-medium is-primary" name="password" placeholder=" Пароль"
                       type="password" required>
                <br>
                <br>

                <div class="field">
                    <p class="control">
                        <button class="button is-link" form="id_create_user_form" type="submit">
                            Создать
                        </button>
                        <button class="button is-link" form="id_create_user_form" type="reset">
                            Очистить форму
                        </button>
                        <a class="button is-link" data-dismiss="modal" href="/users-config">
                            <strong>Отмена</strong>
                        </a>
                    </p>
                </div>

                {{/*<label> Права </label>*/}}
                {{/*<div class="with-border">*/}}

                    {{/*<table>*/}}
                        {{/*<tr>*/}}
                            {{/*<td class="permission">*/}}
                                {{/*<input class="permission" title="Создание" name="create_permission" type="checkbox">*/}}
                            {{/*</td>*/}}
                            {{/*<td>*/}}
                                {{/*<label class="permission">Создание</label>*/}}
                            {{/*</td>*/}}
                        {{/*</tr>*/}}

                        {{/*<tr>*/}}
                            {{/*<td class="permission">*/}}
                                {{/*<input class="permission" title="Редактирование" name="edit_permission" type="checkbox">*/}}
                            {{/*</td>*/}}
                            {{/*<td>*/}}
                                {{/*<label class="permission">Редактирование</label>*/}}
                            {{/*</td>*/}}
                        {{/*</tr>*/}}

                        {{/*<tr>*/}}
                            {{/*<td class="permission">*/}}
                                {{/*<input class="permission" title="Удаление" name="delete_permission" type="checkbox">*/}}
                            {{/*</td>*/}}
                            {{/*<td>*/}}
                                {{/*<label class="permission">Удаление</label>*/}}
                            {{/*</td>*/}}
                        {{/*</tr>*/}}

                        {{/*<tr>*/}}
                            {{/*<td class="permission">*/}}
                                {{/*<input class="permission" title="Конфигурация" name="config_permission" type="checkbox">*/}}
                            {{/*</td>*/}}
                            {{/*<td>*/}}
                                {{/*<label class="permission">Конфигурация</label>*/}}
                            {{/*</td>*/}}
                        {{/*</tr>*/}}

                        {{/*<tr>*/}}
                            {{/*<td class="permission">*/}}
                                {{/*<input class="permission" title="Работа с пользователями" name="users_permission" type="checkbox">*/}}
                            {{/*</td>*/}}
                            {{/*<td>*/}}
                                {{/*<label class="permission">Работа с пользователями</label>*/}}
                            {{/*</td>*/}}
                        {{/*</tr>*/}}

                    {{/*</table>*/}}
                {{/*</div>*/}}

            </form>
        </div>
    </div>
</div>

<!-- Шаблон футера -->
{{ template "footer.tpl" .}}
