<script>
    import LineSymbolsBg from "@components/UI/LineSymbolsBG.svelte";
    import ProjectContent from "./ProjectContent.svelte";
    import ProjectScreenshots from "./ProjectScreenshots.svelte";
    import { Project } from "@models/Project";
    import { layout_properties } from "@stores/layout";

    /** 
     * @type {Project}
    */
    export let project;
</script>

<section id="libery-project-preview" class="" class:debug={false} style:position="relative">
    <div class="background-wrapper">
        <LineSymbolsBg
            line_count={15}
            x_base={layout_properties.IS_MOBILE ? 450 : 1300}
            y_base={layout_properties.IS_MOBILE ? 340 : 120}
            svg_width={layout_properties.IS_MOBILE ? innerWidth + 120 : innerWidth}
            limit_height={!layout_properties.IS_MOBILE}
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

    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 765px) {
            #lpp-project-content {
                flex-direction: column-reverse;
                justify-content: flex-end;
            }
            
        }
    
    /*=====  End of Mobile  ======*/
    
    

</style>