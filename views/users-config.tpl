{{ template "realty_header.tpl" .}}

{{ template "delete-user-modal.tpl" . }}
{{ template "change-password-modal.tpl" . }}

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
                            <a class="delete-user" title="Удалить пользователя"
                               data-login="{{ $val.Login }}" data-name="{{ $val.FullName }}">
                                <img src="/static/img/delete.png" alt="Delete user" width="20" height="20"/>
                            </a>
                            <a class="change-password" title="Изменить пароль"
                               data-login="{{ $val.Login }}" data-name="{{ $val.FullName }}">
                                <img src="/static/img/password.png" alt="Change password" width="20" height="20"/>
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

<script async src="/static/js/user-work.js"></script>

{{ template "footer.tpl" .}}
