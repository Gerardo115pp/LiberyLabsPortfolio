<script>
    import { show_contact_form, layout_properties, hero_visible } from "@stores/layout";
    import { writable } from "svelte/store";
    import MainLogo from "@components/UI/MainLogo.svelte";
    import BurgerBtn from "@components/UI/BurgerBTN.svelte";
    import { fade } from "svelte/transition";
    import DropMenu from "./sub-components/DropMenu.svelte";
    import { SECTIONS as HOME_SECTIONS } from "@pages/Home/sections";

    const dropdown_names = [
        {
            name: "Chat",
            href: HOME_SECTIONS.CHAT,
            options: []
        },
        {
            name: "Top",
            href: HOME_SECTIONS.HERO,
            options: []
        },
        {
            name: "Work",
            href: HOME_SECTIONS.PORTFOLIO,
            options: []
        },
        {
            name: "About me",
            href: HOME_SECTIONS.ABOUT,
            options: []
        },
        {
            name: "Contact",
            href: HOME_SECTIONS.CONTACT,
            options: []
        }
    ]

    let menu_visible = writable(false);
    
    
</script>

<nav class:opaque-background={$menu_visible} id="libery-labs-navbar" class:adebug={false}>
    <div id="lln-content">
        <div id="lln-right-content">
            <div class="lln-rc-burger-wrapper">
                <BurgerBtn on:click={() => menu_visible.set(!$menu_visible)} is_opened={menu_visible} />
            </div>
            {#if !$hero_visible}
                <div in:fade={{duration: 300}} id="lln-main-logo-wrapper">
                    <MainLogo 
                        fill_color="var(--main-dark)" 
                        stroke_color="var(--main-dark)" 
                        subheadline_color="var(--grey-1)"
                        subheadline_font_size="calc(var(--font-size-h3) * 0.35)"
                    />
                </div>
            {/if}
        </div>
        <div id="lln-cta">
            <button on:click={() => show_contact_form.set(true)} id="cta">
                <strong class="highlight-text--CTA">Get in touch</strong>
            </button>
        </div>
    </div>
    {#if $menu_visible}
        <DropMenu sections={dropdown_names} visible={menu_visible} />
    {/if}
</nav>

<style>
    /* * {
        border: 1px solid red;
    } */

    #libery-labs-navbar {
        --navoptions-gap: 81px;

        position: fixed;
        top: 0;
        display: grid;
        height: var(--navbar-height);
        background: linear-gradient(90deg, var(--grey) 0%, var(--grey) 9%, hsla(0, 0%, 5%, .1) 29%, var(--grey) 80%);
        width: 100%;
        place-items: center;
        z-index: var(--z-index-t-2);
        backdrop-filter: blur(10px);
        border-bottom: .5px solid var(--grey-9);
    }
    
    #libery-labs-navbar.opaque-background {
        background: var(--grey) !important;
        border-bottom: 1px solid var(--grey-1);
    }

    #lln-content {
        display: flex;
        width: var(--page-content-width);
        justify-content: space-between;
    }

    #lln-right-content {
        display: flex;
        align-items: center;
        gap: var(--navoptions-gap);
    }

    .lln-rc-burger-wrapper {
        --burger-btn-width: calc(1.1 * var(--vspacing-3));

        box-sizing: border-box;
        width: var(--burger-btn-width);
        height: calc(0.8666666666666667 * var(--burger-btn-width));
    }

    #lln-main-logo-wrapper {
        width: calc(var(--vspacing-7) * 0.5);
    }

    button#cta {
        background: none;
        border: none;
        padding: none;
        color: var(--success-3);
        transform-origin: center;
        transition: transform 0.2s ease-in-out;
    }

    @media(pointer: fine) {
        button#cta:hover {
            transform: scale(1.05);
        }
    }
</style>