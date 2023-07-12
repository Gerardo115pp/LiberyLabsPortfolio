<script>
    export let spacing = 24;
    export let color="var(--grey-5)";
</script>

<svg id="dot-grid-element" viewBox="0 0 1000 500">
    <defs>
        <pattern id="dot-grid" x="0" y="0" width="{spacing}" height="{spacing}" patternUnits="userSpaceOnUse">
            <circle cx="{spacing/3}" cy="{spacing/2}" r="1" fill="{color}" />
        </pattern>
    </defs>
    <rect x="0" y="0" width="100%" height="100%" fill="url(#dot-grid)" />
</svg>