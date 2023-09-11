<script>
    import Router from 'svelte-spa-router';
    import { routes } from './routes';
    import NotificationsPopup from '@components/Popups/NotificationsPopup.svelte';
    import Navbar from '@components/Navbar/Navbar.svelte';
    import ContactForm from '@components/Popups/ContactForm.svelte';
    import { onMount } from 'svelte';
    import { defineLayout, layout_properties } from '@stores/layout';
    
    const ENABLE_DEBUG_ON_MOBILE = false;
    if (ENABLE_DEBUG_ON_MOBILE && layout_properties.IS_MOBILE) {
        debugOnMobile();
    }

    onMount(() => {
        defineLayout();

    })
    function debugOnMobile() { 
        let script = document.createElement('script');
        script.src="https://cdn.jsdelivr.net/npm/eruda"; 
        document.body.append(script); 
        script.onload = () => eruda.init();
    }

</script>

<svelte:head>
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
    <link href="https://fonts.googleapis.com/css2?family=La+Belle+Aurore&family=Permanent+Marker&family=Raleway:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;1,100;1,200;1,300;1,400;1,500;1,600&display=swap" rel="stylesheet">
</svelte:head>
<div id="libery-website-wrapper">
    <Navbar />
    <ContactForm />
    <!-- <NotificationsPopup /> -->
    <div id="router-wrapper">
        <Router {routes}/>
    </div>
</div>

<style global>
    :root {
        /* Color palette */
            /* Main */
            --main: hsla(30, 100%, 50%,1);
            --main-9: hsla(29, 100%, 12.5%, 1);
            --main-8: hsla(30, 100%, 24.9%, 1);
            --main-7: hsla(30, 100%, 37.5%, 1);
            --main-6: hsla(30, 100%, 49.8%, 1);
            --main-5: hsla(33, 100%, 60%, 1);
            --main-4: hsla(29, 100%, 70%, 1);
            --main-3: hsla(29, 100%, 70%, 1);
            --main-2: hsla(36, 100%, 80%, 1);
            --main-1: hsla(26, 100%, 88.6%, 1);
            /* Dark */
            --main-dark: hsla(19, 100%, 49%, 1);
            --main-dark-color-9: hsla(19, 100%, 12.4%, 1);
            --main-dark-color-8: hsla(20, 100%, 24.5%, 1);
            --main-dark-color-7: hsla(19, 100%, 36.9%, 1);
            --main-dark-color-6: hsla(23, 100%, 49.8%, 1);
            --main-dark-color-5: hsla(26, 100%, 50%, 1);
            --main-dark-color-4: hsla(26, 100%, 60%, 1);
            --main-dark-color-3: hsla(23, 98.7%, 69.8%, 1);
            --main-dark-color-2: hsla(20, 96.2%, 79.6%, 1);
            --main-dark-color-1: hsla(4, 85.5%, 89.2%, 1);
            --main-dark-transparent: hsla(19, 100%, 49%, .3);
            /* Accent */
            --accent: hsla(211, 100%, 52.5%, 1);
            --accent-9: hsla(212, 88%, 3%, 1);
            --accent-8: hsla(210, 91%, 13.1%, 1);
            --accent-7: hsla(211, 90.1%, 19.8%, 1);
            --accent-6: hsla(211, 89.6%, 26.5%, 1);
            --accent-5: hsla(211, 90.7%, 29.4%, 1);
            --accent-4: hsla(211, 90%, 39.4%, 1);
            --accent-3: hsla(211, 58.7%, 51.6%, 1);
            --accent-2: hsla(211, 100%, 71.6%, 1);
            --accent-1: hsla(211, 100%, 81%, 1);
            /* Grey */
            --grey: hsla(0, 0%, 5.9%, 1);
            --grey-9: hsla(0, 0%, 7.5%, 1);
            --grey-8: hsla(0, 0%, 13.7%, 1);
            --grey-7: hsla(0, 0%, 14.9%, 1);
            --grey-6: hsla(0, 0%, 22.4%, 1);
            --grey-5: hsla(0, 0%, 20%, 1);
            --grey-4: hsla(0, 0%, 27.5%, 1);
            --grey-3: hsla(0, 0%, 54.9%, 1);
            --grey-2: hsla(0, 0%, 72.9%, 1);
            --grey-1: hsla(0, 0%, 91%, 1);


            --grey-t: hsla(0, 0%, 5.9%, .6);
            /* Success */
            --success: hsla(81, 100%, 51.2%, 1);
            --success-9: hsla(91, 100%, 12.5%, 1);
            --success-8: hsla(91, 100%, 25.1%, 1);
            --success-7: hsla(98, 100%, 37.5%, 1);
            --success-6: hsla(104, 100%, 50%, 1);
            --success-5: hsla(107, 100%, 60%, 1);
            --success-4: hsla(107, 100%, 60%, 1);
            --success-3: hsla(120, 100%, 70%, 1);
            --success-2: hsla(132, 100%, 80%, 1);
            --success-1: hsla(80, 100%, 80.4%, 1);
            /* Warning */
            --warning: hsla(49, 100%, 49.6%, 1);
            --warning-9: hsla(52, 100%, 2.9%, 1);
            --warning-8: hsla(50, 100%, 9.2%, 1);
            --warning-7: hsla(49, 100%, 12.4%, 1);
            --warning-6: hsla(49, 100%, 24.9%, 1);
            --warning-5: hsla(49, 100%, 27.8%, 1);
            --warning-4: hsla(49, 100%, 37.1%, 1);
            --warning-3: hsla(49, 59.7%, 49.6%, 1);
            --warning-2: hsla(49, 59.4%, 62.4%, 1);
            --warning-1: hsla(49, 98.1%, 79.8%, 1);
            /* Danger */
            --danger: hsla(3, 100%, 44.9%, 1);
            --danger-9: hsla(3, 100%, 11.2%, 1);
            --danger-8: hsla(4, 100%, 22.9%, 1);
            --danger-7: hsla(3, 100%, 34.5%, 1);
            --danger-6: hsla(10, 100%, 49.6%, 1);
            --danger-5: hsla(10, 99%, 59.8%, 1);
            --danger-4: hsla(10, 98.7%, 69.8%, 1);
            --danger-3: hsla(7, 94.3%, 79.4%, 1);
            --danger-2: hsla(4, 85.5%, 89.2%, 1);
            --danger-1: hsla(4, 85.5%, 89.2%, 1);

            /* Elements */
            --html-tag-color: var(--grey-8);
        /* Spacing */
            --vspacing-base:8px;
            --vspacing-scale: 1;

            --vspacing-1: calc(var(--vspacing-base) * var(--vspacing-scale)); /* 8px */
            --vspacing-2: calc(var(--vspacing-1) * 2); /* 16px */
            --vspacing-3: calc(var(--vspacing-2) * 2); /* 32px */
            --vspacing-4: calc(var(--vspacing-3) * 2); /* 64px */
            --vspacing-5: calc(var(--vspacing-4) * 2); /* 128px */
            --vspacing-6: calc(var(--vspacing-5) * 2); /* 256px */
            --vspacing-7: calc(var(--vspacing-6) * 2); /* 512px */
            --vspacing-8: calc(var(--vspacing-7) * 2); /* 1024px */
            --vspacing-9: calc(var(--vspacing-8) * 2); /* 2048px */

            --hspacing-1: 4.79%;
            --hspacing-2: 36px;

        /* Fonts */
            --font-size-1: var(--vspacing-2); /* 16px */
            --font-size-2: calc(var(--vspacing-3) * 0.59375); /* 19px */
            --font-size-3: calc(var(--vspacing-3) * 0.65625); /* 21px */
            --font-size-4: calc(var(--vspacing-3) * 0.71875); /* 23px */

            --font-size-p: var(--font-size-3);

            --font-size-h1: calc(var(--vspacing-5) * 0.875); /* 112px */
            --font-size-h2: calc(var(--vspacing-5) * 0.4375);
            --font-size-h3: var(--vspacing-3);

            --font-size-CTA-1: calc(var(--vspacing-4) * 0.5625);

            --font-titles: 'Permanent Marker';
            --font-read: 'Raleway';
            --font-decorative: 'La Belle Aurore';

            font-size: var(--font-size-1) !important;

        /* Z-INDEX */

            /* Use these -b- suffix for background elements */
            --z-index-b-1: -1; /* B stands for bottom, not base */
            --z-index-b-2: -2;
            --z-index-b-3: -3;
            --z-index-b-4: -4;
            --z-index-b-5: -5;
            --z-index-b-6: -6;
            --z-index-b-7: -7;
            --z-index-b-8: -8;
            --z-index-b-9: -9;

            /* Regular elements */
            --z-index-1: 1;
            --z-index-2: 2;
            --z-index-3: 3;
            --z-index-4: 4;
            --z-index-5: 5;
            --z-index-6: 6;
            --z-index-7: 7;
            --z-index-8: 8;
            --z-index-9: 9;

            /* Use these -t- suffix for top elements like modals */
            --z-index-t-1: 11; /* T stands for top */
            --z-index-t-2: 12;
            --z-index-t-3: 13;
            --z-index-t-4: 14;
            --z-index-t-5: 15;
            --z-index-t-6: 16;
            --z-index-t-7: 17;
            --z-index-t-8: 18;
            --z-index-t-9: 19;


        /* Layout */

            --desktop-grid: 12;
            --tablet-grid: 8;
            --mobile-grid: 4;

            --default-grid: var(--desktop-grid);
            --full-width: 1 / -1;



        /* Misc */

            --navbar-height: calc(var(--vspacing-5) * 0.96875);
            --border-radius: 4px;
            --border-radius-2: 19px;
            --page-content-width: 90.41%;
            --page-lateral-spacing: calc(calc(100vw - var(--page-content-width)) / 2);
    }

/*=============================================
=            Normalize            =
=============================================*/

    :global(*) {
        box-sizing: border-box;
    }

    :global(body::-webkit-scrollbar) {
        display: none;
    }

    :global(body) {
        -ms-overflow-style: none;  /* IE and Edge */
        scrollbar-width: none;  /* Firefox */
    }

    :global(body) {
        --body-bg-color: var(--grey);
        --body-text-color: var(--grey-1);

        margin: 0;
        padding: 0;
        font-family: var(--font-read);
        font-size: var(--font-size-1);
        color: var(--body-text-color);
        background: var(--body-bg-color);
    }

    :global(h1, h2, h3, h4, h5, h6) {
        font-family: var(--font-titles);
        color: var(--grey-1);
        margin: 0;
    }

    :global(p) {
        margin: 0;
        font-size: var(--font-size-p);
    }


/*=============================================
=            Layout            =
=============================================*/

    #libery-website-wrapper {
        border-top: .1px solid var(--grey);
        position: relative;
    }

    :global(.section-content-layout) {
        display: grid;
        width: var(--page-content-width);
        grid-template-columns: repeat(12, 1fr);
        column-gap: var(--hspacing-2);
    } 

    :global(.background-wrapper) {
        position: absolute;
        width: 100%;
        height: 100%;
        top: 0;
        left: 0;
        z-index: var(--z-index-b-1);
    }

    :global(.full-vw) {
        width: 100vw;
        transform: translateX(calc(var(--page-lateral-spacing) * -1));
    }

    #router-wrapper {
        margin-top: var(--navbar-height);
    }



/*=============================================
=            Text Styles            =
=============================================*/

    :global(.highlight-text) {
        font-family: var(--font-titles);
        font-size: var(--font-size-2);
    }

    :global(.highlight-text--CTA) {
        font-family: var(--font-titles);
        font-size: var(--font-size-CTA-1);
    }

    :global(.sub-headline) {
        font-family: var(--font-read);
        color: var(--main-5);
        font-size: var(--font-size-4);
    }

    :global(.html-tag) {
        color: var(--html-tag-color);
        font-family: var(--font-decorative);
        font-size: var(--font-size-4);
        line-height: .6;
        user-select: none;
    }

    :global(.html-tag > span) {
        color: var(--grey-4);
    }

/*=============================================
=            Buttons            =
=============================================*/
    :global(button) {
        border-radius: var(--border-radius);
    }
    
    :global(.button-1-wrapper) {
        width: max-content;
        clip-path: polygon(0 0,0 0,100% 0,100% 0,100% calc(100% - 15px),calc(100% - 15px) 100%,15px 100%,0 100%);
        background: var(--success-3);
        padding: 1px;
        border-radius: var(--border-radius);
    }

    @media(pointer: fine) {
        :global(.button-1-wrapper) {
            transition: all .4s cubic-bezier(0.075, 0.82, 0.165, 1);
        }

        :global(.button-1-wrapper:hover) {
            background: var(--success-6) ;
            filter: brightness(1.3) saturate(1.5) hue-rotate(2deg);
        }
    }

    :global(.button-1) {
        border: none;
        font-family: var(--font-titles);
        background: transparent;
        font-size: var(--font-size-1);
        color: var(--grey-7);
        padding: 12px 32px;
        clip-path: polygon(0 0,0 0,100% 0,100% 0,100% calc(100% - 15px),calc(100% - 15px) 100%,15px 100%,0 100%);
    }


/*=============================================
=            Debug            =
=============================================*/

    :global(.debug) {
        border: 1px solid red;
    }
    
    :global(.adebug) {
        border: 1px solid rgb(255, 72, 72);
    }

    :global(.debug > *) {
        border: 1px solid red;
    }

    :global(.adebug *) {
        border: 1px solid rgb(107, 107, 255);
    }

    :global(.dtwo) {
        border: 1px solid blue !important;
    }

    :global(.dthree) {
        border: 1px solid green !important;
    }

    :global(.dfour) {
        border: 1px solid yellow !important;
    }



/*=====  End of Debug  ======*/


/*=============================================
=            Animations            =
=============================================*/

    @keyframes floating {
        0% { transform: translate(0,  0px); }
        50%  { transform: translate(0, 15px); }
        100%   { transform: translate(0, -0px); }   
    }

    :global(.anim-floating) {
        animation: floating 3s ease-in-out infinite;
    }

    @keyframes shake-horizontal {
        0%,
        100% {
            transform: translateX(0);
        }
        10%,
        30%,
        50%,
        70% {
            transform: translateX(-10px) scaleX(1);
        }
        20%,
        40%,
        60% {
            transform: translateX(10px) scaleX(1.03);
        }
        80% {
            transform: translateX(8px);
        }
        90% {
            transform: translateX(-8px);
        }
    }

    :global(.anim-shake-h) {
        animation: shake-horizontal 1s ease-in-out infinite;
    }

    @keyframes shake-vertical {
        0%,
        100% {
            transform: translateY(0);
        }
        10%,
        30%,
        50%,
        70% {
            transform: translateY(-8px);
        }
        20%,
        40%,
        60% {
            transform: translateY(8px);
        }
        80% {
            transform: translateY(6.4px);
        }
        90% {
            transform: translateY(-6.4px);
        }
    }

    :global(.anim-shake-v) {
        animation: shake-vertical 1s ease-in-out infinite;
    }

/*=====  End of Animations  ======*/
/*=============================================
=            Mobile            =
=============================================*/

    :global(.hide-on-mobile) {
        display: none !important;
    }

    :global(.hide) {
        display: none !important;
    }

    @media only screen and (max-width: 765px) {
        :root {
            --vspacing-scale: .714583;

            --font-size-h1: calc(var(--vspacing-5) * 0.575); /* 112px */
            --font-size-h2: calc(var(--vspacing-5) * 0.3375);

            --default-grid: var(--mobile-grid);
            --html-tag-color: var(--grey-6);
        }

        :global(.section-content-layout) {
            display: flex;
            width: var(--page-content-width);
            flex-direction: column;
            column-gap: var(--vspacing-3);
        } 

        :global(.html-tag) {
            font-size: var(--font-size-h3);
            line-height: .6;
            user-select: none;
        }
    }

/*=====  End of Mobile  ======*/





</style>

