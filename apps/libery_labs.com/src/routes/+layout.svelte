<script>
    import '@app/app.css';
    import Navbar from "@components/Navbar/Navbar.svelte";
    import { defineLayout, hasChangedLayout } from '@stores/layout';
    import { onMount } from 'svelte';


    onMount(() => {
        defineLayout();
    });

    /*=============================================
    =            Methods            =
    =============================================*/
    
        /**
         * Handles the layout resize event
         * @param {UIEvent} event
         */
         const handleLayoutResize = e => {
            if (!browser) return;

            if (hasChangedLayout()) {
                defineLayout();
            }
        }
                
    
    /*=====  End of Methods  ======*/
</script>

<svelte:window 
    on:resize|passive={handleLayoutResize}
/>
<div id="libery-website-wrapper">
    <Navbar />
    <slot></slot>
</div>

<style>
    #libery-website-wrapper {
        padding: var(--navbar-height) 0 0 0;
    }
</style>