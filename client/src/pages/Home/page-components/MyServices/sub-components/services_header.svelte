<script>
    import LiberyBorderStar from "@components/UI/LiberyBorderStar.svelte";
    import LiberyHeadline from "@components/UI/LiberyHeadline.svelte";
    import { fade, fly } from "svelte/transition";
    import { bounceIn } from "svelte/easing";

    const getLines = () => {
        const lines = [];
        const num_lines = 200;

        for (let i = 0; i < num_lines; i++) {
            let p = i*19.2;
            lines.push(p);
        }

        return lines;
    }
</script>

<header class="services-header">
    <div id="white-star-wrapper">
        <LiberyBorderStar />
    </div>
    <LiberyHeadline headline_text="what can i craft for you?" headline_color="var(--grey-1)" text_transform="lowercase" />
    <svg id="vertical-lines" viewBox="0 0 1920 100" style:width="110%">
        {#each getLines() as line, h}
             <line in:fly={{y:10*h, duration: 400, delay: 30*h, opacity: 1}} class="vl-line" x1="0" x2="{h}%" y1="{line}%" y2="0" stroke="var(--grey-8)" stroke-width=".2%"/> 
        {/each}
    </svg>
</header>    
    
<style>
    .services-header {
        grid-column: 1 / -1;
        display: flex;
        align-items: center;
        flex-direction: column;
    }

    #white-star-wrapper {
        width: 12.67%;
    }

    :global(#white-star-wrapper svg) {
        width: 100%;
    }
</style>