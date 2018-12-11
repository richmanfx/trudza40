/**
 * Created by Александр Ящук (R5AM, Zoer) on 11.12.2018.
 */

/// Обработка пунктов в navbar меню


$(document).ready(function() {

    // Если кликнули по пункту "О программе"
    $('#about-menu').on('click',function() {
        // alert('Ho-ho!');

        // Сделать модальное окно активным - добавить "is-active"
        $(".about-modal").toggleClass("is-active");
    });

    $('#ok_button').on('click',function() {

        // Сделать модальное окно НЕактивным - убрать "is-active"
        $(".about-modal").removeClass("is-active");

    });

});
