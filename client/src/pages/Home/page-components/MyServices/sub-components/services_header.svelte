<script>
    import LiberyBorderStar from "@components/UI/LiberyBorderStar.svelte";
    import LiberyHeadline from "@components/UI/LiberyHeadline.svelte";
    import { layout_properties } from "@stores/layout";

</script>

<header class="services-header">
    <div id="white-star-wrapper">
        <LiberyBorderStar />
    </div>
    <LiberyHeadline animated={false} forced_font_size="{layout_properties.IS_MOBILE ? 'calc(var(--font-size-h2) * 1.03)' : 'var(--font-size-h1)'}" headline_text="what can i craft for you?" headline_color="var(--grey-1)" text_transform="lowercase" />
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

    
    /*=============================================
    =            mobile            =
    =============================================*/
    
        @media only screen and (max-width: 765px) {
            .services-header {
                row-gap: var(--vspacing-3);
            }

            #white-star-wrapper {
                width: 30%;
            }
        }
    
    /*=====  End of mobile  ======*/
    
    
</style>