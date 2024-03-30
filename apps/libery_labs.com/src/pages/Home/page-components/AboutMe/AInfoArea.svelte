<script>
    import AnchorsOne from "@components/UI/AnchorsOne.svelte";
    import LiberyHeadline from "@components/UI/LiberyHeadline.svelte";
    import TaggedText from "@components/Wrappers/TaggedText.svelte";
    import { fly } from "svelte/transition";
    import { bounceOut } from "svelte/easing";

    export let visible = false;
    export let animation_delay_increment = 300;
    export let animation_duration = 800;

</script>

{#key visible}
    <div style:visibility={visible ? "visible" : "hidden"} id="text-area" style:position="relative">
        <div id="text-area-anchor-wrapper" >
            <AnchorsOne animation_delay={(animation_duration*1.2)+animation_delay_increment}/>
        </div>
        <div in:fly={{x: 1500, duration: animation_duration*2, delay: animation_delay_increment, easing: bounceOut }} class="burger-section">
            <LiberyHeadline 
                vspacing="var(--vspacing-2)" 
                headline_tag="h2" 
                headline_text="Hi, I,m Lalo!" 
                extra_props='role=”<span>passionate-engineer</span>”'
                text_transform="capitalize"
            />
        </div>
        <div in:fly={{x: 1500, duration: animation_duration, delay: animation_delay_increment*2 }} class="burger-section">
            <TaggedText tag_name="p">
                <p class="bio-block">
                    Welcome! My name is Lalo, I'm a fullstack web engineer based in Jalisco, Mexico. I'm passionate about turning ideas into reality and taking businesses to the next level. I understand front-end, backend and I'm in love with Linux, Docker and K8S, i know exactly what how each of these technologies a must be combined to give <strong>your business the scalability, reliability and efficiency it needs to grow and succeed</strong>.
                </p>
            </TaggedText>
        </div>
        <div in:fly={{x: 1500, duration: animation_duration, delay: animation_delay_increment*3 }} class="burger-section">
            <TaggedText tag_name="p">
                <p class="bio-block">
                    Ever since i was a child i always loved building things, and getting new tools on my toolkit to achieve bigger and better things. Aside from coding, my other passions are carpentry, wood carving, 3D modeling, printing and painting miniatures. For me the only thing better than building things is helping others through something i built.
                </p>
            </TaggedText>
        </div>
    </div>
{/key}

<style>
    /*----------  Anchor  ----------*/
        
        #text-area-anchor-wrapper {
            position: absolute;
            top: -20%;
            left: -23%;
            height: 0px;
        }

    
    /*----------  End of Anchor  ----------*/
     
        #text-area {
            grid-column: 6 / span 7;
            display: flex;
            flex-direction: column;
            row-gap: var(--vspacing-3);
        }

        @media only screen and (max-width: 768px) {
            #text-area-anchor-wrapper {
                position: static;
                width: max-content;
                height: max-content
                /* transform: translateX(-55%); */
            }
        }

</style>