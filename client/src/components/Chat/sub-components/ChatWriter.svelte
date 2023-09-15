<script>
    import { layout_properties } from "@stores/layout";
    import { ChatDialer } from "@models/Chat";
    import { createEventDispatcher } from "svelte";

    /** @type {ChatDialer}*/
    export let chat_dialer;
    const chat_focus_dispatcher = createEventDispatcher();
    let chat_active = false;

    const handleChatFocus = () => {
        if (chat_active) return;

        chat_active = true;

        chat_focus_dispatcher('chatFocus');
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
</script>


<div id="writer-bar" class:debug={false}>
    <div id="sender-wrapper">
        <input 
            aria-label="Write any doubts you have about my services and my AI assistant will clear them for you. press enter to send the message"
            data-name="message-bar"
            on:keydown={handleMessageBarKeypress}
            on:focus={handleChatFocus}
            class:hide={layout_properties.IS_MOBILE}
            type="text"
            id="sender-writer"
            placeholder="Clear you'r doubts"
        >
        <div id="sender-btn" class="button-1-wrapper">
            <button class="button-1">{ layout_properties.IS_MOBILE ? 'open chat' : 'send'}</button>
        </div>
    </div>
</div>

<style>
        #writer-bar {
        position: relative;
        display: grid;
        width: 100%;
        height: var(--writer-height);
        border-top: 1px solid var(--grey-4);
        place-items: center;
        z-index: var(--z-index-2);
    }

    #sender-wrapper {
        width: 70%;
        display: flex;
        gap: var(--vspacing-2);
    }

    #sender-wrapper input {
        width: 80%;
        font-family: var(--font-read);
        background: var(--grey-8);
        border: none;
        color: var(--grey-1);
        border-radius: var(--border-radius);
        padding: 0 var(--vspacing-2);
        outline: none;
    }

    #sender-wrapper #sender-btn {
        background: var(--main-dark);
        color: var(--grey-1);
    }

    #sender-wrapper #sender-btn .button-1 {
        color: var(--grey-1);
    }

    @media only screen and (max-width: 768px) {
        #writer-bar {
            padding: calc(.75 * var(--vspacing-1));
        }

        #sender-wrapper {
            width: max-content;
            height: 100%;
        }

        #sender-wrapper .button-1-wrapper button {
            height: 100%;
            padding: 0 var(--vspacing-2);
        }
    }
</style>