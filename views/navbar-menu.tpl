
{{/* Шаблон модального окна "О программе" */}}
{{ template "about-modal.tpl" . }}

<nav class="navbar is-dark" role="navigation" aria-label="main navigation">

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
                    <a class="navbar-item" href="/realty/users-config">
                        Работа с Пользователями
                    </a>
                    <a class="navbar-item" href="/realty/main-settings">
                        Основные настройки
                    </a>
                    <a class="navbar-item" href="/realty/object-parameters">
                        Параметры объекта
                    </a>
                </div>
            </div>


            <a class="navbar-item" id="about-menu">
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



    <script async src="/static/js/menu-work.js"></script>

</nav>
