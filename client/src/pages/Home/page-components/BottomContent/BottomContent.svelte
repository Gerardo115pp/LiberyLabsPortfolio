<script>
    import ContactMe from "../ContactMe/ContactMe.svelte";
    import ChatBot from "../ChatBot/ChatBot.svelte";
    import Footer from "@components/Footer/Footer.svelte";

    let bottom_content_height = 0;
    let chat_section_content_height = 0;
    let contact_section_content_height = 0;
    const chat_top_offset = 0.3;

    $: bottom_content_height = chat_section_content_height + contact_section_content_height;
</script>

<div id="stupid-wrapper" style:position="relative">
    <div id="bottom-content" style:height="{(bottom_content_height) - (bottom_content_height * chat_top_offset) }px" style:position="relative" class:debug={false}>
        <section bind:clientHeight={contact_section_content_height} id="contact-me-section" class="page-section">
            <div class="section-content">
                <ContactMe />
            </div>
        </section>
        <section bind:clientHeight={chat_section_content_height} style:top="{chat_top_offset*100}%" id="chat-bot-section" class="page-section ">
            <div class="section-content">
                <ChatBot />
            </div>
        </section>
    </div>
    <Footer/>
</div>

<style>

    #stupid-wrapper {
        max-height: 220vh;
        overflow: hidden;
    }

    .page-section {
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 100%;
        overflow: hidden;
    }

    .section-content {
        display: flex;
        width: 100%;
        justify-content: center;
    }

    #bottom-content {
        height: max-content;
        width: 100%;
    }

    #chat-bot-section {
        position: absolute;
        top: 30%;
        width: 100%;
        z-index: var(--z-index-2);
    }
</style>