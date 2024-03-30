<script>
    import Hero from "@pages/Home/page-components/Hero/Hero.svelte";
    import { onMount } from "svelte";
    import { recaptcha_pk } from "@stores/env";
    import AboutMe from "@pages/Home/page-components/AboutMe/AboutMe.svelte";
    import MyServices from "@pages/Home/page-components/MyServices/MyServices.svelte";
    import NoBoundaries from "@pages/Home/page-components/NoBoundaries/NoBoundaries.svelte";
    import MyWork from "@pages/Home/page-components/MyWork/MyWork.svelte";
    import BottomContent from "@pages/Home/page-components/BottomContent/BottomContent.svelte";
    import { writable } from "svelte/store";


    
    /*=============================================
    =            Properties            =
    =============================================*/
    
        let has_recaptcha_loaded = writable(false);
    
    /*=====  End of Properties  ======*/
    
    

    onMount(() => {
        loadRecaptcha();
    });
    
    /*=============================================
    =            Methods            =
    =============================================*/
    
        function loadRecaptcha() {
            let script = document.createElement('script');
            script.src=`https://www.google.com/recaptcha/api.js?render=${recaptcha_pk}`; 
            document.body.append(script); 

            script.onload = () => {
                console.log('Recaptcha loaded');
                console.log(grecaptcha);
                has_recaptcha_loaded.set(true);
            }
        }
    
    /*=====  End of Methods  ======*/
    
    

</script>

<main id="home-page">
    <section id="hero-section" class="page-section">
        <div class="section-content">
            <Hero />
        </div>
    </section>
    <section id="about-me-section" class="page-section">
        <div class="section-content" >
            <AboutMe />
        </div>
    </section>
    <section id="tech-stack-section" class="page-section">
        <div class="section-content">
            <MyServices />
        </div>
    </section>
    <section id="broke-text-section" class="page-section" style:overflow="visible">
        <div class="section-content">
            <NoBoundaries />
        </div>
    </section>
    <section id="recent-work-section" class="page-section">
        <div class="section-content">
            <MyWork />
        </div>
    </section>
    <!--vvv Contact, Chat section and Footer vvv-->
    <BottomContent /> 
</main>

<style>
    /* * {
        border: 1px solid red;
    } */

    #home-page {
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 100%;
        row-gap: var(--vspacing-6);
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


</style>