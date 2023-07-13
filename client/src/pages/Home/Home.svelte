<script>
    import Hero from "./page-components/Hero/Hero.svelte";
    import AboutMe from "./page-components/AboutMe/AboutMe.svelte";
    import LineSymbolsBg from "@components/UI/LineSymbolsBG.svelte";
    import HrDivider from "@components/UI/HrDivider.svelte";
    import MyServices from "./page-components/MyServices/MyServices.svelte";
    import NoBoundaries from "./page-components/NoBoundaries/NoBoundaries.svelte";
</script>

<main id="home-page">
    <section id="hero-section" class="page-section">
        <div class="section-content" style:position="relative">
            <Hero />
            <LineSymbolsBg />
        </div>
        <HrDivider />
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