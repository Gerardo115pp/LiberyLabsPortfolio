<script>
    import viewport from "@components/viewport_actions/useViewportActions";
    import LiberyHeadline from "@components/UI/LiberyHeadline.svelte";
    import Bracket from "@components/UI/Bracket.svelte";
    import { onMount, onDestroy } from "svelte";
    import { SECTIONS as HOME_SECTIONS } from "@pages/Home/sections";

    const tech_stack = [
        "svelte",
        "nginx",
        "kubernetes",
        "react",
        "C#",
        "unity",
        "go",
        "python",
        "C",
        "GCP",
        "selenuim",
        "web assembly",
        "javascript",
        "docker",
        "HTML",
        "tensorflow",
        "pandas",
        "flask",
        "CSS",
        "api design",
        "mysql",
        "bind9",
        "prometheus",
        "grafana",
        "git",
        "linux",
        "redis",
        "prefect",
        "stripe",
        "nasm",
        "gsap",
        "three.js",
        "ebpf",
        "systemd",
        "OpenAI",
        "Discord Bots"
    ].slice(0, 36);

    let random_tech_items = [];

    
    /*----------  Animation settings  ----------*/
    const animation_round_length = 4500; // in ms
    const items_transition_duration = animation_round_length * .55; 
    let animation_enabled = false;

    let animator_timer;
    

    onMount(() => {
        random_tech_items = selectRandomTechItems(tech_stack);

        animator_timer = window.setInterval(() => {
            if (!animation_enabled) return;

            random_tech_items = selectRandomTechItems(tech_stack);

        }, animation_round_length);
    })

    onDestroy(() => {
        window.clearInterval(animator_timer);
    })

    /**
     * @description Selects a item from each of the 6 columns of the tech stack. Mainly used for animations.
     * @param {string[]} tech_stack
     * @returns {int[]} Array of 6 random numbers in ranges 0-5, 6-11, 12-17, 18-23, 24-29, 30-35
    */
    const selectRandomTechItems = tech_stack => {

        const columns = tech_stack.length / 6;

        /** @type {int[]} */
        const random_teck_items = [];
        const occupied_positions = new Set();

        for (let h = 0; h < columns; h++) {
            let random_number = Math.floor(Math.random() * 6) + (h * 6);
            let position = random_number%columns;

            if (occupied_positions.has(position)) {
                random_number = h * columns; // sets random_number%columns to be 0
                random_number -= 1; // so the loop sets it back to 0
                do {
                    random_number++;
                    position = random_number%columns;
                } while (occupied_positions.has(position) && random_number < ((h + 1) * columns));
            }

            random_teck_items.push(random_number);
            occupied_positions.add(position);
        }

        return random_teck_items;
    }

</script>

<article id="tech-stack-content" class="section-content-layout">
    <div class="section-headline-wrapper-reverse">
        <LiberyHeadline headline_text="My Arsenal" headline_color="var(--main)" text_transform="capitalize"/>
    </div>
    <div id="tsc-bottom-content" on:viewportEnter={() => animation_enabled = true} on:viewportLeave={() => animation_enabled = false} use:viewport>
        <div class="bracket-wrapper">
            <Bracket/>
        </div>
        <ul data-scroll-section={HOME_SECTIONS.ARSENAL} id="teck-stack-container">
            {#each tech_stack as ti, k}
                <li 
                    data-item-index={k}
                    class="teck-stack-item"
                    class:ts-item-selected={random_tech_items.includes(k)}
                    style:animation-duration="{animation_round_length* .87}ms"
                    style:animation-delay={random_tech_items.includes(k) ? `${Math.random() * (animation_round_length * .13)}ms` : `${Math.random() * (animation_round_length * .73)}ms`} 
                >{ti}</li>
            {/each}
        </ul>
        <div class="bracket-wrapper">
            <Bracket/>
        </div>
    </div>
</article>

<style>

    /* #tech-stack-content * {
        border: 1px solid red;
    } */

    #tech-stack-content {
        column-gap: none;
        row-gap: var(--vspacing-4);
    }

    .section-headline-wrapper-reverse {
        grid-column: 1 / span 12;
        display: flex;
        justify-content: flex-end;
    }

    #tsc-bottom-content {
        grid-column: 1 / span 12;
        display: flex;
        justify-content: center;
    }

    .bracket-wrapper {
        grid-column: 1 / span 2;
        height: 100%;
    }

    .bracket-wrapper:last-of-type {
        grid-column: 11 / span 2;
        transform: rotate(180deg);
    }

    :global(.bracket-wrapper svg) {
        fill: var(--accent-4);
        opacity: 0.1;
    }

    #teck-stack-container {
        width: min(80%, 1348px);
        display: grid;
        grid-template-columns: repeat(6, minmax(160px, 1fr));
        row-gap: var(--vspacing-2);
        /* column-gap: var(--hspacing-2); */
        align-content: start;
        list-style: none;
        margin: 0;
        padding: 0;
        outline: none;
        /* padding-inline-start: 0 !important; */
    }

    @keyframes LightUp {
        0% {
            color: var(--accent-8);
        }
        10% {
            color: var(--accent-7);
        }
        22% {
            color: var(--accent-3);
            filter: blur(.6px);
        }
        70% {
            color: var(--accent-4);
            filter: blur(0px);
        }
        100% {
            color: var(--accent-8);
        }
    }

    @keyframes LightGently {
        0% {
            filter: brightness(1.2);
        }
        30% {
            filter: brightness(1.22);
        }
        45% {
            filter: brightness(1.4);
        }
        57% {
            filter: brightness(1.22);
        }
        65% {
            filter: brightness(1.2);
        }
        78% {
            filter: brightness(1.22);
        }
        100% {
            filter: brightness(1.2);
        }
    }

    .teck-stack-item {
        white-space: nowrap;
        /* font-family: var(--font-titles); */
        /* font-family: var(--font-read); */
        font-family: var(--font-decorative);
        color: var(--accent-8);
        font-size: var(--font-size-h3);
        text-transform: capitalize;
        line-height: 1;
        font-weight: lighter;
        text-align: center;
        animation-name: LightGently;
        animation-iteration-count: infinite;
    }

    :global(.teck-stack-item.ts-item-selected) {
        animation-name: LightUp;
        animation-iteration-count: 1;
        animation-fill-mode: forwards;
        /* font-weight: bold; */
    }

    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
    @media (max-width: 768px) {
        #tsc-bottom-content {
            flex-direction: column;
            align-items: center;
            /* row-gap: var(--vspacing-4); */
        }

        .bracket-wrapper {
            display: flex;
            width: 100%;
            height: 22vw;
            justify-content: center;
            align-items: flex-start;
        }
        
        :global(.bracket-wrapper svg) {
            transform: translateY(-30%) rotate(90deg);
            transform-origin: center;
            height: 100vw !important;
            width: max-content;
            opacity: .1;

        }

        #teck-stack-container {
            grid-template-columns: repeat(3, minmax(80px, 1fr));
            width: 100%;
        }

        .teck-stack-item {
            font-size: var(--font-size-3);
        }
    }
    
    /*=====  End of Mobile  ======*/
    
    

</style>