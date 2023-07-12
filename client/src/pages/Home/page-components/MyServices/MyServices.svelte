<script>
    import ServicesHeader from "./sub-components/services_header.svelte";
    import ServiceDescriptor from "./sub-components/ServiceDescriptor/ServiceDescriptor.svelte";
</script>

<article id="my-services-section" class="section-content-layout">
    <ServicesHeader />
    <ServiceDescriptor />
</article>

<style>
    #my-services-section {
        row-gap: var(--vspacing-3);
    }
</style>