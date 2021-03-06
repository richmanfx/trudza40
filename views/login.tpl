{{ template "header.tpl" .}}


<div class="columns">
    <div class="column"></div>
    <div class="column">
        <form method="post" class="" action="/realty/login-processing">

            <div class="field">
                <label class="label is-medium">Логин:</label>
                <div class="control">
                    <label>
                        <input class="input is-medium is-primary" name="user_name"
                               placeholder=" Имя пользователя" required>
                    </label>
                </div>
            </div>

            <div class="field">
                <label class="label is-medium">Пароль: </label>
                <div class="control">
                    <label>
                        <input class="input is-medium is-primary" name="user_password"
                               placeholder=" Пароль" type="password" required>
                    </label>
                </div>
            </div>

            <div class="field">
                <div class="control">
                    <button class="button is-success">
                        Войти
                    </button>
                </div>
            </div>
        </form>
    </div>
    <div class="column"></div>
</div>

{{ template "footer.tpl" .}}
