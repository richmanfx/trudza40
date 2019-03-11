<div class="change-password-modal modal">
    <div class="modal-background">
    </div>
    <div class="modal-content">
        <div class="box box-change">
            <span class="has-text-danger is-size-4 has-text-weight-bold">Изменение пароля пользователя</span>
            <hr>

            <form id="id_change_password_form" method="post" action="/realty/change-password">

                <label class="label is-medium" for="id_login_change"> Пользователь </label>
                <input class="input is-medium is-primary" id="id_login_change" name="login" required>
                <br>
                <br>

                <label class="label is-medium" for="id_new_password"> Новый пароль </label>
                <input class="input is-medium is-primary" id="id_new_password"
                       name="new_password" type="password" required>
                <br>
                <br>

                <input type="hidden" id="id_id_change" name="id">
                <input type="hidden" id="id_full_name_change" name="full_name">

            </form>
        </div>

        <!-- Автофокус в поле ввода в модальном окне -->
        <script>

            // При перемещении мыши по боксу с полями вода
            $('.box-change').hover(function(event) {
                $('#id_new_password', event.target).focus();
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
