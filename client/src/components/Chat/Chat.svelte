<script>
    import { layout_properties } from "@stores/layout";
    import { token as chat_token, is_available, messages } from "@stores/chat";
    import { ChatDialer } from "@models/Chat";
    import ChatMessages from "./sub-components/ChatMessages.svelte";
    import DottedStar from "./sub-components/DottedStar.svelte";
    import { onMount, afterUpdate } from "svelte";
    import ChatWriter from "./sub-components/ChatWriter.svelte";

    /** @type {ChatDialer}*/
    let chat_dialer;
    let is_awaiting_response = false;
    
    onMount(() => {
        // chat_dialer = new ChatDialer(); // To styleize the chat component enable this line
    });

    const handleChatFocus = () => {
        if (chat_dialer === undefined) {
            chat_dialer = new ChatDialer();
        }
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
        <DottedStar />
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
            <ChatMessages />
        {/if}
    </div>
    <ChatWriter {chat_dialer} on:chatFocus={handleChatFocus} />
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
        background: var(--main-dark);
        align-items: center;
        padding: 0 var(--vspacing-3);
        z-index: var(--z-index-2);
    }

    #sb-controls {
        height: 34%;
    }
    
    /*=============================================
    =            Chat content            =
    =============================================*/
    
    
    
        #chat-content {
            --connection-status-bar-height: 50px;
            height: var(--chat-content-height);
            padding: var(--vspacing-1);
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

    /*=====  End of Chat content  ======*/



    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 768px) {
            #chat-content {
                opacity: 0;
            }

            #chat-component {
                height: calc(var(--vspacing-8) * 0.455859375);
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
</style>