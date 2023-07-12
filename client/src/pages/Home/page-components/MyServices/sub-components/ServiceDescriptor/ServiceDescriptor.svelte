<script>
    import ServicePointer from "./ServicePointer.svelte";

    export let service_description = "We conceptualize and design the blueprint of your software solution, establishing the foundational structure that ensures growth and scalability. This involves creating wireframes, defining user flows, and setting up the perfect software architecture."

</script>

<div id="service-descriptor-wrapper" style:position="relative">
    <ServicePointer/>
    <div id="service-description-wrapper">
        <p aria-label="service description" id="service-description" style:position="relative">
            {service_description}
        </p>
    </div>
</div>

<style>
    #service-descriptor-wrapper {
        display: flex;
        width: 100%;
        height: 280px;
        grid-column: 3 / span 8;
        align-items: center;
        justify-content: center;
    }

    #service-description {
        z-index: var(--z-index-1);
    }
</style>