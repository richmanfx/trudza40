<!-- Шаблон хедера -->
{{ template "realty_header.tpl" .}}

<div class="modal is-active" >
    <div class="modal-background">
    </div>
    <div class="modal-content">

        <div class="box">
            <article class="message">
                <div class="message-header">
                    <h1>{{ .message1 }}</h1>
                </div>
                <div class="message-body">
                    {{ .message2 }}
                </div>
                <div class="message-body">
                    {{ .message3 }}
                </div>
                    <div class="has-text-centered">
                        <button class="button is-success ">
                            <a href="/login"> В начало </a>
                        </button>
                    </div>
            </article>
        </div>
    </div>
    <button class="modal-close is-large" aria-label="close">
        <span>&times;</span>
    </button>
</div>

<!-- Шаблон футера -->
{{ template "footer.tpl" .}}
