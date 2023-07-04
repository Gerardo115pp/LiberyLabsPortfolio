<script>
    import TaggedText from "@components/Wrappers/TaggedText.svelte";

    export let headline_tag = "h1";
    export let headline_text;
    // @type {string} - extra props to be added to the headline tag as html
    export let extra_props;
    export let headline_color;
    export let vspacing = "var(--vspacing-1)";
    
    
    /*----------  Style  ----------*/

    export let text_transform = "uppercase";
    
    
    
    const headline_settings = {
        h1: {
            font_size: "var(--font-size-h1)",
            color: "var(--main)",
        },
        h2: {
            font_size: "var(--font-size-h2)",
            color: "var(--main-5)",
        },
    }

    $: headline_color = headline_color || headline_settings[headline_tag]?.color;
    const headline_font_size = headline_settings[headline_tag]?.font_size || "var(--font-size-h1)";

</script>

<TaggedText {vspacing} tag_name={headline_tag} {extra_props}>
    <div class="headline-wrapper">
        <h1 class="libery-headline" 
            style:color="{headline_color}" 
            style:font-size={headline_font_size} 
            style:text-transform={text_transform}
        >
            {headline_text}
        </h1>
        {#if headline_tag === "h1"}
            <div class="bottom-lines">
                <svg viewBox="0 0 275 23" fill="none" preserveAspectRatio="xMidYMax">
                    <path d="M1 16.5088L124.385 2" stroke="#BC3D00" stroke-width="3"/>
                    <path d="M95.8305 16.5088L274.28 20.5088" stroke="#BC3D00" stroke-width="3"/>
                </svg>                
            </div>
        {/if}
    </div>
</TaggedText>

<style>
    .headline-wrapper {
        display: flex;
        width: 100%;
        flex-direction: column;
        row-gap: var(--vspacing-2);
    }

    h1.libery-headline {
        font-size: 112px;
        font-weight: 400;
        letter-spacing: -3.36px;
        line-height: .75;
        margin: 0;
        padding: 0;
    }

    .bottom-lines {
        max-width: var(--vspacing-6);
        padding-left: var(--vspacing-4);
    }
</style>