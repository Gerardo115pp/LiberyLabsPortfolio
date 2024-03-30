<script>
    import { layout_properties } from "@stores/layout";
    const svg_viewbox = $layout_properties.IS_MOBILE ? { width: 100, height: 50 } : { width: 100, height: 86 };
    const line_length = 10;

</script>

<svg viewBox="0 0 {svg_viewbox.width} {svg_viewbox.height}" class="pcw-pi-corners" fill="none" preserveAspectRatio={ $layout_properties.IS_MOBILE  ? 'none' : "xMidYMax meet"}>
    <path d="M1 {1+line_length}l0 -{line_length}l{line_length} 0" />
    <path d="M{svg_viewbox.width - 11} 1L{svg_viewbox.width - 1} 1L{svg_viewbox.width - 1} {1+line_length}" />
    <path d="M{svg_viewbox.width - 1} {svg_viewbox.height - (1 + line_length)}L{svg_viewbox.width - 1} {svg_viewbox.height - 1}L{svg_viewbox.width - (1 + line_length)} {svg_viewbox.height - 1}"/>
    <path d="M1 {svg_viewbox.height - (1 + line_length)}L1 {svg_viewbox.height - 1}L{(1 + line_length)} {svg_viewbox.height - 1}"/>
</svg>

<style>
    svg.pcw-pi-corners {
        position: absolute;
        top: 0;
        left: 0;
        height: 100%;
        width: 100%;
        z-index: var(--z-index-b-1);
    }

    svg.pcw-pi-corners path {
        stroke: var(--grey-1);
        stroke-width: .8;
    }
</style>