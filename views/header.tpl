<!doctype html>
<html lang="ru">
<head>
    <title>{{ .title }}</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    {{/*<script defer src="https://use.fontawesome.com/releases/v5.1.0/js/all.js"></script>*/}}
    <meta charset="UTF-8">



    <!-- Бутстрап -->
{{/*<link rel="stylesheet" href="/css/bootstrap.min.css">*/}}
{{/*<script async src="/js/bootstrap.min.js"></script>*/}}

    <!-- favicon -->
    <link rel="shortcut icon" href="data:image/png;base64,
    iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAABmJLR0QA/wD/AP+gvaeTAAAACXBIWXMAAC4jAAAuIwF4pT92AAAAB3RJTUUH4gkKCS
    sF/D6gVQAAAJtJREFUOMudU9ERxSAIC9zbpVs4mAM4mFt0mvSj5R31LJbyg3omEHIISVzxP0jtmAVb8VcBALkIGAEDItEsePhLnYF9q2xlbP1GopFO
    tgKpHVL7lAQA1Fc3QEaK4kP4bjRb3cCWdXx80joOz/Jv5cLCxm8zuLkQVTP7nubDVk4JUYurweobzZEbaluVIfHLdG6jbAD33DrLJuCOA2M1Weix0t
    NQAAAAAElFTkSuQmCC" />

    <!-- Bulma -->
    <link rel="stylesheet" type="text/css" href="/static/css/bulma.min.css">

    <!-- Стили приложения -->
    <link rel="stylesheet" type="text/css" href="/static/css/trudza40.css">

    <!-- jQuery -->
    <script src="/static/js/jquery.min.js"></script>
</head>

<body>
<div class="hero is-fullheight">

    <div class="columns">
        <div class="column is-one-fifth">
            <div class="navbar-brand" style="padding: 20px;">
                {{/*<div style="width:180px; alignment: left; margin:auto; padding: 10px;">*/}}
                    <a class="" href="//trudza40.ru">
                        <img src="/static/img/40+_trud.svg.png" width="160" height="160">
                    </a>
                    {{/*<p style="line-height:50px;">Место для меню</p>*/}}
                {{/*</div>*/}}
            </div>

        </div>
        <div class="column">Заголовок. Основная форма поиска.</div>
    </div>

    <!-- Шаблон меню -->
