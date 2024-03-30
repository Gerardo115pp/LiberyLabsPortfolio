<script>
    import '@app/app.css';
    import Navbar from "@components/Navbar/Navbar.svelte";
</script>

<div id="libery-website-wrapper">
    <Navbar />
    <slot></slot>
</div>

<style>
    #libery-website-wrapper {
        padding: var(--navbar-height) 0 0 0;
    }
</style>