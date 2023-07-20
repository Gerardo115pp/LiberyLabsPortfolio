<script>
    import { onMount } from "svelte";

    onMount(() => {
        const a = document.querySelector('.test_b')
        console.log(a)
    })

    export let projects = [
        "test_a0",
        "test_a",
        "test_a",
        "test_a",
        "test_a",
        "test_a",
        "test_a",
        "test_a",
        "test_b"
    ]
</script>

<div id="projects-carousel-wrapper">
    <div id="pcw-projects-container" class:debug={false} style:justify-content={projects.length > 7 ? 'flex-start' : 'center'}>
        {#each projects as project}
            <div class="pcw-project-item" style:position="relative">
                <span class="pcw-pi-name">{project}</span>
                <svg viewBox="0 0 100 86" class="pcw-pi-corners" fill="none" preserveAspectRatio="xMidYMax meet">
                    <path d="M1 10V0H10" />
                    <path d="M89.2 0h10v10" />
                    <path d="M1 75.3v10h10" />
                    <path d="M89.2 85.3h10v-10" /> 
                </svg>
            </div>
        {/each}
    </div>
</div>

<style>
    #projects-carousel-wrapper {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        width: 100%;
        height: calc(var(--vspacing-5) * 1.6844);
        border-top: 3px dashed var(--danger-9);
        border-bottom: 3px dashed var(--danger-9);
    }

    #pcw-projects-container {
        --project-item-width: calc(var(--vspacing-6) * 0.9414);

        display: flex;
        height: 100%;
        width: var(--page-content-width);
        overflow: hidden;
        flex-wrap: nowrap;
        scroll-behavior: smooth;
        gap: 0px;
    }
    
    .pcw-project-item {
        font-family: var(--font-decorative);
        display: grid;
        place-items: center;
        height: 100%;
        width: var(--project-item-width);
        flex-shrink: 0;
    }

    svg.pcw-pi-corners {
        position: absolute;
        top: 0;
        left: 0;
        height: 100%;
        width: 100%;
        z-index: var(--z-index-b-1);
    }

    svg.pcw-pi-corners path {
        stroke: var(--grey-1);
        stroke-width: .8;
    }
</style>