{{ template "realty_header.tpl" .}}

<div class="columns">
    <div class="column"></div>
    <div class="column is-two-thirds">
        <div class="content has-text-centered">

            <h2>Список пользователей</h2>
            <table class="table is-striped">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Login</th>
                        <th>Полное имя</th>
                        <th>Действия</th>
                    </tr>
                </thead>

                <tbody>
                    {{ range $val := .users }}
                    <tr>
                        <td>{{ $val.Id }}</td>
                        <td>{{ $val.Login }}</td>
                        <td>{{ $val.FullName }}</td>
                        <td>
                            <a data-toggle="modal" data-target="#deleteUser" href="#" title="Удалить пользователя"
                               data-login="{{ $val.Login }}" data-name="{{ $val.FullName }}">
                                <img src="/static/img/delete.png" alt="Delete" width="20" height="20"/>
                            </a>
                            <a data-toggle="modal" data-target="#changePassword" href="#" title="Новый пароль"
                               data-login="{{ $val.Login }}">
                                <img src="/static/img/password.png" alt="Delete" width="20" height="20"/>
                            </a>
                        </td>
                    </tr>
                    {{end }}
                </tbody>
            </table>

            <br><br>

            <a class="button is-link" title="Создать нового пользователя" href="/create-user">
                <strong>Создать нового пользователя</strong>
            </a>

        </div>
    </div>
    <div class="column"></div>
</div>

{{ template "footer.tpl" .}}
