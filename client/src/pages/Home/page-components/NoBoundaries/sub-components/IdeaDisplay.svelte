<script>
    import LiberyBorderStar from "@components/UI/LiberyBorderStar.svelte";
    import { onMount } from "svelte";
    import { fade } from "svelte/transition";

    export let ideas = [];
    export let flip = false;
    let idea = "gatos pedorros";
    let idea_index = 0;

    let show_idea = false;
    let show_animation_duration = 800;
    let show_animation_delay = 500;
    let animation_active = true;
    const fuck = () => {idea= ideas[0]; show_idea = !show_idea; return idea};
    window.fuck = fuck;
    
    onMount(() => {
        showIdea();
    });

    const showIdea = () => {
        if (ideas.length === 0 || !animation_active) return;

        if(!show_idea) {
            idea_index = idea_index === ideas.length - 1 ? 0 : idea_index + 1;

            idea = ideas[idea_index];
            show_idea = true;

            setTimeout(() => {
                showIdea();
            }, 3.5 * (show_animation_duration + show_animation_delay));
        } else {
            show_idea = false;  
            setTimeout(() => {
                showIdea();
            }, show_animation_duration + show_animation_delay);
        }
    }

</script>

<div in:fade|local={{duration: 400, delay: 500}} class:debug={false} class:idea-wrapper-reverse={flip} class="bt-idea-display-wrapper">
    <div class="bt-idw-content">
        <div class="idea-wrapper">
            <h4>{idea}</h4>
        </div>
        <div 
            style:animation-duration={`${show_animation_duration}ms`}
            style:animation-delay={`${show_animation_delay}ms`}
            class="curtain-wrapper {show_idea ? 'unfold' : 'fold'}"
        >
            <div class="bt-idw-star-wrapper">
                <LiberyBorderStar
                stroke_color="var(--grey-3)"
                fill_color="var(--grey-9)"
                />
            </div>
            <div class="idea-hide-curtain"></div>
        </div>
    </div>
</div>

<style>

    @keyframes slide-out-right {
        0% {
            transform: translateX(0%);
        }
        100% {
            transform: translateX(105%);
        }
    }

    @keyframes slide-in-right {
        0% {
            transform: translateX(105%);
        }
        100% {
            transform: translateX(0%);
        }
    }

    @keyframes slide-out-left {
        0% {
            transform: translateX(0%);
        }
        100% {
            transform: translateX(-105%);
        }
    }

    @keyframes slide-in-left {
        0% {
            transform: translateX(-105%);
        }
        100% {
            transform: translateX(0%);
        }
    }



    .bt-idea-display-wrapper {
        --star-wrapper: calc(var(--vspacing-3) * 1.6);
        --curtain-width: calc(100% - var(--star-wrapper));
        --position-offset: 10%;

        display: flex;
    }
    
    .bt-idea-display-wrapper .bt-idw-content {
        position: relative;
        width: max-content;
        display: flex;
        align-items: center;
        gap: var(--vspacing-1);
    }

    .bt-idea-display-wrapper.idea-wrapper-reverse {
        flex-direction: row-reverse;
    }

    .curtain-wrapper {
        position: absolute;
        width: calc(100% + var(--position-offset));
        display: flex;
        top: calc(var(--star-wrapper) * -0.25);
        left: calc(var(--position-offset) * -1);
        transform-origin: left;


        /* /* animation: slide-in 1.2s ease-in-out; */
        animation-timing-function: ease-in-out;
        animation-fill-mode: both; 
        animation-name: slide-out-right;
        
    }
    
    .curtain-wrapper.fold {
        animation-name: slide-in-right;
    }

    .curtain-wrapper.unfold {
        animation-name: slide-out-right;
    }



    .idea-wrapper-reverse .curtain-wrapper {
        transform-origin: right;
        flex-direction: row-reverse;
        left: auto;
        right: calc(var(--position-offset) * -1);
    }

    .idea-wrapper-reverse .curtain-wrapper.fold {
        animation-name: slide-in-left;
    }

    .idea-wrapper-reverse .curtain-wrapper.unfold {
        animation-name: slide-out-left;
    }

    .idea-hide-curtain {
        width: var(--curtain-width);
        height: var(--star-wrapper);

        background: var(--body-bg-color);
    }

    .idea-wrapper-reverse .idea-hide-curtain {
        transform-origin: left;
    }

    .idea-wrapper h4 {
        font-family: var(--font-decorative);
        /* text-transform: lowercase; */
        color: var(--grey-8);
        font-size: var(--font-size-4);
        font-weight: 400;
        margin: 0;
        width: max-content;
        min-width: max-content;
        white-space: nowrap;
    }

    .bt-idw-star-wrapper {
        width: var(--star-wrapper);
        background-color: var(--body-bg-color);
    }
</style>