<script>
    import { onMount } from "svelte";
    import { ProjectListItem } from "@models/Project";
    import { createEventDispatcher } from "svelte";
    import CarouselItemCorners from "./CarouselItemCorners.svelte";

    onMount(() => {
        if (projects.length > 0) {
            selectProject(0);
        }
    })

    /**
     * @type {ProjectListItem[]}
     */
    export let projects = [];
    let selected_project;
    const selected_project_dispatcher = createEventDispatcher();

    const disableSelectionMark = () => {
        const selected_project_element = document.querySelector('.selected-project');

        if (selected_project_element !== null) {
            selected_project_element.classList.remove('selected-project');
        }
    }

    const emitSelectedProject = project => {
        selected_project_dispatcher('selectedProject', { project });
    }

    const markSelectedProject = index => {
        disableSelectionMark();

        const project_element = document.getElementById(`pcw-pc-project-${index}`);

        if (project_element !== null) {
            project_element.classList.add('selected-project');
        }
    }

    const selectProject = project_index => {
        if (selected_project === projects[project_index]) return;

        selected_project = projects[project_index];

        emitSelectedProject(selected_project);
        markSelectedProject(project_index);
    }
</script>

<div id="projects-carousel-wrapper">
    <div 
        id="pcw-projects-container" 
        class:debug={false} 
        style:justify-content={projects.length > 7 ? 'flex-start' : 'center'}
        style:grid-template-columns="{projects.length}fr"
    >
        {#each projects as project, h}
            <div tabindex="0" role="button" style:grid-column="{h+1} / span 1" id="pcw-pc-project-{h}" on:click={() => selectProject(h)} class="pcw-project-item" style:position="relative">
                <span class="pcw-pi-name">{project.name}</span>
                <CarouselItemCorners />
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
        container-type: size;
        container-name: project-carousel;
        width: 100%;
        height: calc(var(--vspacing-5) * 1.6844);
        border-top: 3px dashed var(--danger-9);
        border-bottom: 3px dashed var(--danger-9);
    }

    #pcw-projects-container {
        --project-item-width: calc(var(--vspacing-6) * 0.9414);

        display: grid;
        height: 100%;
        grid-template-rows: 1fr;
        grid-template-columns: auto;
        grid-auto-columns: var(--project-item-width);
        width: max-content;
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

    .pcw-pi-name {
        cursor: default;
        transition: all .6s ease-in-out;
    }

    @media(pointer: fine) {
        .pcw-project-item:not(.selected-project):hover .pcw-pi-name{
            transform: scale(1.1);
        }
    }

    :global(.selected-project) {
        color: var(--main-dark);
        font-size: var(--font-size-4);
    }

    
    /*=============================================
    =            Mobile            =
    =============================================*/
    
    @media (max-width: 768px) {
        #projects-carousel-wrapper {
            height: calc(var(--vspacing-4) * 2.2844);
            justify-content: flex-start;
            align-items: flex-start;
            overflow-y: auto;
        }

        #pcw-projects-container {
            --project-item-width: 40cqw;
        }

        .pcw-project-item {
            padding: var(--vspacing-2) var(--vspacing-2);
        }
    }
    
    /*=====  End of Mobile  ======*/
    
    
</style>