<nav class="navbar is-dark" role="navigation" aria-label="main navigation">
    {{/*<div class="navbar-brand">*/}}
        {{/*<a class="navbar-item" href="https://bulma.io">*/}}
            {{/*<img src="https://bulma.io/images/bulma-logo.png" width="112" height="28">*/}}
        {{/*</a>*/}}

        {{/*<a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false"*/}}
           {{/*data-target="navbarBasicExample">*/}}
            {{/*<span aria-hidden="true"></span>*/}}
            {{/*<span aria-hidden="true"></span>*/}}
            {{/*<span aria-hidden="true"></span>*/}}
        {{/*</a>*/}}
    {{/*</div>*/}}

    <div id="navbarBasicExample" class="navbar-menu">
        <div class="navbar-start">

            <a class="navbar-item" href="/realty">
                В начало
            </a>

            <div class="navbar-item has-dropdown is-hoverable">

                <a class="navbar-item">
                    Конфигурация
                </a>
                <div class="navbar-dropdown">
                    <a class="navbar-item" href="/users-config">
                        Работа с Пользователями
                    </a>
                    <a class="navbar-item" href="/main-settings">
                        Основные настройки
                    </a>
                    <a class="navbar-item" href="/object-parameters">
                        Параметры объекта
                    </a>
                </div>
            </div>


            <a class="navbar-item">
                О программе
            </a>

        </div>

        <div class="navbar-end">
            <div class="navbar-item">
                <div class="buttons">

                    <a class="button is-dark" title="Выход">
                        <strong>Выход</strong>
                    </a>


                </div>
            </div>
        </div>
    </div>
</nav>
