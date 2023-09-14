<script>
    import { layout_properties } from "@stores/layout";
    import { token as chat_token, is_available, messages } from "@stores/chat";
    import { ChatDialer } from "@models/Chat";
    import { onMount, afterUpdate } from "svelte";

    let chat_dialer;
    /** @type {HTMLDivElement}*/
    let chat_messages_element;
    
    onMount(() => {
        chat_dialer = new ChatDialer(); // To styleize the chat component enable this line
    });

    afterUpdate(() => {
        if (chat_messages_element === undefined) return;

        chat_messages_element.scrollTo(0, chat_messages_element.scrollHeight) 
    });

    const handleChatFocus = () => {
        // chat_dialer = new ChatDialer(); // To styleize the chat component disable this line
    }

    const handleMessagebarKeypress = e => {
        if (e.key !== 'Enter') return;  

        /** @type string */
        let new_message_content = e.target.value;
        e.target.value = '';

        if (new_message_content.length < 3) return;

        chat_dialer.sendMessage(new_message_content);     
    }

</script>

<div id="chat-component">
    <div id="status-bar">
        <svg id="sb-controls" viewBox="0 0 48 14">
            <g>
                <ellipse cx="6.65505" cy="6.72339" rx="6.65505" ry="6.65503" fill="#E50D00"/>
                <ellipse cx="23.692" cy="6.72339" rx="6.65505" ry="6.65503" fill="#FDCF00"/>
                <ellipse cx="40.7288" cy="6.72339" rx="6.65505" ry="6.65503" fill="#A9FF06"/>
            </g>
        </svg>
    </div>
    <div id="chat-content" style:position="relative">
        <div class="chat-bg-star">
            <svg width="{layout_properties.IS_MOBILE ? layout_properties.VIEWPORT_WIDTH * .3 : 355}" height="{layout_properties.IS_MOBILE ? (layout_properties.VIEWPORT_WIDTH * .3) * 0.940 : 334}" viewBox="0 0 355 334" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M180.344 165.358L223.772 96.3071L229.588 219.714L180.344 165.358Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M128.996 225.124L176.322 165.357H180.345L229.589 219.713L229.64 220.838L183.361 212.764L177.825 212.435L128.996 225.124Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M132.013 93.835L176.322 165.358L128.996 225.125L130.645 182.133L132.013 93.835Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M132.013 93.835L176.488 122.249L180.345 122.069L223.773 96.3071L180.345 165.358H176.322L132.013 93.835Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M77.8032 114.112L92.2343 135.99L175.306 49.9419L77.8032 114.112Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M182.522 49.9419L265.127 129.889L280.947 108.998L182.522 49.9419Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M132.013 93.835L131.204 147.321L82.9868 145.846L132.013 93.835Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M177.483 49.9419H175.306L132.013 93.8361L164.69 115.364L177.483 49.9419Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M177.483 49.9419H178.935L193.936 114.346L180.345 122.07L176.488 122.25L164.69 115.364L177.483 49.9419Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M228.179 93.73L193.936 114.346L178.935 49.9419H182.522L228.179 93.73Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M131.204 147.321L82.9869 145.846L81.4629 147.321L130.645 182.132L131.204 147.321Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M128.996 225.127L81.4629 147.322L130.645 182.134L128.996 225.127Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M81.4629 147.322L49.9468 180.108L121.947 231.312L164.234 258.103L128.996 225.127L81.4629 147.322Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M128.996 225.125L177.825 212.435L179.111 264.913L179.837 266.589L160.803 254.886L128.996 225.125Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M177.825 212.435L179.111 264.913L179.837 266.589L181.35 264.913L183.361 212.764L177.825 212.435Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M183.361 212.765L229.64 220.839L230.19 223.576L227.162 229.762L215.085 242.134L179.836 266.59L181.35 264.914L183.361 212.765Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M225.929 142.175H279.081L228.178 93.7295L223.772 96.3078L225.929 142.175Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M225.929 142.176H279.081L227.66 178.929L225.929 142.176Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M227.66 178.929L229.64 220.839L279.081 142.176L227.66 178.929Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M279.081 142.176L308.731 173.083L232.201 231.821L215.085 242.134L227.163 229.763L230.19 223.577L229.64 220.839L279.081 142.176Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M179.836 266.59L215.085 242.134L232.201 231.821L239.468 226.25L277.008 308.596V312.511L179.836 266.59Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M283.57 192.394L277.029 308.331L239.489 225.975L283.57 192.394Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M179.837 266.59L118.816 229.083L82.1677 318.389L179.837 266.59Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M118.816 229.083L78.5703 200.467L82.1677 318.389L118.816 229.083Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M136.554 289.55L179.837 332.818L221.481 286.271L179.837 266.589L136.554 289.55Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M63.0404 196.785L37.869 231.82L80.1358 251.725L79.576 233.422L63.0404 196.785Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M280.657 243.948L324.469 224.255L297.11 192.669L294.632 195.757L281.652 226.674L280.657 243.948Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M77.8032 114.112L81.4628 3.50098L114.876 89.7509L77.8032 114.112Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M114.876 89.7509L146.589 68.9126L81.4629 3.50098L114.876 89.7509Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M81.4629 3.50098L175.306 49.9416L146.589 68.9126L81.4629 3.50098Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M182.522 49.9411L271.472 0.98584L212.493 67.989L182.522 49.9411Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M212.493 67.9889L271.472 0.98584L242.703 86.0899L212.493 67.9889Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M242.703 86.0899L280.947 108.997L271.472 0.98584L242.703 86.0899Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M154.261 39.5223L175.306 12.668L200.975 39.7769L182.522 49.9414H175.306L154.261 39.5223Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M175.306 49.9414V12.668L182.522 49.9414H175.306Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M78.3319 98.1431L77.8031 114.111L92.2342 135.989L79.4412 149.433L1.61505 125.676L78.3319 98.1431Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M79.4412 149.432L49.9467 180.106L1.61511 125.676L79.4412 149.432Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M265.127 129.889L280.947 108.997L279.61 93.7295L353.03 121.56L279.081 142.175L265.127 129.889Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M279.081 142.175L308.731 173.082L353.03 121.56L279.081 142.175Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M308.731 173.083L283.57 192.393L281.652 226.675L294.632 195.757L297.11 192.669L308.731 173.083Z" stroke="#7D2900" stroke-dasharray="2 2"/>
                <path d="M79.5761 233.422L63.0405 196.786L49.9468 180.106L78.5705 200.467L79.5761 233.422Z" stroke="#7D2900" stroke-dasharray="2 2"/>
            </svg>   
        </div>
        {#if chat_dialer !== undefined}
            <div id="chat-connection-status">
                <ul id="status-labels">
                    <li class="status-label">
                        <span class="status-name">Valaria:</span>
                        <span class="status-value" class:status-ok={$is_available} >{$is_available ? 'Available' : 'Unavailable'}</span>
                    </li>
                    {#if $is_available}
                            <li class="status-label">
                                <span class="status-name">Connection status:</span>
                                <span class="status-value" class:status-ok={$chat_token !== ''}>{$chat_token !== '' ? 'Connected' : 'Connecting...'}</span>
                            </li>
                    {/if}
                </ul>
            </div>
            <ul bind:this={chat_messages_element} id="messages-container">
            {#if $messages.length > 0}
                {#each $messages as message}
                    {@debug message}
                    <li class="message" class:external-message={message.sender !== 'user'}>
                        <span class="message-sender">{message.sender === 'user' ? 'You' : 'Valaria'}</span>
                        <p class="message-content">
                            {message.content}
                        </p>
                        <div class="message-metadata">
                            <span class="metadata-item">
                                {message.send_date}
                            </span>
                        </div>
                    </li>
                {/each}
            {/if}
            </ul>
        {/if}
    </div>
    <div id="writer-bar" class:debug={false}>
        <div id="sender-wrapper">
            <input aria-label="Write any doubts you have about my services and my AI assistant will clear them for you. press enter to send the message" data-name="message-bar" on:keydown={handleMessagebarKeypress} on:focus={handleChatFocus} class:hide={layout_properties.IS_MOBILE} type="text" id="sender-writer" placeholder="Clear you'r doubts">
            <div id="sender-btn" class="button-1-wrapper">
                <button class="button-1">{ layout_properties.IS_MOBILE ? 'open chat' : 'send'}</button>
            </div>
        </div>
    </div>
</div>

<style>
    #chat-component {
        --status-bar-height: 7%;
        --writer-height: 11%;
        --chat-content-height: calc(100% - calc(var(--status-bar-height) + var(--writer-height)));

        position: relative;

        width: 100%;
        height: calc(var(--vspacing-8) * 0.655859375);
        background-color: var(--grey-9);
        border-radius: var(--border-radius-2);
        box-shadow: 9px -1px 37px 3px rgba(0, 0, 0, 0.25);
        overflow: hidden;
    }

    #status-bar {
        box-sizing: border-box;
        display: flex;
        height: var(--status-bar-height);
        background: var(--main);
        align-items: center;
        padding: 0 var(--vspacing-3);
        z-index: var(--z-index-2);
    }

    #sb-controls {
        height: 34%;
    }

    #chat-content {
        --connection-status-bar-height: 50px;
        height: var(--chat-content-height);
    }

    #chat-content .chat-bg-star {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        opacity: .5;
    }

    #chat-content #chat-connection-status {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: var(--connection-status-bar-height);
        background: var(--accent-9);
        display: flex;
        align-items: center;
        padding: 0 var(--vspacing-3);
        z-index: var(--z-index-2);
    }

    #messages-container {
        position: relative;
        width: 100%;
        overflow-y: auto;
        height: 100%;
        display: flex;
        flex-direction: column;
        margin: 0;
        padding: calc(var(--connection-status-bar-height) + var(--vspacing-2)) var(--vspacing-3) 0;
        row-gap: var(--vspacing-3);
        list-style: none;
        z-index: var(--z-index-1);
    }
    
    #messages-container .message {
        display: flex;
        max-width: 50%;
        align-self: end;
        flex-direction: column;
        padding: var(--vspacing-1) var(--vspacing-3) var(--vspacing-1) var(--vspacing-2);
        border-radius: var(--border-radius);
        background: var(--grey-8);
        gap: var(--vspacing-1);
        clip-path: polygon(0 0, 100% 0, 97% 30%, 97% 100%, 0 100%, 0 0);
    }
    
    #messages-container .message.external-message {
        align-self: start;
        background: var(--main-7);
        padding: var(--vspacing-1) var(--vspacing-2) var(--vspacing-1) var(--vspacing-3);
        clip-path: polygon(0 0, 100% 0, 100% 100%, 3% 100%, 3% 20%, 0 0);
    }

    .message-sender {
        font-weight: 500;
        color: var(--main);
    }

    #messages-container .message.external-message .message-sender {
        width: 100%;
        color: var(--danger-9);
        text-align: end;
    }

    #messages-container .message .message-content {
        box-sizing: content-box;
        min-height: 4ch;
        font-size: var(--font-size-1);
        padding: 0 0 0 var(--vspacing-2);
        
    }
    
    #messages-container .message.external-message .message-content {
        padding: 0 var(--vspacing-2) 0 0;
    }

    .message-metadata {
        display: flex;
        justify-content: flex-end;
        gap: var(--vspacing-1);
    }

    #status-labels {
        display: flex;
        width: 100%;
        gap: var(--vspacing-3);
        list-style: none;
    }

    .status-value {
        font-weight: 500;
        color: var(--danger-7);
    }

    .status-value.status-ok {
        color: var(--success-5);
    }

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
        background: var(--main-dark-color-5);
    }

    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 768px) {
            #chat-content {
                opacity: 0;
            }

            #writer-bar {
                padding: calc(.75 * var(--vspacing-1));
            }

            #sender-wrapper {
                width: max-content;
                height: 100%;
            }

            #chat-component {
                height: calc(var(--vspacing-8) * 0.455859375);
            }

            #sender-wrapper .button-1-wrapper button {
                height: 100%;
                padding: 0 var(--vspacing-2);
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
</style>