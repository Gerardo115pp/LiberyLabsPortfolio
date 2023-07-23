<script>
    import Hero from "./page-components/Hero/Hero.svelte";
    import AboutMe from "./page-components/AboutMe/AboutMe.svelte";
    import MyServices from "./page-components/MyServices/MyServices.svelte";
    import NoBoundaries from "./page-components/NoBoundaries/NoBoundaries.svelte";
    import MyWork from "./page-components/MyWork/MyWork.svelte";
    import BottomContent from "./page-components/BottomContent/BottomContent.svelte";
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
    <section id="broke-text-section" class="page-section">
        <div class="section-content">
            <NoBoundaries />
        </div>
    </section>
    <section id="recent-work-section" class="page-section">
        <div class="section-content">
            <MyWork />
        </div>
    </section>
    <BottomContent /> <!-- Contact, Chat section and Footer -->
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
        overflow-x: hidden;
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