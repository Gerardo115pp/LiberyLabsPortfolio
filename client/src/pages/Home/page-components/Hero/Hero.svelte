<script>
    import libery_labs_textlogo from "@svg/main_icon_1.svg";
    import TaggedText from "@components/Wrappers/TaggedText.svelte";
    import LiberyFloatingStar from "@components/UI/LiberyFloatingStar.svelte";
    import LineSymbolsBg from "@components/UI/LineSymbolsBG.svelte";
    import { layout_properties } from "@stores/layout";
    import { writable } from "svelte/store";
    import MainLogo from "@components/UI/MainLogo.svelte";
    import { SECTIONS as HOME_SECTIONS, scrollToSection } from "@pages/Home/sections";

    const hero_content_height = writable(0);
    
    /**
     * @type {HTMLDivElement}
     */
    let hero_star_element;
    hero_content_height.subscribe(updateStarHeight);
    
    function updateStarHeight(new_height){
        if (hero_star_element === undefined) return; 
        
        window.hero_star_element = hero_star_element;
        

        hero_star_element.style.height = `${new_height}px`;
    }

</script>

<article data-scroll-section="{HOME_SECTIONS.HERO}" id="hero-content" class="section-content-layout"  style:position="relative" class:debug={false}>
    <div class="background-wrapper full-vw">
        <LineSymbolsBg
            line_count={layout_properties.IS_MOBILE ? 23 : 10}
            step={ layout_properties.IS_MOBILE ? 0.028 : 0.026}
            x_base={layout_properties.IS_MOBILE ? 50 : 1400}
            y_base={layout_properties.IS_MOBILE ? 680 : 310}
            svg_width={layout_properties.IS_MOBILE ? innerWidth : innerWidth + 120}
        />
    </div>
    <div id="libery-hero-info">
        <div id="lhi-top">
            <div id="lhi-headline-wrapper">
                <MainLogo
                    fill_color="var(--main)"
                    stroke_color="var(--main)"
                    subheadline_color="var(--main-5)"
                /> 
            </div>
            <div id="libery-description-wrapper">
                <TaggedText tag_name="p">
                    <p id="libery-description">
                        In Libery Labs, your wildest ideas find form in cutting-edge software.<br/> We craft bespoke digital solutions that deliver on innovation,<br/> functionality, and scalability.
                    </p>
                </TaggedText>
            </div>
        </div>
        <div id="cta-wrapper" class:hide-on-mobile={layout_properties.IS_MOBILE}>
            <div class="button-1-wrapper">
                <button on:click={() => scrollToSection(HOME_SECTIONS.PORTFOLIO)} class="button-1">Check my work</button>
            </div>
        </div>
    </div>
    <div bind:this={hero_star_element} id="hero-star-wrapper">
        <LiberyFloatingStar height_store={hero_content_height}/>
    </div>
</article>

<style>

    /* * {
        border: 1px solid red;
    } */

    #hero-content {
        height: calc(100vh - var(--navbar-height));
        padding-top: var(--vspacing-4);
    }
    
    /*=============================================
    =            Hero Info            =
    =============================================*/
        
        #libery-hero-info {
            grid-column: 1 / span 5;
            width: 100%;
            display: flex;
            flex-direction: column;
            row-gap: var(--vspacing-3);
        }

        #lhi-top {
            display: flex;
            flex-direction: column;
            row-gap: var(--vspacing-2);
        }

        #lhi-headline-wrapper {
            display: flex;
            flex-direction: column;
            gap: var(--vspacing-1);
        }

        :global(#headline-logo svg) {
            width: 387px
        }

        #cta-wrapper{
            display: flex;
            width: 100%;
            justify-content: flex-end;
        }

    /*=====  End of Hero Info  ======*/

    
    /*=============================================
    =            Libery Star            =
    =============================================*/
    
        #hero-star-wrapper {
            grid-column: 6 / span 7;
            /* background: blue; */
        }
    
    /*=====  End of Libery Star  ======*/
    
    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 765px) {
            #hero-content {
                flex-direction: column-reverse;
                justify-content: flex-end;
                row-gap: var(--vspacing-4);
                padding-top: var(--vspacing-3);
            }

            #libery-hero-info {
                grid-column: 1 / span 4; /* section-content-layout not sets a 4 column grid as material design suggests */
            }

            #hero-star-wrapper {
                display: flex;
                justify-content: center;
            }

            #lhi-top {
                row-gap: var(--vspacing-3);
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    

    
</style>
    
    
