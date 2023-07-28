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
                <ul id="tech-stack" class:debug={false}>
                    {#each project.tech_stack as technology, h}
                        <div class="tech-stack-item">{technology + (h+1 === project.tech_stack.length ? '' : ',')}</div>
                    {/each}
                </ul>
            </TaggedText>
        </div>
        <div id="previewer-controls">
            {#if project.github !== ''}
                <div class="button-1-wrapper redirect">
                    <a href="{project.github}" target="_blank" rel="noopener noreferrer">
                        <button class="button-1">
                            GITHUB
                        </button>
                    </a>
                </div>
            {/if}
            {#if project.url !== ''}
                <div class="button-1-wrapper redirect">
                    <a href="{project.url}"  target="_blank" rel="noopener noreferrer">
                        <button class="button-1">
                            See for yourself!
                        </button>
                    </a>
                </div>
            {/if}
            <div id="bcw-cta" class="button-1-wrapper">
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
        width: calc(var(--vspacing-6) * 1.1);
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: flex-start;
        align-items: center;
        
        column-gap: var(--vspacing-1);
        list-style: none;
        padding: 0;
        margin: 0;
    }

    .tech-stack-item {
        font-family: var(--font-decorative);
        color: var(--main);
        font-size: var(--font-size-4);
        text-transform: capitalize;
    }

    #previewer-controls {
        display: flex;
        justify-content: center;
        align-items: center;
        column-gap: var(--vspacing-1);
    }

    /* #previewer-controls button {
    } */
    
    .button-1-wrapper.redirect {
        background: var(--grey-7);
    }
    
    .button-1-wrapper.redirect button.button-1 {
        font-family: var(--font-titles);
        font-size: calc(var(--font-size-1) * .8) !important;
        background: transparent;
        color: var(--grey-1);
    }

    @media(pointer: fine) {
        .button-1-wrapper.redirect:hover {
            background: var(--grey-5);
        }
    }

    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 765px) {
            #bottom-content-wrapper {
                flex-direction: column;
                row-gap: var(--vspacing-3);
            }

            #tech-stack {
                width: 100%;
                justify-content: center;
            }

            #previewer-controls {
                width: 100%;
                gap: var(--vspacing-2);
            }

            #previewer-controls .button-1-wrapper {
                display: flex;
                justify-content: center;
            }

            #previewer-controls .button-1-wrapper .button-1 {
                padding: var(--vspacing-2) var(--vspacing-3);
                white-space: nowrap;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    

</style>