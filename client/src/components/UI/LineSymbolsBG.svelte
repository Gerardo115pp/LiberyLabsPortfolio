<script>
    import line_symbols from "@svg/line_symbols.svg";
</script>

<div class="lines-wrapper">
    {@html line_symbols}
</div>

<style>
    .lines-wrapper {
        --liners-symbols-height: 81.76%;

        position: absolute;
        width: 94.14%;
        height: var(--liners-symbols-height);
        top: calc(52% - var(--liners-symbols-height) / 2);
        left: 0;
        z-index: -1;
    }
</style>