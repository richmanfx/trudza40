<div class="modal   ">
    <div class="modal-background">
    </div>
    <div class="modal-content">
        Удалить пользователя

        <div class="box">
            <form id="id_delete_user_form" name="delete_user_form" method="post" action="/delete-user">
                <label class="label is-medium" for="id_login"> Логин </label>
                <input class="input is-medium is-primary" id="id_login" name="login" required>
                <br>
                <br>

                <label class="label is-medium" for="id_full_name"> Полное имя (ФИО) </label>
                <input class="input is-medium is-primary" id="id_full_name" name="full_name" required>
                <br>
                <br>
            </form>
        </div>

        <!-- Автофокус в поле ввода в модальном окне -->
        <script>
            $('#deleteUser').on('shown.bs.modal', function (e) {
                $('#id_login', e.target).focus();
            });
        </script>

        <div class="modal-footer">
            <button class="button is-link" form="id_delete_user_form">
                Удалить
            </button>

            <button class="button is-link" form="id_delete_user_form" type="reset">
                Очистить форму
            </button>

            <a class="button is-link" data-dismiss="modal" href="/users-config">
                <strong>Отмена</strong>
            </a>

    </div>
</div>

<!-- Шаблон футера -->
{{ template "footer.tpl" .}}
