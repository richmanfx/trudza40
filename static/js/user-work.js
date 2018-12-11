/**
 * Created by Александр Ящук (R5AM, Zoer) on 10.12.2018.
 */

/// Подстановка Логина и Полного имени пользователя в поля в модальном окне
$(document).ready(function() {

    // Для "Удалить пользователя"
    // Если кликнули по элементу с соответствующим удалению классом
    $('.delete-user').on('click',function() {
        // alert('Ho-ho!');

        // Сделать модальное окно активным - добавить "is-active"
        $(".delete-user-modal").toggleClass("is-active");

        // Извлечь информацию из "data-*" полей у кликнутого элемента
        let deletedLogin = $(this).data('login');
        let deletedFullName = $(this).data('name');

        // В input-ы вставить значения из "data-*"
        let modal_box = $('.box');
        modal_box.find('#id_login').val(deletedLogin);
        modal_box.find('#id_full_name').val(deletedFullName);
    });

    // Для "Изменить пароль"
    // Если кликнули по элементу с соответствующим изменению пароля классом
    $('.change-password').on('click',function() {

        // Сделать модальное окно активным - добавить "is-active"
        $(".change-password-modal").toggleClass("is-active");

        // Извлечь информацию из "data-*" полей у кликнутого элемента
        let Login = $(this).data('login');
        let deletedFullName = $(this).data('name');

        // В input-ы вставить значения из "data-*"
        let modal_box = $('.box');
        modal_box.find('#id_login').val(Login);
        modal_box.find('#id_full_name').val(deletedFullName);

    });

});
