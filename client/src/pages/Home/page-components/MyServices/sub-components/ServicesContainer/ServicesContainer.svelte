<script>
    import DotGridBg from "@components/UI/DotGridBG.svelte";
    import { fade, fly } from "svelte/transition";

    // @typedef {Object} Service
    // @property {string} title
    // @property {string} description
    
    // @type {Service[]}
    export let services;
    export let service_selected = 0;
    export let user_clicked = false;
    
    /*=============================================
    =            Lines animation            =
    =============================================*/
    
        const num_lines = 110;
        
        const getLines = () => {
            const lines = [];
            
    
            for (let i = 0; i < num_lines; i += 1) {
                let p = i*19.2;
                lines.push(p);
            }
    
            return lines;
        }
    
    /*=====  End of Lines animation  ======*/
    
    const onServiceClick = index => {
        service_selected = index;
        user_clicked = true;
    }
    
</script>

<section id="services-selector">
    <header id="services-container-header">
        <h2>
            Services
        </h2>
    </header>
    <div id="services-container">
        {#each services as service, h}
            <button on:click={() => onServiceClick(h)} id="ms-service-{h}" data-index="{h}" class="service-item {h === service_selected ? 'service-selected' : ''}">
                <svg class="service-item-corners" viewBox="0 0 545 125">
                    {#if h === service_selected}
                        <g id="bg-lines">
                            {#each getLines() as line, h}
                                    <line in:fade={{duration: 200, delay: .4*(h*(num_lines-h)/2), opacity: 1}} class="vl-line" x1="0" x2="{h+h}%" y1="{line}%" y2="0" stroke="var(--grey-8)" stroke-width="2"/> 
                            {/each}
                        </g>
                    {/if}
                    <path d="M1.31537 18.3618V1.36182H15.824" stroke="#E8E8E8" stroke-width="2"/>
                    <path d="M1.31537 106.362V123.362H15.824" stroke="#E8E8E8" stroke-width="2"/>
                    <path d="M543.315 18.3618V1.36182H528.807" stroke="#E8E8E8" stroke-width="2"/>
                    <path d="M543.315 106.362V123.362H528.807" stroke="#E8E8E8" stroke-width="2"/>
                </svg>    
                <h3>
                    {service.title}
                </h3>
            </button>
        {/each}
    </div>
    <div id="dot-grid-wrapper">
        <DotGridBg />
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

    #dot-grid-wrapper {
        position: absolute;
        top: 0;
        left: -7%;
        width: 100vw;
        height: 100%;
        overflow: hidden;
        z-index: var(--z-index-b-2);
    }

    header#services-container-header {
        grid-column: var(--full-width);
    }

    header#services-container-header h2 {
        font-size: calc(var(--font-size-h3) * 1.21);
        margin: calc(var(--vspacing-2) * .7) 0;
        text-align: center;
    }

    #services-container {
        grid-column: var(--full-width);
        display: grid;
        grid-template-columns: repeat(3, calc(var(--vspacing-7) + calc(var(--vspacing-7) * .058)));
        grid-auto-rows: var(--vspacing-5);
        column-gap: var(--vspacing-3);
        row-gap: calc(var(--vspacing-3) + 4px);
    }

    .service-item {
        background: none;
        position: relative;
        display: flex;
        justify-content: center;
        align-items: center;
        border: none;
        height: var(--vspacing-5);
        z-index: var(--z-index-1);
        transition: all .2s ease-in-out;
    }

    .service-item h3 {
        width: max-content;
        font-size: var(--font-size-3);
        font-family: var(--font-read);
        font-weight: 400;
        color: var(--grey-1);
        text-align: center;
    }

    .service-item-corners {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        z-index: var(--z-index-b-1);
    }

    .service-item-corners path {
        fill: none;
    }

    @media(pointer: fine) {
        .service-item:not(.service-selected):hover {
            background: var(--main-dark-transparent);
        }

        .service-item:not(.service-selected):hover h3 {
            color: var(--grey-1);
            filter: brightness(1.6);
        }
    }

    .service-selected h3 {
        color: var(--main-dark);
        background: var(--grey);
        padding: var(--vspacing-1);
    }

</style>