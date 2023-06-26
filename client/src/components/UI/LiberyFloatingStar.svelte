<script>
    import star from "@svg/libery_star_full.svg"
    import blue_star from "@svg/libery_star_blue.svg"
    import libery_line from "@svg/libery_line.svg"

    const star_x_offset = 74; // px
    const star_y_offset = 37; // px
</script>

<div id="libery-star-wrapper">
    <div class="svg-container" id="lsw-star" style:top="{star_y_offset}px" style:left="{star_x_offset}px">
        {@html star}
    </div>
    <div class="svg-container" id="lsw-blue-bg" style:top="{star_y_offset - 37}px" style:left="{star_x_offset - 32}px">
        {@html blue_star}
    </div>
    <div id="lsw-lines">
        <div id="lsw-line-one" class="line-wrapper svg-container">
            {@html libery_line}
        </div>
        <div id="lsw-line-two" class="line-wrapper svg-container">
            {@html libery_line}
        </div>
        <div id="lsw-line-three" class="line-wrapper svg-container">
            {@html libery_line}
        </div>
    </div>
</div>

<style>
    #libery-star-wrapper {
        --line-offset: 13%;
        --limit-lines-rotation: 5deg;
        --star-width: 736.585px;

        position: relative;
        width: 100%;
        height: 100%;
        z-index: 10;
    }

    :global(#libery-star-wrapper svg) {
        width: 100%;
    }

    
    /*----------  star  ----------*/
    
    

    #lsw-star {
        position: absolute;
        width: var(--star-width);
        top: 0;
        left: 0;
        z-index: 3;
    }

    :global(#lsw-star polygon.libery-star-full-element) {
        stroke: var(--grey) !important;
    }


    /*----------  blue star  ----------*/


    #lsw-blue-bg {
        position: absolute;
        width: var(--star-width);
        z-index: 1;
    }

    
    /*----------  Lines  ----------*/
    
    
    #lsw-lines {
        position: absolute;
        width: 100%;
        height: 100%;
        z-index: 2;
    }

    .line-wrapper {
        width: calc(100% - var(--line-offset));
    }

    #lsw-line-one {
        position: absolute;
        top: 101px;
        left: 0;
        transform: translateX(calc(var(--line-offset) * 1.2)) rotate(calc(var(--limit-lines-rotation) * -1));
    }

    #lsw-line-two {
        position: absolute;
        top: 50%;
        left: 0;
    }

    #lsw-line-three {
        position: absolute;
        bottom: 101px;
        left: 0;
        transform: translateX(calc(var(--line-offset) * 1.2)) rotate(calc(var(--limit-lines-rotation) * 1));
    }

    #lsw-line-one, #lsw-line-three {
        
        /* transform-origin: top left; */
        
    }

</style>