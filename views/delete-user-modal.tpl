<div class="delete-user-modal modal">
    <div class="modal-background">
    </div>
    <div class="modal-content">

        <div class="box">
            <span class="has-text-danger is-size-4 has-text-weight-bold">Удалить пользователя</span>
            <hr>
            <form id="id_delete_user_form" name="delete_user_form" method="post" action="/delete-user">
                <label class="label is-medium" for="id_login_delete"> Логин </label>
                <input class="input is-medium is-primary" id="id_login_delete" name="login" required>
                <br>
                <br>

                <label class="label is-medium" for="id_full_name"> Полное имя (ФИО) </label>
                <input class="input is-medium is-primary" id="id_full_name" name="full_name" required>
                <br>
                <br>
            </form>
        </div>

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
