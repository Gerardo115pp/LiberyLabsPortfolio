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
</article>

<style>

    /* #tech-stack-content * {
        border: 1px solid red;
    } */

    #tech-stack-content {
        row-gap: var(--vspacing-4);
        border-top: 1px solid transparent; /* margin collapse fix*/
    }

    .section-headline-wrapper-reverse {
        grid-column: 1 / span 12;
        display: flex;
        justify-content: flex-end;
    }

    .bracket-wrapper {
        height: 100%;
    }

    .bracket-wrapper:last-of-type {
        transform: rotate(180deg);
    }

    #teck-stack-container {
        grid-column: 2 / span 10;
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(218px, 1fr));
        row-gap: var(--vspacing-2);
        column-gap: var(--vspacing-2);
        list-style: none;
        margin: 0;
        padding: 0;
    }

    .teck-stack-item {
        font-family: var(--font-titles);
        color: var(--accent-7);
        font-size: var(--font-size-h3);
        text-transform: capitalize;
        line-height: 45.9px;
        font-weight: 400;
        text-align: center;
    }



</style>