<script>
    export let message;
    export let is_typing_message = false;
</script>

{#if !is_typing_message}
    <li class="message" class:external-message={message.sender !== 'user'}>
        <span class="message-sender">{message.sender === 'user' ? 'You' : 'Valaria'}</span>
        <p class="message-content">
            {@html message.content.replaceAll('\n', '<br/>')}
        </p>
        <div class="message-metadata">
            <span class="metadata-item">
                {message.send_date}
            </span>
        </div>
    </li>
{:else}
    <li id="writing-message" class="message external-message">
        <span class="message-sender">Valaria</span>
        <p class="message-content writing-animation">
            <span>typing</span>
            <span style:animation-delay=".2s" class="dot"></span>
            <span style:animation-delay=".4s" class="dot"></span>
            <span style:animation-delay=".8s" class="dot"></span>
        </p>
        <div class="message-metadata">
            <span class="metadata-item">
                ---
            </span>
        </div>
    </li>
{/if}


<style>
    @keyframes writing-animation {
        0% {
            transform: translateY(0);
            background-color: var(--grey-1);
        }
        28% {
            transform: translateY(-7px);
            background-color: var(--main-5);
        }
        64% {
            transform: translateY(0);
            background-color: var(--main-dark-color-6);
        }
    }

    .message {
        --message-main-color: var(--grey-7);
        --message-accent-color: var(--main-dark);
        --message-text-color: var(--grey-1);

        display: flex;
        max-width: 50%;
        align-self: end;
        flex-direction: column;
        padding: var(--vspacing-1) var(--vspacing-3) var(--vspacing-1) var(--vspacing-2);
        border-radius: var(--border-radius);
        background: var(--message-main-color);
        gap: var(--vspacing-1);
        clip-path: polygon(0 0, 100% 0, 97% 30%, 97% 100%, 0 100%, 0 0);
    }
    
    .message.external-message {
        align-self: start;
        background: var(--message-accent-color);
        padding: var(--vspacing-1) var(--vspacing-2) var(--vspacing-1) var(--vspacing-3);
        clip-path: polygon(0 0, 100% 0, 100% 100%, 3% 100%, 3% 20%, 0 0);
    }

    .message-sender {
        font-family: var(--font-titles);
        font-size: var(--font-size-2);
        color: var(--message-accent-color);
        font-weight: 400;
    }

    .message.external-message .message-sender {
        width: 100%;
        color: var(--message-main-color);
    }

    .message .message-content {
        box-sizing: content-box;
        min-height: 4ch;
        color: var(--message-text-color);
        font-size: var(--font-size-1);
        padding: 0 0 0 var(--vspacing-1);
    }

    .message-metadata {
        display: flex;
        justify-content: flex-end;
        gap: var(--vspacing-1);
    }

    #writing-message.message {
        min-width: 40%;
    }

    #writing-message.message .message-content{
        font-style: italic;
        color: var(--message-text-color);
        font-size: var(--font-size-2);
        min-height: 2ch;
        height: max-content;
        white-space: pre-line;
    }

    #messages-container #writing-message.message .message-content.writing-animation {
        display: flex;
        align-items: flex-end;
        line-height: 1;
        gap: 3px;
    }

    #messages-container #writing-message.message .message-content.writing-animation .dot {
        width: 4px;
        height: 4px;
        border-radius: 50%;
        background-color: var(--message-text-color);
        animation: writing-animation 1.8s infinite ease-in-out;
        transition: background-color .1s ease-in;
    }
</style>