<script>
    import ContactMe from "../ContactMe/ContactMe.svelte";
    import ChatBot from "../ChatBot/ChatBot.svelte";
    import Footer from "@components/Footer/Footer.svelte";
    import BlueBg from "../ContactMe/sub-components/BlueBg.svelte";
    import { layout_properties } from "@stores/layout.js"


</script>


<div id="bottom-content"  class:debug={false}>
    <div id="bc-wrapper" style:position="relative">
        <div id="blue-bg-wrapper">
            <BlueBg 
                match_viewport={true}
            />
        </div>
        <div id="bc-content">
            <div id="bc-relative-wrapper" style:position="relative">
                <section id="contact-me-section" class="page-section" >
                    <div class="section-content">
                        <ContactMe />
                    </div>
                </section>
                <section id="chat-bot-section" class="page-section" >
                    <div class="section-content">
                        <ChatBot />
                    </div>
                </section>
            </div>
        </div>
    </div>
    <div id="home-footer-wrapper">
        <Footer/>
    </div>
</div>

<style>
    #bottom-content {
        height: max-content;
        width: 100%;
    }

    #bc-content {
        position: absolute;
        width: 100%;
        height: 100%;
        top: 0;
        left: 0;
    }

    #bc-relative-wrapper {
        height: 100%;
        width: 100%;
    }

    #bc-relative-wrapper .page-section {
        width: max-content;
        position: absolute;
    }

    #bc-relative-wrapper #contact-me-section {
        top: 8.726%;
        left: 19.021%;
    }

    #bc-relative-wrapper #chat-bot-section {
        width: 100%;
        top: 52.726%;
    }
</style>