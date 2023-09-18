<script>
    import { getAllProjects, ProjectListItem, Project, getProjectByID } from "@models/Project";
    import ProjectsCarousel from "./sub-components/ProjectsCarousel.svelte";
    import ProjectPreviewer from "./sub-components/ProjectPreviewer.svelte";
    import LiberyHeadline from "@components/UI/LiberyHeadline.svelte";
    import { onMount } from "svelte";

    /**
     * @type {ProjectListItem[]}
    */
    let projects = [];

    /**
     * @type {Object.<string, Project>}
    */
    let full_projects = {}

    /**
     * @type {Project}
    */
    let selected_project;

    onMount(() => {
        requestProjects();
    });
    
    const handleSelectedProject = async (e) => {
        /**
         * @type {ProjectListItem}
        */
        let new_selected_project = e.detail.project;
        
        if (!(new_selected_project.ID in full_projects)) {
            try {
                const new_project = await requestSelectedProject(new_selected_project.ID);
                full_projects[new_selected_project.ID] = new_project;
            } catch (error) {
                console.error(`Http error code when getting project with ID ${new_selected_project.ID}: ${error}`);
                return;
            }
        }
        
        selected_project = full_projects[new_selected_project.ID];

    }

    const requestSelectedProject = (project_id) => {
        return new Promise((resolve, reject) => {
            const on_success = project => resolve(project);
            const on_error = error => reject(error);

            getProjectByID(project_id, on_success, on_error);
        });
    }

    const requestProjects = () => {
        const on_success = projects_list => {
            let sorted_projects = projects_list.sort((a, b) => b.relevance - a.relevance)
            projects = sorted_projects;
        };
        const on_error = error => console.error(`Http error code when getting portfolio projects: ${error}`);

        getAllProjects(on_success, on_error);
    }

</script>

{#if projects.length > 0}     
    <section id="my-work-section" class="section-content-layout" style:display="flex">
        <div class="section-headline-wrapper" style:grid-column="1 / -1">
            <LiberyHeadline animation_delay={300} animation_duration={800} headline_text="Recent work" headline_color="var(--main)" text_transform="capitalize"/>
        </div>
        <div id="projects-carousel-component" class="full-vw">
            <ProjectsCarousel on:selectedProject={handleSelectedProject}  {projects} />
        </div> 
        {#if selected_project !== undefined}
            <div id="project-preview-component" class="full-vw">
                <ProjectPreviewer project={selected_project} />
            </div>
        {/if}
    </section>
{/if}

<style>
    section#my-work-section {
        flex-direction: column;
        row-gap: var(--vspacing-5);
    }


    #projects-carousel-component {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }
</style>