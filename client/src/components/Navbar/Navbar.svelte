<script>
    import { isMobile } from "@libs/utils";

    const dropdown_names = {
        HOME: {
            name: "Home",
            href: "/#libery-home-hero",
            options: []
        },
        WORK: {
            name: "Work",
            href: "/#libery-work",
            options: []
        },
        ABOUT: {
            name: "About me",
            href: "/#libery-about",
            options: []
        },
        CONTACT: {
            name: "Contact",
            href: "/#libery-contact",
            options: []
        }
    }
    
    const is_mobile = isMobile();
</script>

<nav id="libery-labs-navbar">
    <div id="lln-content">
        <div id="lln-navoptions-container">
            {#each Object.values(dropdown_names) as ddn}
                <div class="lln-navoption">
                    <a class="highlight-text" href="{ddn.href}">{ddn.name}</a>
                    {#if ddn.options.length > 0}
                        <div class="lln-dropdown">
                            {#each ddn.options as option}
                                <a href="{option.href}">{option.name}</a>
                            {/each}
                        </div>
                    {/if}
                </div>
            {/each}
        </div>
        <div id="lln-cta">
            <button id="cta">
                <strong class="highlight-text--CTA">Get in touch</strong>
            </button>
        </div>
    </div>
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
        background: var(--grey);
        width: 100%;
        place-items: center;
        z-index: 1000;
    }

    #lln-content {
        display: flex;
        width: var(--page-content-width);
        justify-content: flex-end;
    }

    #lln-navoptions-container {
        display: flex;
        gap: var(--navoptions-gap);
        margin-right: var(--navoptions-gap);
    }

    #lln-navoptions-container .lln-navoption {
        display: flex;
        align-items: center;
    }

    #lln-navoptions-container .lln-navoption a {
        text-decoration: none;
        color: var(--main-dark);
        text-transform: uppercase;
    }   

    button#cta {
        background: none;
        border: none;
        padding: none;
        color: var(--success-3);
    }
</style>