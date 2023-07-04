<script>
    import LiberyHeadline from "@components/UI/LiberyHeadline.svelte";
    import TaggedText from "@components/Wrappers/TaggedText.svelte";
    import ImagePortrait from "@components/UI/ImagePortrait.svelte";
    import LineRect from "@components/UI/LineRect.svelte";
    import TechStack from "./tech_stack.svelte";
    
</script>


<div id="about-me-section-wrapper" aria-hidden="true">
    <article id="about-me-content" class="section-content-layout">
        <div class="section-headline-wrapper">
            <LiberyHeadline headline_text="Me" headline_color="var(--main)"/>
        </div>
        <div id="bio-wrapper">
            <div id="picture-area">
                <ImagePortrait image_src="resources/images/about-me-image.png" alt="Lalo's portrait" />
            </div>
            <div id="text-area" style:position="relative">
                <div id="text-area-anchor-wrapper" >
                    <svg id="anchor-symbol" viewBox="0 0 697 102" fill="none" xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="none">
                        <g id="anchor-element">
                            <path class="anchors-1-cls-2" d="M1.93344 0.195312C1.93344 0.195312 57.6934 97.1411 272.036 61.6883C486.378 26.2355 638.165 43.3881 695.765 84.1149" stroke="#053160" stroke-miterlimit="10"/>
                            <path class="anchors-1-cls-1" d="M278.094 59.2579L273.518 60.0799L274.926 62.752L279.502 61.93L278.094 59.2579Z" fill="#053160"/>
                            <path class="anchors-1-cls-2" d="M181.277 100.034L376.886 19.9971" stroke="#053160" stroke-miterlimit="10"/>
                            <path class="anchors-1-cls-1" d="M181.862 102C183.663 102 185.124 101.147 185.124 100.095C185.124 99.0436 183.663 98.191 181.862 98.191C180.061 98.191 178.601 99.0436 178.601 100.095C178.601 101.147 180.061 102 181.862 102Z" fill="#053160"/>
                            <path class="anchors-1-cls-1" d="M377.68 21.572C379.482 21.572 380.942 20.7193 380.942 19.6675C380.942 18.6157 379.482 17.763 377.68 17.763C375.879 17.763 374.419 18.6157 374.419 19.6675C374.419 20.7193 375.879 21.572 377.68 21.572Z" fill="#053160"/>
                        </g>
                    </svg>    
                </div>
                <LiberyHeadline 
                    vspacing="var(--vspacing-2)" 
                    headline_tag="h2" 
                    headline_text="Hi, I,m Lalo!" 
                    extra_props='role=”<span>passionate-engineer</span>”'
                    text_transform="capitalize"
                />
                <TaggedText tag_name="p">
                    <p class="bio-block">
                        Lorem ipsum dolor sit amet consectetur. Aliquet viverra vitae volutpat non. Tellus amet vitae dolor id. Fames sed eget ac mauris et eget ut. Eget nullam placerat at leo. Faucibus tellus bibendum eu a consequat diam volutpat. Rhoncus blandit massa sed ut pellentesque maecenas. Magna rhoncus in at nulla amet egestas a amet netus. Nisl viverra congue malesuada volutpat eleifend. Non lectus sollicitudin aenean amet molestie. Tincidunt consequat massa purus odio massa sed vitae.<br/>Morbi velit libero eu mauris viverra. Mattis turpis sagittis vitae nullam purus mi amet. Libero 
                    </p>
                </TaggedText>
                <TaggedText tag_name="p">
                    <p class="bio-block">
                        Lorem ipsum dolor sit amet consectetur. Aliquet viverra vitae volutpat non. Tellus amet vitae dolor id. Fames sed eget ac mauris et eget ut. Eget nullam placerat at leo. Faucibus tellus bibendum eu a co Morbi velit libero eu mauris viverra. Mattis turpis sagittis vitae nullam purus mi amet. Libero 
                    </p>
                </TaggedText>
            </div>
        </div>
        <div id="line-rect" style:position="relative">
            <div id="line-rect-wrapper">
                <LineRect/>
            </div>
        </div>
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

        
        /*----------  Anchor  ----------*/
        
            #text-area-anchor-wrapper {
                position: absolute;
                top: -20%;
                left: -23%;
                height: 0px;
            }

            svg#anchor-symbol {
                max-width: 595.609px;
                height: 102px;
            }

            g#anchor-element {
                width: 100%;
            }

            .anchors-1-cls-1 {
                fill: var(--accent-8);
            }
                
            .anchors-1-cls-2 {
                fill: none;
                stroke: var(--accent-8);
                stroke-width: 1.3px;
                stroke-miterlimit: 10;
            }
        
        /*----------  End of Anchor  ----------*/
     
        #text-area {
            grid-column: 6 / span 7;
            display: flex;
            flex-direction: column;
            row-gap: var(--vspacing-3);
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
    
    
</style>