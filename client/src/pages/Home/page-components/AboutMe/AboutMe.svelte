<script>
    import LiberyHeadline from "@components/UI/LiberyHeadline.svelte";
    import TaggedText from "@components/Wrappers/TaggedText.svelte";
    import ImagePortrait from "@components/UI/ImagePortrait.svelte";
    import LineRect from "@components/UI/LineRect.svelte";
    import TechStack from "./tech_stack.svelte";
    import AnchorsOne from "@components/UI/AnchorsOne.svelte";
    import { layout_properties } from "@stores/layout";
    
</script>


<div id="about-me-section-wrapper" aria-hidden="true">
    <article id="about-me-content" class="section-content-layout">
        <div class="section-headline-wrapper">
            <LiberyHeadline headline_text="Myself" text_transform="none" headline_color="var(--main)"/>
        </div>
        <div id="bio-wrapper">
            <div id="picture-area" >
                <ImagePortrait image_src="resources/images/about-me-image.png" alt="Lalo's portrait" />
            </div>
            <div id="text-area" style:position="relative">
                <div id="text-area-anchor-wrapper" >
                    <AnchorsOne/>
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
                        Yo soy lalo y soy bien chido. Lorem ipsum dolor sit amet consectetur. Aliquet viverra vitae volutpat non. Tellus amet vitae dolor id. Fames sed eget ac mauris et eget ut. Eget nullam placerat at leo. Faucibus tellus bibendum eu a consequat diam volutpat. Rhoncus blandit massa sed ut pellentesque maecenas. Magna rhoncus in at nulla amet egestas a amet netus. Nisl viverra congue malesuada volutpat eleifend. Non lectus sollicitudin aenean amet molestie. Tincidunt consequat massa purus odio massa sed vitae.<br/>Morbi velit libero eu mauris viverra. Mattis turpis sagittis vitae nullam purus mi amet. Libero 
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

            #text-area-anchor-wrapper {
                position: static;
                width: max-content;
                height: max-content
                /* transform: translateX(-55%); */
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