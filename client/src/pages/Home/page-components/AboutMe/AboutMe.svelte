<script>
    import LiberyHeadline from "@components/UI/LiberyHeadline.svelte";
    import TaggedText from "@components/Wrappers/TaggedText.svelte";
    import ImagePortrait from "@components/UI/ImagePortrait.svelte";
    import LineRect from "@components/UI/LineRect.svelte";
    import TechStack from "./tech_stack.svelte";
    import AnchorsOne from "@components/UI/AnchorsOne.svelte";
    import { layout_properties } from "@stores/layout";
    import { SECTIONS as HOME_SECTIONS } from "@pages/Home/sections";
    
</script>


<div id="about-me-section-wrapper" aria-hidden="true">
    <article id="about-me-content" class="section-content-layout">
        <div class="section-headline-wrapper">
            <LiberyHeadline headline_text="Myself" text_transform="none" headline_color="var(--main)"/>
        </div>
        <div data-scroll-section={HOME_SECTIONS.ABOUT} id="bio-wrapper">
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
                        Welcome! My name is Lalo, I'm a fullstack web engineer based in Jalisco, Mexico. I'm passionate about turning ideas into reality and taking businesses to the next level. I understand front-end, backend and I'm in love with Linux, Docker and K8S, i know exactly what how each of these technologies a must be combined to give <strong>your business the scalability, reliability and efficiency it needs to grow and succeed</strong>.
                    </p>
                </TaggedText>
                <TaggedText tag_name="p">
                    <p class="bio-block">
                        Ever since i was a child i always loved building things, and getting new tools on my toolkit to achieve bigger and better things. Aside from coding, my other passions are carpentry, wood carving, 3D modeling, printing and painting miniatures. For me the only thing better than building things is helping others through something i built.
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