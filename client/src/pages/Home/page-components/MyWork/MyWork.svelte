<script>
    import LiberyHeadline from "@components/UI/LiberyHeadline.svelte";
    import ProjectsCarousel from "./sub-components/ProjectsCarousel.svelte";
    import ProjectPreviewer from "./sub-components/ProjectPreviewer.svelte";
</script>

<section id="my-work-section" class="section-content-layout" style:display="flex">
    <div class="section-headline-wrapper" style:grid-column="1 / -1">
        <LiberyHeadline headline_text="Recent work" headline_color="var(--main)" text_transform="capitalize"/>
    </div>
    <div id="projects-carousel-component" class="full-vw">
        <ProjectsCarousel />
    </div> 
    <div id="project-preview-component" class="full-vw">
        <ProjectPreviewer />
    </div>
</section>

<style>
    section#my-work-section {
        flex-direction: column;
        row-gap: var(--vspacing-5);
    }


    #projects-carousel-component {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }
</style>