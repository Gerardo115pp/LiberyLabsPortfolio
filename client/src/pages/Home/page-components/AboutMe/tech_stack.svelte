<script>
    import LiberyHeadline from "@components/UI/LiberyHeadline.svelte";
    import Bracket from "@components/UI/Bracket.svelte";

    const teck_stack = [
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

</script>

<article id="tech-stack-content" class="section-content-layout">
    <div class="section-headline-wrapper-reverse">
        <LiberyHeadline headline_text="My Arsenal" headline_color="var(--main)" text_transform="capitalize"/>
    </div>
    <div id="tsc-bottom-content" >
        <div class="bracket-wrapper">
            <Bracket/>
        </div>
        <ul id="teck-stack-container">
            {#each teck_stack as ti}
                <li class="teck-stack-item">{ti}</li>
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

    .teck-stack-item {
        white-space: nowrap;
        font-family: var(--font-titles);
        color: var(--accent-8);
        font-size: var(--font-size-h3);
        text-transform: capitalize;
        line-height: 1;
        font-weight: lighter;
        text-align: center;
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