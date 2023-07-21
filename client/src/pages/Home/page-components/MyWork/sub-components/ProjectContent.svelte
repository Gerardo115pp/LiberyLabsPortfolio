<script>
    import LiberyBorderStar from "@components/UI/LiberyBorderStar.svelte";
    import LiberyHeadline from "@components/UI/LiberyHeadline.svelte";
    import TaggedText from "@components/Wrappers/TaggedText.svelte";
    import { Project } from "@models/Project";

    /** 
     * @type {Project}
    */
    export let project;

</script>


<!-- HTMl -->
    <header id="project-name-wrapper" style:position="relative">
        <LiberyHeadline headline_color="var(--grey-1)" headline_text="{project.name}" text_transform="capitalize" headline_tag="h2"/>
        <div class="pnw-libery-star-wrapper background-wrapper">
            <LiberyBorderStar stroke_color="var(--danger-9)" fill_color="none" stroke_width="2px"/>
        </div>
    </header>
    <div id="project-description-wrapper">
        <TaggedText tag_name="p">
            <p class="project-description">
                {project.description}
            </p>
        </TaggedText>
    </div>
    <div id="bottom-content-wrapper">
        <div id="bcw-teck-stack-wrapper">
            <TaggedText tag_name="stack">
                <ul id="tech-stack">
                    <div class="tech-stack-item">Svelte</div>
                    <div class="tech-stack-item">Nginx</div>
                    <div class="tech-stack-item">CSS</div>
                </ul>
            </TaggedText>
        </div>
        <div id="bcw-cta">
            <div class="button-1-wrapper">
                <button class="button-1">Contact me!</button>
            </div>
        </div>
    </div>
<!-- END OF HTML -->

<style>
    #project-name-wrapper {
        width: max-content;
    }

    .pnw-libery-star-wrapper {
        width: calc(var(--vspacing-6) * .7981);
        top: -40%;
        left: 75%;
    }

    #bottom-content-wrapper {
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    #tech-stack {
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: flex-start;
        align-items: center;
        row-gap: var(--vspacing-2);
        column-gap: var(--vspacing-2);
        list-style: none;
        padding: 0;
        margin: 0;
    }

    .tech-stack-item {
        font-family: var(--font-decorative);
        color: var(--main);
        font-size: var(--font-size-4);
    }
</style>