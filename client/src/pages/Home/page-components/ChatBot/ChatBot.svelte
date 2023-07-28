<script>
    import LiberyHeadline from "@components/UI/LiberyHeadline.svelte";
    import OrangeBg from "./sub-components/OrangeBG.svelte";
    import Chat from "./sub-components/Chat.svelte";
</script>

<section id="ask-chat-section" class:debug={false}>
    <div id="acs-section-content" class="section-content-layout" >
        <div id="acs-information">
            <header id="acs-info-header">
                <LiberyHeadline headline_text="Ask Away" forced_font_size="var(--font-size-h2)" headline_color="var(--grey-1)"/>
            </header>
            <p id="acs-subheadline">
                Clear Doubts, Discover Possibilities with Valaria powered by ChatGPT.
            </p>
        </div>
        <div id="acs-chat">
            <Chat />
        </div>
    </div>
</section>

<style>
    #ask-chat-section {
        width: 100%;
        display: flex;
        justify-content: center;
    }

    #acs-information {
        grid-column: 1 / span 4;
        display: flex;
        flex-direction: column;
        gap: var(--vspacing-3);
    }

    #acs-subheadline {
        color: var(--success-3);
    }

    #acs-chat {
        grid-column:6 / -1;
    }

</style>