<script>
    import DotGridBg from "@components/UI/DotGridBG.svelte";
    import { layout_properties } from "@stores/layout";
    import { fade, fly } from "svelte/transition";
    import ServiceWrapper from "./ServiceWrapper.svelte";

    // @typedef {Object} Service
    // @property {string} title
    // @property {string} description
    
    // @type {Service[]}
    export let services;
    export let service_selected = 0;
    export let user_clicked = false;
    
    const onServiceClick = index => {
        service_selected = index;
        user_clicked = true;
    }
    
</script>

<section id="services-selector" class:debug={false}>
    <header id="services-container-header">
        <h2>
            Services
        </h2>
    </header>
    <div id="services-container" >
        {#each services as service, h}
            <ServiceWrapper on:service_selected={e => onServiceClick(e.detail.service_index)} is_selected={h === service_selected} service_name={service} service_index={h}/>
        {/each}
    </div>
    <div id="dot-grid-wrapper">
        <DotGridBg
            spacing={ layout_properties.IS_MOBILE ? 20 : 24 }
        />
    </div>
</section>

<style>
    #services-selector {
        position: relative;
        display: grid;
        grid-column: 1 / -1;
        grid-template-columns: repeat(var(--default-grid), 1fr);
        justify-items: center;
        row-gap: var(--vspacing-2);
        padding: var(--vspacing-2) 0 calc(var(--vspacing-3)) 0;
    }

    
    /*=============================================
    =            dot grid            =
    =============================================*/
    
        #dot-grid-wrapper {
            position: absolute;
            top: 0;
            left: -7%;
            width: 100vw;
            height: 100%;
            overflow: hidden;
            z-index: var(--z-index-b-2);
        }    
    
    /*=====  End of dot grid  ======*/
    /*=============================================
    =            Header            =
    =============================================*/

        header#services-container-header {
            grid-column: var(--full-width);
        }

        header#services-container-header h2 {
            font-size: calc(var(--font-size-h3) * 1.21);
            margin: calc(var(--vspacing-2) * .7) 0;
            text-align: center;
        }
    
    /*=====  End of Header  ======*/
    
    /*=============================================
    =            Services            =
    =============================================*/
    
        #services-container {
            grid-column: var(--full-width);
            display: grid;
            grid-template-columns: repeat(3, calc(var(--vspacing-7) + calc(var(--vspacing-7) * .058)));
            grid-auto-rows: var(--vspacing-5);
            column-gap: var(--vspacing-3);
            row-gap: calc(var(--vspacing-3) + 4px);
        }

    /*=====  End of Services  ======*/

    
    


    @media only screen and (max-width: 765px) {
        #services-container {
            grid-template-columns: repeat(2, 1fr);
            row-gap: 0;
            column-gap: var(--vspacing-2);
            grid-auto-rows: var(--vspacing-4);
        }

    }

</style>