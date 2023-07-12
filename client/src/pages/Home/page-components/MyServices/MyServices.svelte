<script>
    import ServicesHeader from "./sub-components/services_header.svelte";
    import ServiceDescriptor from "./sub-components/ServiceDescriptor/ServiceDescriptor.svelte";
    import ServicesContainer from "./sub-components/ServicesContainer/ServicesContainer.svelte";
</script>

<article id="my-services-section" class="section-content-layout">
    <ServicesHeader />
    <ServiceDescriptor />
    <ServicesContainer />
</article>

<style>
    #my-services-section {
        row-gap: var(--vspacing-3);
    }
</style>