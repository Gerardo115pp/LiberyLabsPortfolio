<script>
    import { ProjectImages } from "@models/ProjectImage";
    import { Project } from "@models/Project";
    import { onMount } from "svelte";
    import { layout_properties } from "@stores/layout";

    /**
     * @type {Project}
     */
    export let project;

    /**
     * @type {Object.<string, ProjectImages>}
     */
    const loaded_project_images = {};

    /**
     * @type {ProjectImages}
     */
    let project_images;
    let has_images = false;

    /**
     * @type {SVGImageElement} 
    */
    let desktop_image_display;

    /**
     * @type {SVGImageElement} 
    */
    let mobile_image_display;

    /** @type {boolean} - wether the project has both mobile and desktop screenshots or just one of the two.*/
    let has_desktop_mobile = false;

    $: if (mobile_image_display !== undefined) {
        // images have loaded at least once
        setProjectImages(), project;
    }

    const svg_sizes = getDisplaySizes();

    onMount(() => {
        // setProjectImages();
    });

    const advanceMobileImage = (backwards=false) => {
        if (!has_images) return;

        console.log(`backwards: ${backwards}`)  

        if (backwards) {
            console.log('popo')
            mobile_image_display.href.baseVal = project_images.PreviousMobileImage;
            return;
        }

        mobile_image_display.href.baseVal = project_images.NextMobileImage;
    }

    const advanceDesktopImage = (backwards) => {
        if (!has_images) return;

        if (backwards) {
            desktop_image_display.href.baseVal = project_images.PreviousDesktopImage;
            return;
        }

        desktop_image_display.href.baseVal = project_images.NextDesktopImage;
    }

    const setProjectImages = async () => {
        if( project_images !== undefined && project_images === loaded_project_images[project.ID]) return;

        if (loaded_project_images[project.ID] === undefined) {
            loaded_project_images[project.ID] = new ProjectImages(project.ID);
            await loaded_project_images[project.ID].loadProjectImages();
        }

        project_images = loaded_project_images[project.ID];
        project_images.resetImageCounter();


        has_images = project_images.TotalImages > 0;
        setImagesHref();
    }

    const setImagesHref = () => {
        if (!has_images) return;
        let has_desktop = false;
        let has_mobile = false;

        const mobile_display_wrapper = document.getElementById('mobile-view-wrapper');
        const desktop_display_wrapper = document.getElementById('desktop-view-wrapper');

        if (mobile_display_wrapper === null || desktop_display_wrapper === null) return;
        
        if (desktop_image_display !== undefined && project_images.CurrentDesktopImage !== null) {
            desktop_display_wrapper.style.visibility = 'visible';
            desktop_image_display.href.baseVal = project_images.CurrentDesktopImage;
            has_desktop = true;
        } else {
            desktop_display_wrapper.style.visibility = 'hidden';
        }

        if (mobile_image_display !== undefined && project_images.CurrentMobileImage !== null) {
            mobile_display_wrapper.style.visibility = 'visible';
            mobile_image_display.href.baseVal = project_images.CurrentMobileImage;
            has_mobile = true;
        } else {
            mobile_display_wrapper.style.visibility = 'hidden';
        }

        has_desktop_mobile = has_desktop && has_mobile;
    }

    function getDisplaySizes() {
        const display_sizes = {
            desktop_display: {
                w: 74.358,
                h: 49.864
            },
            mobile_display: {
                w: 45.441,
                h: 87.907
            }
        }

        if (display_sizes.desktop_display.w >= $layout_properties.VIEWPORT_WIDTH || display_sizes.desktop_display.h >= $layout_properties.VIEWPORT_HEIGHT) {
            const wh_ratio = display_sizes.desktop_display.w / display_sizes.desktop_display.h;
            display_sizes.desktop_display.w = $layout_properties.VIEWPORT_WIDTH * 0.7;
            display_sizes.desktop_display.h = display_sizes.desktop_display.w / wh_ratio;
        }

        console.log(`mobile h: ${display_sizes.mobile_display.h} | desktop h: ${display_sizes.desktop_display.h}`)

        if (display_sizes.mobile_display.h > (3 * display_sizes.desktop_display.h) || display_sizes.mobile_display.w >= $layout_properties.VIEWPORT_WIDTH || display_sizes.mobile_display.h >= $layout_properties.VIEWPORT_HEIGHT) {
            console.log('mobile display too big')
            const wh_ratio = display_sizes.mobile_display.w / display_sizes.mobile_display.h;
            display_sizes.mobile_display.w = $layout_properties.VIEWPORT_WIDTH * 0.5;
            display_sizes.mobile_display.h = display_sizes.mobile_display.w / wh_ratio;
        }

        console.log(display_sizes)

        return display_sizes;
    }
</script>

<div id="screenshot-displayer"  class:adebug={false}>
    <div id="desktop-view-wrapper" class:single-mode-display={!has_desktop_mobile}>
        <svg 
            style:width="{has_desktop_mobile ? svg_sizes.desktop_display.w : (svg_sizes.desktop_display.w * 1.3)}cqw"
            style:height="{has_desktop_mobile ? svg_sizes.desktop_display.h : (svg_sizes.desktop_display.h * 1.3)}cqh"
            viewBox="0 0 522 367"
            fill="none"
        >
            <path id="desktop-bg" d="M20 3L502 3L502 264L20 264Z"/>
            <image bind:this={desktop_image_display} x="10" id="desktop-image-display" width="500" height="267"/>
            <path class="device-body" d="M334.154 366.729H188.545C188.545 366.729 187.752 355.193 213.102 344.472C238.442 333.761 232.9 282.653 232.9 282.653H289.809C289.809 282.653 284.258 333.761 309.607 344.472C334.947 355.193 334.154 366.729 334.154 366.729Z"/>
            <path class="device-body" d="M511.492 0.729004H11.2077C5.51417 0.729004 0.845276 5.16188 0.845276 10.5778V279.398C0.845276 284.813 5.51417 289.246 11.2077 289.246H511.483C517.185 289.246 521.845 284.813 521.845 279.398V10.5778C521.845 5.16188 517.176 0.729004 511.483 0.729004H511.492ZM504.382 247.616C504.382 253.737 499.57 258.745 493.689 258.745H29.6338C23.7531 258.745 18.9417 253.737 18.9417 247.616V32.3156C18.9417 26.1949 23.7531 21.187 29.6338 21.187H493.689C499.57 21.187 504.382 26.1949 504.382 32.3156V247.616Z"/>
            <path tabindex="0" role="button" on:click={() => advanceDesktopImage(false)} class="device-button next-btn action-btn" d="M254.204 268.167L262.57 283.255L270.946 268.167H254.204Z"/>
        </svg>    
    </div>
    <div id="mobile-view-wrapper" class:single-mode-display={!has_desktop_mobile}>
        <svg style:width="{svg_sizes.mobile_display.w}cqw" style:height="{svg_sizes.mobile_display.h}cqh" viewBox="0 0 319 647" fill="none">
            <path id="cellphone-bg" d="M10 60Q0 0 50 10H269Q319 0 309 60V587Q319 647 269 637H50Q0 647 10 587Z"></path>
            <image bind:this={mobile_image_display} x="10" y="10" id="mobile-image-display" width="300" height="627"/>
            <path class="device-body" d="M283.293 0.729004H40.1346C40.1346 0.729004 8.27126 4.57248 2.65857 37.3335V61.7365L3.53011 608.19C3.53011 608.19 3.53011 641.308 41.0062 646.537H282.805C282.805 646.537 315.54 643.923 318.155 609.061V37.8477C318.155 37.8477 317.283 8.57283 283.293 0.729004ZM274.142 632.157H45.7647C45.7647 632.157 22.2681 627.381 17.9105 604.494V36.0262C17.9105 36.0262 17.9105 19.467 37.0842 15.1094H79.6936C79.6936 15.1094 84.1472 14.2378 84.1472 19.467C84.1472 24.6962 85.0187 28.1824 85.0187 28.1824C85.0187 28.1824 88.5049 36.8978 97.2202 37.7693H223.305C223.305 37.7693 233.18 36.0262 234.923 28.1824C236.666 20.3386 235.794 19.467 235.794 18.5955C235.794 17.724 236.666 15.1094 239.281 15.1094C241.895 15.1094 283.346 15.0571 283.346 15.0571C283.346 15.0571 301.16 16.8524 302.903 35.1547C304.646 53.4569 302.903 602.525 302.903 602.525C302.903 602.525 304.646 623.442 274.142 632.157Z"/>
            <path class="device-button" d="M5.27316 87.8827C5.27316 86.4387 4.10256 85.2681 2.65855 85.2681C1.21455 85.2681 0.0439453 86.4387 0.0439453 87.8827V109.671C0.0439453 111.115 1.21455 112.286 2.65855 112.286C4.10256 112.286 5.27316 111.115 5.27316 109.671V87.8827Z"/>
            <path tabindex="0" role="button" on:click={() => advanceMobileImage(false)} class="device-button next-btn action-btn" d="M5.27316 134.074C5.27316 132.63 4.10256 131.459 2.65855 131.459C1.21455 131.459 0.0439453 132.63 0.0439453 134.074V178.522C0.0439453 179.966 1.21455 181.137 2.65855 181.137C4.10256 181.137 5.27316 179.966 5.27316 178.522V134.074Z"/>
            <path tabindex="0" role="button" on:click={() => advanceMobileImage(true)} class="device-button action-btn" d="M5.27316 195.081C5.27316 193.637 4.10256 192.467 2.65855 192.467C1.21455 192.467 0.0439453 193.637 0.0439453 195.081V237.787C0.0439453 239.231 1.21455 240.401 2.65855 240.401C4.10256 240.401 5.27316 239.231 5.27316 237.787V195.081Z"/>
            <path class="device-button" d="M128.404 30.0822C132.019 30.0822 134.949 27.1518 134.949 23.5369C134.949 19.9221 132.019 16.9917 128.404 16.9917C124.789 16.9917 121.859 19.9221 121.859 23.5369C121.859 27.1518 124.789 30.0822 128.404 30.0822Z"/>
            <path class="device-button" d="M104.001 30.1872C107.568 30.1872 110.459 27.2958 110.459 23.7291C110.459 20.1624 107.568 17.271 104.001 17.271C100.434 17.271 97.5427 20.1624 97.5427 23.7291C97.5427 27.2958 100.434 30.1872 104.001 30.1872Z"/>
            <path class="device-button" d="M189.734 27.5724C191.77 27.5724 193.42 25.9219 193.42 23.8858C193.42 21.8498 191.77 20.1992 189.734 20.1992C187.698 20.1992 186.047 21.8498 186.047 23.8858C186.047 25.9219 187.698 27.5724 189.734 27.5724Z"/>
            <path class="device-button" d="M214.651 29.8559C218.001 29.8559 220.717 27.1401 220.717 23.79C220.717 20.4399 218.001 17.7241 214.651 17.7241C211.301 17.7241 208.585 20.4399 208.585 23.79C208.585 27.1401 211.301 29.8559 214.651 29.8559Z"/>
            <path class="device-button" d="M174.787 22.5176H146.026C144.823 22.5176 143.847 23.4931 143.847 24.6964C143.847 25.8998 144.823 26.8753 146.026 26.8753H174.787C175.99 26.8753 176.966 25.8998 176.966 24.6964C176.966 23.4931 175.99 22.5176 174.787 22.5176Z"/>
        </svg>    
    </div>
</div>

<style>
    @keyframes blink-buttons-color {
        0% {
            fill: var(--main-5);
            scale: 1;
        }

        3% {
            fill: var(--main-dark-color-2);
            scale: 1.07;
        }

        6% {
            fill: var(--main-5);
            scale: 1;
        }

        9% {
            fill: var(--main-dark-color-2);
            scale: 1.07;
        }

        12% {
            fill: var(--main-5);
            scale: 1;
        }

        15% {
            fill: var(--main-dark-color-2);
            scale: 1.07;
        }

        18% {
            fill: var(--main-5);
            scale: 1;
        }
    }

    #screenshot-displayer {
        container-type: size;
        position: relative;
        height: 100%;
    }

    #mobile-view-wrapper {
        position: absolute;
        top: 0;
        left: 65%;
        width: max-content;
        height: max-content;
        z-index: var(--z-index-1);
    }

    #mobile-view-wrapper.single-mode-display {
        left: -5%;
        transform-origin: center;
        transform: translateX( 50%);
    }

    #mobile-view-wrapper svg {
        max-width: 319px;
        max-height: 647px;
    }

    #desktop-view-wrapper {
        position: absolute;
        bottom:  5%;
        left: 0%;
        width: max-content;
        height: max-content;
        z-index: var(--z-index-2);
    }

    #desktop-view-wrapper.single-mode-display {
        left: -25%;
        bottom: 50%;
        transform-origin: center;
        transform: translate(50%, 50%);
    }

    #desktop-view-wrapper svg {
        max-width: 522px;
        max-height: 367px;
    }

    #mobile-image-display {
        clip-path: inset(0 0 0 0 round 40px 40px 40px 40px);
    }

    path#cellphone-bg, path#desktop-bg {
        fill: var(--grey-9);
        stroke: none;
    }

    .device-body {
        fill: var(--grey-9);
        stroke: var(--grey-8);
        stroke-width: 1.2px;
        filter: brightness(1.04);
    }

    .device-button {
        fill: var(--main-dark-color-5);
    }

    @media (pointer: fine) {
        .action-btn {
            fill: var(--main-5);
            transform-box: fill-box;
            transform-origin: center;
            transition: all .2s ease-in-out;
            animation-name: blink-buttons-color;
            animation-duration: 6.5s;
            animation-iteration-count: infinite;
            animation-direction: alternate;
            animation-delay: 1s;
        }

        .action-btn:hover {
            fill: var(--main-dark-color-4);
            scale: 1.1;
            animation: none;
        }
    }

    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
        @media only screen and (max-width: 765px){
            #screenshot-displayer {
                height: 60vh;
            }

            #mobile-view-wrapper {
                left: 45%;
                top: 0%;
            }

            #desktop-view-wrapper.single-mode-display {
                left: 0%;
                transform: translate(0%, 50%);   
            }

            #mobile-view-wrapper.single-mode-display {
                left: 50%;
                transform: translateX(-50%);
            }
        }
    
    /*=====  End of Mobile  ======*/
    
    
</style>