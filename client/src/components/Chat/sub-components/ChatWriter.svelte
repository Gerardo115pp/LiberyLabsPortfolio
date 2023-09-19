<script>
    import { layout_properties } from "@stores/layout";
    import { ChatDialer } from "@models/Chat";
    import { createEventDispatcher } from "svelte";
    import { recaptcha_pk } from "@stores/env";
    import { GetRecaptchaVerificationRequest } from "@libs/HttpRequests";
    import { is_user_human } from "@stores/chat";
    import { is_big_mode_enabled } from "@stores/chat";

    /** @type {ChatDialer}*/
    export let chat_dialer;
    const chat_focus_dispatcher = createEventDispatcher();
    const big_mode_on_dispatcher = createEventDispatcher();
    let chat_active = false;

    const handleChatFocus = () => {
        if (chat_active) return;

        chat_active = true;

        grecaptcha.ready(() => {
            grecaptcha.execute(recaptcha_pk, {action: 'chatFocus'})
                .then(token => {
                    const verification_request = new GetRecaptchaVerificationRequest(token);

                    const on_success = () => {
                        is_user_human.set(true);
                        console.log('User is human');
                        chat_focus_dispatcher('chatFocus');
                    }

                    const on_error = () => {
                        is_user_human.set(false);
                    }

                    verification_request.do(on_success, on_error);
                })
                .catch(err => {
                    console.error(err);
                    is_user_human.set(false);
                })
        })

    }

    const handleMessageBarKeypress = e => {
        if (e.key !== 'Enter') return;  

        if (chat_dialer === undefined) {
            alert('Chat dialer is not defined');
        }
        

        /** @type string */
        let new_message_content = e.target.value;
        e.target.value = '';

        if (new_message_content.length < 3) return;

        chat_dialer.sendMessage(new_message_content);
    }

    const handleBigMode = () => {
        is_big_mode_enabled.set(!$is_big_mode_enabled);
    }
</script>

<div id="writer-bar" class:big-mode={$is_big_mode_enabled} class:debug={false}>
    <div id="sender-wrapper">
        <textarea 
            aria-label="Write any doubts you have about my services and my AI assistant will clear them for you. press enter to send the message"
            data-name="message-bar"
            on:keydown={handleMessageBarKeypress}
            on:focus={handleChatFocus}
            class:hide={layout_properties.IS_MOBILE && !$is_big_mode_enabled}
            type="text"
            id="sender-writer"
            placeholder="Clear you'r doubts, ask us anything!"
            user
        />
        <div id="sender-btn" class:hide={!layout_properties.IS_MOBILE} class="button-1-wrapper">
            <button on:click={handleBigMode} class="button-1">{ $is_big_mode_enabled ? 'close chat' : 'open chat'}</button>
        </div>
    </div>
    {#if !layout_properties.IS_MOBILE || $is_big_mode_enabled}
         <p id="grecaptcha-policy">
             This component is protected by reCAPTCHA and the Google
             <a href="https://policies.google.com/privacy">Privacy Policy</a> and
             <a href="https://policies.google.com/terms">Terms of Service</a> apply.
         </p>
    {/if}
</div>

<style>
    #writer-bar {
        position: relative;
        display: grid;
        width: 100%;
        height: var(--writer-height);
        border-top: 1px solid var(--grey-4);
        justify-items: center;
        padding: var(--vspacing-2) var(--vspacing-2);
        row-gap: var(--vspacing-1);
        z-index: var(--z-index-2);
    }

    #sender-wrapper {
        width: 70%;
        display: flex;
        gap: var(--vspacing-2);
    }

    #sender-wrapper textarea {
        width: 100%;
        font-family: var(--font-read);
        background: var(--grey-8);
        border: none;
        color: var(--grey-1);
        border-radius: var(--border-radius);
        padding: var(--vspacing-1) var(--vspacing-2);
        resize: none;
        outline: none;
    }

    #sender-wrapper #sender-btn {
        background: var(--main-dark);
        color: var(--grey-1);
    }

    #sender-wrapper #sender-btn .button-1 {
        color: var(--grey-1);
    }

    #grecaptcha-policy {
        color: var(--grey-6);
        font-size: var(--font-size-fineprint);
        justify-self: end;
        align-self: last baseline;
    }

    #grecaptcha-policy a {
        color: var(--main-8);
        text-decoration: none;
    }

    @media only screen and (max-width: 768px) {
        #writer-bar {
            padding: calc(1.2 * var(--vspacing-1));
        }

        #sender-wrapper {
            display: flex;
            justify-content: center;
            width: 90%;
            height: 100%;
        }

        #sender-wrapper .button-1-wrapper button {
            height: 100%;
            padding: 0 var(--vspacing-2);
        }
    }
</style>