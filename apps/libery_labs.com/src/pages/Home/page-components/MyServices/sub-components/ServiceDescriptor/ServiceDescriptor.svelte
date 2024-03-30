<script>
    import ServicePointer from "./ServicePointer.svelte";
    import { SECTIONS as HOME_SECTIONS } from "@pages/Home/sections";

    export let service_description;
    export let point_direction;

</script>

<div data-scroll-section={HOME_SECTIONS.SERVICES} id="service-descriptor-wrapper" style:position="relative">
    <ServicePointer {point_direction} />
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