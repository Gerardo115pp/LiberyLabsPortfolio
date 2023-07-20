<script>
    import libery_labs_textlogo from "@svg/main_icon_1.svg";
    import TaggedText from "@components/Wrappers/TaggedText.svelte";
    import LiberyFloatingStar from "@components/UI/LiberyFloatingStar.svelte";
    import LineSymbolsBg from "@components/UI/LineSymbolsBG.svelte";
    
</script>

<article id="hero-content" class="section-content-layout"  style:position="relative" style:padding-top="var(--vspacing-4)">
    <div class="background-wrapper full-vw">
        <LineSymbolsBg
            line_count={8}
            step={0.026}
            x_base={1400}
            y_base={310}
            svg_width={innerWidth + 120}
        />
    </div>
    <div id="libery-hero-info">
        <div id="lhi-top">
            <div id="lhi-headline-wrapper">
                <div id="headline-logo">
                    {@html libery_labs_textlogo}    
                </div>
                <h2 class="sub-headline">Where I Cast Your Dreams into Digital Reality</h2>
            </div>
            <div id="libery-description-wrapper">
                <TaggedText tag_name="p">
                    <p id="libery-description">
                        In Libery Labs, your wildest ideas find form in cutting-edge software.<br/> We craft bespoke digital solutions that deliver on innovation,<br/> functionality, and scalability.
                    </p>
                </TaggedText>
            </div>
        </div>
        <div id="cta-wrapper">
            <div class="button-1-wrapper">
                <button class="button-1">Check my work</button>
            </div>
        </div>
    </div>
    <div id="hero-star-wrapper">
        <LiberyFloatingStar/>
    </div>
</article>

<style>

    /* * {
        border: 1px solid red;
    } */

    #hero-content {
        height: calc(100vh - var(--navbar-height));
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
    
    
</style>
    
    
