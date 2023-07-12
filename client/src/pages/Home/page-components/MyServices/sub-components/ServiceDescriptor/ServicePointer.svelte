<script>
    import center_pointer from "./pointer-svgs/center-pointer.svg";
</script>

<div id="pointer-wrapper">
    {@html center_pointer}
</div>

<style>
    #pointer-wrapper {
        position: absolute;
        height: inherit;
        width: max-content;
        left: 50%;
        transform: translateX(-50%);
        z-index: var(--z-index-b-2);
    }

    :global(#pointer-wrapper > svg) {
        
        height: inherit;
        width: max-content;
    }

    :global(#pointer-wrapper > svg > path) {
        stroke: var(--main-dark-color-9);
    }
</style>