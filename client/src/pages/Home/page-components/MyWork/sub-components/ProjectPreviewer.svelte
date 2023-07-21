<script>
    import LiberyBorderStar from "@components/UI/LiberyBorderStar.svelte";
    import LiberyHeadline from "@components/UI/LiberyHeadline.svelte";
    import LineSymbolsBg from "@components/UI/LineSymbolsBG.svelte";
    import TaggedText from "@components/Wrappers/TaggedText.svelte";
    import ProjectContent from "./ProjectContent.svelte";
    import ProjectScreenshots from "./ProjectScreenshots.svelte";
    import { Project } from "@models/Project";

    /** 
     * @type {Project}
    */
    export let project;
</script>

<section id="libery-project-preview" class="" class:debug={false} style:position="relative">
    <div class="background-wrapper">
        <LineSymbolsBg
            line_count={15}
            x_base={1300}
            y_base={120}
            step={0.026}
        />
    </div>
    <div id="lpp-project-content" class="section-content-layout" class:debug={false}>
        <div id="llp-pc-left">
            <ProjectContent {project}/>
        </div>
        <div id="llp-pc-right">
            <ProjectScreenshots {project} />
        </div>
    </div>
</section>

<style>
    #libery-project-preview {
        display: flex;
        justify-content: center;
        align-items: flex-start;
    }

    #lpp-project-content {
        min-height: inherit;
        grid-template-rows: 100%;
    }

    #llp-pc-left {
        display: flex;
        flex-direction: column;
        grid-column: 1 / span 6;
        row-gap: var(--vspacing-3);
    }

    #llp-pc-right {
        grid-column: 7 / span 5;
        height: 100%;
    }

</style>