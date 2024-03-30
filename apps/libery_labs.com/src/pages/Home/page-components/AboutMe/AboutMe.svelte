<script>
    import viewport from "@components/viewport_actions/useViewportActions";
    import LiberyHeadline from "@components/UI/LiberyHeadline.svelte";
    import { SECTIONS as HOME_SECTIONS } from "@pages/Home/sections";
    import ImagePortrait from "@components/UI/ImagePortrait.svelte";
    import LineRect from "@components/UI/LineRect.svelte";
    import { layout_properties } from "@stores/layout";
    import TechStack from "./tech_stack.svelte";
    import AInfoArea from "./AInfoArea.svelte";
    import { fly } from "svelte/transition";
    
    let is_component_visible = false;
    $: console.log(is_component_visible);

    let animation_delay_increment = 300;
    let animation_duration = 800;

</script>


<div id="about-me-section-wrapper" aria-hidden="true">
    <article id="about-me-content" class="section-content-layout">
        <div class="section-headline-wrapper" on:viewportEnter={() => is_component_visible = true} use:viewport>
            <LiberyHeadline animated={false} headline_text="Myself" text_transform="none" headline_color="var(--main)"/>
        </div>
        <div data-scroll-section={HOME_SECTIONS.ABOUT} id="bio-wrapper">
            <div id="picture-area" >
                <ImagePortrait image_src="resources/images/about-me-image-original.webp" alt="Lalo's portrait" />
            </div>
            <AInfoArea visible={is_component_visible} {animation_delay_increment}/>
        </div>
        {#if is_component_visible}
            <div in:fly={{x: 1500, duration: animation_duration, delay: animation_delay_increment*4}} id="line-rect" style:position="relative">
                <div id="line-rect-wrapper">
                    <LineRect/>
                </div>
            </div>
        {/if}
    </article>
    <TechStack/>
</div>

<style>

    
    /*=============================================
    =            About me content            =
    =============================================*/
    
        #about-me-section-wrapper {
            display: flex;
            width: 100%;
            flex-direction: column;
            gap: var(--vspacing-5);
            align-items: center;
        }

        #about-me-section-wrapper > * {
            width: var(--page-content-width);
        }

        #about-me-content {
            row-gap: var(--vspacing-4);
        }

        .section-headline-wrapper {
            grid-column: 1 / span 12;
        }

        #bio-wrapper {
            display: grid;
            grid-column: 1 / span 12;
            grid-template-columns: repeat(12, 1fr);
        }

        #picture-area {
            grid-column: 1 / span 5;
            height: 100%;
            display: flex;
            justify-content: center;
        }

        

    /*=====  End of About me content  ======*/

    



    
    /*=============================================
    =            Line rect            =
    =============================================*/
    
        #line-rect {
            /* width: 1101px; */
            height: 140.19px;
            grid-column: 5 / span 8;
        }
        
        #line-rect-wrapper {
            position: absolute;
            width: 100%;
            height: 100%;
            left: 9.53%;
        }
    
    
    /*=====  End of Line rect  ======*/

    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 768px) {
            #about-me-content {
                row-gap: calc(1.4 * var(--vspacing-4));
            }

            #bio-wrapper {
                display: flex;
                flex-direction: column;
                row-gap: var(--vspacing-3);
            }

            #line-rect {
                width: 100vw;
                height: auto;
            }
            
            #line-rect-wrapper {
                position: static;
                width: 170vw;
                height: auto;
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
    
    
</style>