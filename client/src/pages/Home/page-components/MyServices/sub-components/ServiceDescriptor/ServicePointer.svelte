<script>

    /**
     * @typedef {Object} PathData
     * @property {string} d
     * @property {string} view_box
     */

    /**
     * @type {Object.<string, PathData>}
     */
    const path_values = {
        CENTER: {
            d: "M250 1L375 50L250 400L125 50L185.7 25Z", // use exactly 1 move command, 4 line commands and 1 close command or 5 line commands and no close command.
            view_box: "0 0 500 400"
        },
        LEFT: {
            d:  "M1.99996 298.883L208 2.32293L373.5 28.7254L445 63.0484L1.99996 298.883Z",
            view_box: "0 0 500 400"
        },
        RIGHT: {
            d: "M392.5 342.908L4 25.8521L89 2.40796L244.5 41.99L392.5 342.908Z",
            view_box: "0 0 500 400"
        }
    }

    export let point_direction = -1; // -1: left, 0: center, 1: right
    let old_point_direction = point_direction;

    /** @type {SVGPathElement} */
    let svg_pointer;
    /** @type {SVGAnimateElement} */
    let svg_pointer_animate;

    let direction_shape_data = path_values.LEFT;

    $: direction_shape_data = updateDirection(point_direction);

    /**
     * Sets the values for the animate element within the svg_pointer path element
     * @param f {PathData} - Old direction pointer path data
     * @param t {PathData} - New direction pointer path data
     */    
    const animatePointer = (f,t) => {
       if (svg_pointer === undefined) return;

        if (svg_pointer_animate === undefined) {
            svg_pointer_animate = document.createElementNS("http://www.w3.org/2000/svg", "animate");
            svg_pointer_animate.setAttribute("attributeName", "d");
            svg_pointer_animate.setAttribute("dur", ".5s");
            svg_pointer_animate.setAttribute("fill", "freeze");

        }
        
        svg_pointer_animate.setAttribute("from", f.d);
        svg_pointer_animate.setAttribute("to", t.d);

        svg_pointer.appendChild(svg_pointer_animate);
        svg_pointer_animate.beginElement();
    }

    const getDirectionData = direction => {
        let new_direction_shape_data = path_values.LEFT;
        if (direction === 0) {
            new_direction_shape_data = path_values.CENTER;
        } else if (direction === -1) {
            new_direction_shape_data = path_values.RIGHT;
        }

        return new_direction_shape_data;
    }

    function updateDirection(direction) {
        if (direction === old_point_direction) return direction_shape_data;
        
        let old_direction_shape_data = getDirectionData(old_point_direction);
        let new_direction_shape_data = getDirectionData(direction);
        
        animatePointer(old_direction_shape_data, new_direction_shape_data);
        
        old_point_direction = direction;

        return new_direction_shape_data;
    }
</script>

<div id="pointer-wrapper">
    <svg viewBox="{path_values.LEFT.view_box}" fill="none">
        <path bind:this={svg_pointer} stroke="var(--grey-8)" stroke-width="3" d="{path_values.LEFT.d}">
            <!-- <animate bind:this={svg_pointer_animate}  attributeName="d" dur=".5s" fill="freeze" from="{path_values.LEFT.d}" to="{path_values.RIGHT.d}"></animate> -->
        </path>
    </svg>
</div>

<style>
    @keyframes morphSVG {
        0% {
            d: "M216.226 2.6615L463.226 67.6615L1.22605 314.662L216.226 2.6615Z";
        }
        100% {
            d: "M3.60449 167.941L192.766 2.8541L396.244 352.331L3.60449 167.941Z";
        }
    }

    #pointer-wrapper {
        position: absolute;
        height: inherit;
        width: max-content;
        left: 50%;
        transform: translateX(-50%);
        z-index: var(--z-index-b-2);
    }

    :global(#pointer-wrapper > svg) {
        
        height: inherit;
        width: max-content;
    }

    :global(#pointer-wrapper > svg > path) {
        stroke: var(--main-dark-color-9);
    }
</style>