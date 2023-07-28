<script>
    import { createEventDispatcher } from "svelte";
    import { fade } from "svelte/transition";

    export let service_name;
    export let service_index;
    export let is_selected = false;
    export let lines_spacing = .02;

    const select_dispatcher = createEventDispatcher();

    /**
     * @typedef {Object} Line
     * @property {number} x1
     * @property {number} x2
     * @property {number} y1
     * @property {number} y2
     */

    /** @type {Line[]} */
    let lines_data = [];
    const svg_viewbox = { width: 545, height: 125}
    const corners_length = 18.3618;

    lines_data = getLines();
    
    const serviceSelectedHandler = () => {
        select_dispatcher("service_selected", {service_index});
    }

    function getLines() {

        /** @type {Line[]} */
        const lines = [];

        const dx = svg_viewbox.width * lines_spacing;
        const dy = dx;
        let last_x1 = 0;
        let last_y2 = 0;
        let k = 0;
        let new_line = { x1: last_x1, x2: 0, y1: 0, y2: last_y2 };

        while(last_x1 < svg_viewbox.width && last_y2 < svg_viewbox.height) {
            new_line = { x1: 0, x2: 0, y1: 0, y2: 0 };

            new_line.y1 = (k * dy) <= svg_viewbox.height ? k * dy : svg_viewbox.height;

            if (new_line.y1 >= svg_viewbox.height) {
                new_line.x1 = ((k * dy) - svg_viewbox.height) 
                last_x1 += dx;
            }

            new_line.x2 = (dx * k) >= svg_viewbox.width ? svg_viewbox.width : dx * k;

            if (new_line.x2 >= svg_viewbox.width) {
                new_line.y2 = ((dx * k) - svg_viewbox.width);
                last_y2 += dy;
            }
            
            k++;
            lines.push(new_line);
        }

        return lines;
    }
</script>

<button on:click={serviceSelectedHandler} id="ms-service-{service_index}" data-index="{service_index}" class="service-item" class:service-selected={is_selected}>
    <svg style:overflow="visible" class="service-item-corners" viewBox="0 0 {svg_viewbox.width} {svg_viewbox.height}">
        {#if is_selected}
            <g id="bg-lines">
                {#each getLines() as line, h}
                        <!-- <line  class="vl-line" x1="0" x2="{h+h}%" y1="{line}%" y2="0" stroke="var(--grey-8)" stroke-width="2"/>  -->
                        <line in:fade={{duration: 200, delay: .4 + ((13*h) * ((h)/lines_data.length)), opacity: 1}} class="vl-line" x1="{line.x1}" x2="{line.x2}" y1="{line.y1}" y2="{line.y2}" stroke-linecap="round" stroke="var(--grey-8)" stroke-width="2"/> 
                {/each}
            </g>
        {/if}
        <path d="M0 {corners_length}V0H{corners_length}" stroke="#E8E8E8" stroke-width="2"/>
        <path d="M{svg_viewbox.width - corners_length} 0h{corners_length}v{corners_length}" stroke="#E8E8E8" stroke-width="2"/>
        <path d="M{svg_viewbox.width} {svg_viewbox.height - corners_length}v{corners_length}h-{corners_length}" stroke="#E8E8E8" stroke-width="2"/>
        <path d="M{corners_length} {svg_viewbox.height}h-{corners_length}v-{corners_length}" stroke="#E8E8E8" stroke-width="2"/>
    </svg>    
    <h3>
        {service_name.title}
    </h3>
</button>

<style>
        .service-item {
            box-sizing: border-box;
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

        svg.service-item-corners {
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
    
        @media only screen and (max-width: 765px) {

            .service-item {
                /* height: var(--vspacing-4); */
                height: 100%;
            }

            .service-item h3 {
                font-size: calc(.65 * var(--font-size-1));
            }

            .service-selected h3 {
                color: var(--main-dark);
                padding: calc(var(--vspacing-1) / 2);
            }
        }

</style>