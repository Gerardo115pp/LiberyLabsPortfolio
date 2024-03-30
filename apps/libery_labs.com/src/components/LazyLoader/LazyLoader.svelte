<script>
    import viewport from './actions/useViewPortAction';
    export let id;
    export let className;
    export let image_url;

    let image_element;
    let image_src = "";

    let isLoaded = false;

    const handleViewportEnter = e => {
        image_element = document.createElement('img');

        image_element.onload = () => {
            image_src = image_element.src;
            isLoaded = true;
        };

        image_element.src = image_url;
    };
</script>

<div use:viewport on:viewportEnter={handleViewportEnter}  class={`lazy-wrapper ${className !== undefined ? className : ''}`} {id}>
    {#if isLoaded}
        <slot name="lazy-wrapper-image" {image_src}></slot>
    {/if}
</div>

<style>
    
    @keyframes placeholderShimmer {
        0% {
            background-position: 220% 0
        }
        100% {
            background-position:  -210% 0
        }
    }
    
    .lazy-wrapper:empty {
        animation-duration: 1s;
        animation-fill-mode: forwards;
        animation-iteration-count: infinite;
        animation-name: placeholderShimmer;
        animation-timing-function: linear;
        background: rgb(236,236,236);
        background: linear-gradient(98deg, rgba(236,236,236,1) 34%, rgba(255,255,255,1) 38%, rgba(255,255,255,1) 50%, rgba(235,235,235,1) 71%); 
        background-size: 130%;
    }
</style>