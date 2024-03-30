<script>
    import { messages } from "@stores/chat";
    import { afterUpdate, onMount } from "svelte";
    import Message from "./Message.svelte";

    /** @type {HTMLDivElement}*/
    let chat_messages_element;
    export let is_awaiting_response = false;
    let chat_messages_updated_unsubscriber;

    afterUpdate(() => {
        if (chat_messages_element === undefined) return;

        chat_messages_element.scrollTo(0, chat_messages_element.scrollHeight) 
    });

    onMount(() => {
        chat_messages_updated_unsubscriber = messages.subscribe(handleChatMessagesUpdated);
    });

    const handleChatMessagesUpdated = new_messages => {
        if (new_messages.length < 2) return;

        is_awaiting_response = new_messages[new_messages.length - 1].sender === 'user';
    }


</script>

<ul bind:this={chat_messages_element} id="messages-container">
    {#if $messages.length > 0}
        {#each $messages as message}
            <Message {message} />
        {/each}
        {#if is_awaiting_response}  
            <Message is_typing_message={true} />
        {/if}
    {/if}
</ul>

<style>
    #messages-container {
        position: relative;
        width: 100%;
        overflow-y: auto;
        height: 100%;
        display: flex;
        flex-direction: column;
        margin: 0;
        padding: calc(var(--connection-status-bar-height) + var(--vspacing-2)) var(--vspacing-3) var(--vspacing-3) var(--vspacing-3);
        row-gap: var(--vspacing-3);
        list-style: none;
        z-index: var(--z-index-1);
    }

    #messages-container::-webkit-scrollbar {
        color: var(--grey-5);
        width: var(--vspacing-1);
        opacity: .1 !important;
    }
    
    #messages-container::-webkit-scrollbar-thumb {
        background: var(--grey-5);
        border-radius: calc(.27777 * var(--border-radius-2));
        opacity: .1;
    }
    
    #messages-container::-webkit-scrollbar-track {
        background: transparent;
        margin-right: var(--vspacing-1);
        padding: var(--vspacing-1);
    }
</style>