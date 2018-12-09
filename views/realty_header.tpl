<!doctype html>
<html lang="ru">
<head>
    <title>{{ .title }}</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta charset="UTF-8">

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

<!-- Шаблон основного меню -->
{{ template "navbar-menu.tpl" . }}
