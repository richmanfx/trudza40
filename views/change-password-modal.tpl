<div class="change-password-modal modal">
    <div class="modal-background">
    </div>
    <div class="modal-content">
        <div class="box">
            <span class="has-text-danger is-size-4 has-text-weight-bold">Изменение пароля пользователя</span>
            <hr>

            <form id="id_change_password_form" method="post" action="/change-password">

                <label class="label is-medium" for="id_login"> Пользователь </label>
                <input class="input is-medium is-primary" id="id_login" name="login" required>
                <br>
                <br>

                <label class="label is-medium" for="id_new_password"> Новый пароль </label>
                <input class="input is-medium is-primary" id="id_new_password"
                       name="new_password" type="password" required>
                <br>
                <br>
                <input type="hidden" name="full_name" id="id_full_name">
            </form>
        </div>

        <!-- Автофокус в поле ввода в модальном окне -->
        {{/*TODO: Отладить - по id не находит?  */}}
        <script>
            $('#deleteUser').on('shown.bs.modal', function (e) {
                $('#id_login', e.target).focus();
            });
        </script>

        <div class="modal-footer">
            <button class="button is-link" form="id_change_password_form">
                Сохранить
            </button>

            <a class="button is-link" data-dismiss="modal" href="/users-config">
                <strong>Отмена</strong>
            </a>

        </div>
    </div>

    <!-- Шаблон футера -->
    {{ template "footer.tpl" .}}
