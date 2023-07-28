<script>
    import { layout_properties } from "@stores/layout";
    import LiberyHeadline from "@components/UI/LiberyHeadline.svelte";
    import BlueBg from "./sub-components/BlueBg.svelte";
</script>

<section id="cms-wrapper" style:position="relative" class:debug={false}>
    <div id="cms-libery-star-wrapper">
        <svg viewBox="0 0 557 484" fill="none" xmlns="http://www.w3.org/2000/svg">
            <g clip-path="url(#clip0_44_12)">
                <path d="M283.276 239.984L351.908 139.797L361.099 318.849L283.276 239.984Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M202.127 326.7L276.919 239.983H283.276L361.099 318.849L361.181 320.481L288.044 308.766L279.295 308.288L202.127 326.7Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M206.895 136.21L276.919 239.984L202.127 326.701L204.732 264.322L206.895 136.21Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M206.895 136.21L277.181 177.436L283.276 177.175L351.908 139.797L283.276 239.984H276.919L206.895 136.21Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M121.224 165.628L144.03 197.372L275.313 72.5229L121.224 165.628Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M286.716 72.5229L417.263 188.52L442.264 158.208L286.716 72.5229Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M206.895 136.21L205.617 213.813L129.416 211.673L206.895 136.21Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M278.754 72.5229H275.313L206.895 136.21L258.536 167.445L278.754 72.5229Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M278.754 72.5229H281.048L304.755 165.967L283.276 177.174L277.181 177.436L258.536 167.445L278.754 72.5229Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M358.871 136.056L304.755 165.967L281.048 72.5229H286.717L358.871 136.056Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M205.617 213.813L129.416 211.673L127.007 213.813L204.732 264.322L205.617 213.813Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M202.127 326.7L127.007 213.813L204.732 264.322L202.127 326.7Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M127.007 213.813L77.2004 261.382L190.986 335.675L257.815 374.546L202.127 326.7L127.007 213.813Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M202.127 326.7L279.295 308.289L281.326 384.429L282.473 386.862L252.392 369.882L202.127 326.7Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M279.295 308.289L281.326 384.429L282.473 386.862L284.865 384.429L288.044 308.766L279.295 308.289Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M288.044 308.766L361.181 320.481L362.049 324.453L357.265 333.427L338.178 351.377L282.473 386.862L284.865 384.429L288.044 308.766Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M355.316 206.347H439.315L358.871 136.056L351.908 139.796L355.316 206.347Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M355.315 206.347H439.315L358.052 259.673L355.315 206.347Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M358.052 259.673L361.181 320.481L439.315 206.347L358.052 259.673Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M439.315 206.347L486.173 251.191L365.228 336.414L338.178 351.378L357.265 333.428L362.049 324.453L361.181 320.481L439.315 206.347Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M282.473 386.862L338.178 351.377L365.228 336.414L376.713 328.332L436.038 447.808V453.489L282.473 386.862Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M446.409 279.208L436.071 447.423L376.745 327.932L446.409 279.208Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M282.473 386.862L186.038 332.442L128.121 462.017L282.473 386.862Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M186.038 332.443L122.436 290.924L128.121 462.017L186.038 332.443Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M214.071 420.175L282.473 482.954L348.287 415.418L282.473 386.862L214.071 420.175Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M97.8932 285.582L58.1133 336.414L124.91 365.294L124.025 338.739L97.8932 285.582Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M441.806 354.01L511.043 325.438L467.807 279.609L463.891 284.089L443.378 328.948L441.806 354.01Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M121.224 165.628L127.007 5.1416L179.812 130.283L121.224 165.628Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M179.812 130.283L229.93 100.048L127.007 5.1416L179.812 130.283Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M127.007 5.1416L275.313 72.523L229.93 100.048L127.007 5.1416Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M286.716 72.523L427.289 1.49316L334.082 98.7089L286.716 72.523Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M334.082 98.7089L427.289 1.49316L381.824 124.972L334.082 98.7089Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M381.824 124.972L442.264 158.208L427.289 1.49316L381.824 124.972Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M242.054 57.4057L275.313 18.4424L315.88 57.7751L286.716 72.523H275.313L242.054 57.4057Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M275.313 72.523V18.4424L286.716 72.523H275.313Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M122.059 142.46L121.224 165.629L144.03 197.372L123.812 216.877L0.819214 182.408L122.059 142.46Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M123.812 216.876L77.2005 261.382L0.819214 182.408L123.812 216.876Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M417.263 188.52L442.264 158.208L440.151 136.056L556.181 176.435L439.315 206.347L417.263 188.52Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M439.315 206.347L486.173 251.191L556.181 176.436L439.315 206.347Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M486.173 251.191L446.409 279.209L443.378 328.948L463.891 284.089L467.807 279.609L486.173 251.191Z" stroke="#7F3F00" stroke-linejoin="round"/>
                <path d="M124.025 338.739L97.8932 285.582L77.2004 261.382L122.436 290.924L124.025 338.739Z" stroke="#7F3F00" stroke-linejoin="round"/>
            </g>
            <defs>
                <clipPath id="clip0_44_12">
                <rect width="557" height="483" fill="white" transform="translate(0 0.723633)"/> 
                </clipPath>
            </defs>
        </svg>
    </div>
    <article id="cms-contact-cta">
        <div id="cms-cc-top-content">
            <LiberyHeadline  
                headline_color="var(--grey-1)"
                headline_text="Like what you see?"
                text_transform="none"
                forced_font_size={layout_properties.IS_MOBILE ? "calc(var(--font-size-h2) * 1.2)" : "calc(var(--font-size-h1)*1)"}
            />
        </div>
        <div id="cms-cc-bottom-content">
            <div id="cms-cc-bc-cta" class:hide={layout_properties.IS_MOBILE}>
                <LiberyHeadline headline_tag="h2" headline_text="Let's Talk" text_transform="none"/>
            </div>
            <div id="cms-cc-bc-cta-btn" class="button-1-wrapper">
                <button class="button-1">Contact me!</button>
            </div>
        </div>
    </article>
</section>

<style>

    #cms-libery-star-wrapper {
        position: absolute;
        top: 62%;
        left: 99.4%;
        width: calc(var(--vspacing-7) * 1.08789);
        transform: translate(-50%, -50%);
        z-index: var(--z-index-b-2);
    }

    #cms-contact-cta {
        display: flex;
        flex-direction: column;
        gap: var(--vspacing-2);
    }

    #cms-cc-bottom-content {
        display: flex;
        justify-content: center;
        align-items: center;
        gap: var(--vspacing-4);
        
    }

    #cms-cc-bc-cta-btn button {
        font-size: calc(var(--font-size-h2) * .86);
        line-height: 1;
    }

    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 765px) {
            #cms-libery-star-wrapper {
                top: 37%;
                left: 65.3%;
                width: calc(var(--vspacing-6) * 1.33789);
                transform: translate(-50%, -50%);
                z-index: var(--z-index-b-2);
            }

            #cms-contact-cta {
                align-items: center;
                gap: var(--vspacing-3);
            }

            #cms-cc-bc-cta-btn button {
                font-size: calc(var(--font-size-h3) * .86);
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
</style>