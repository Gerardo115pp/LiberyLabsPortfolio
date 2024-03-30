import { writable } from "svelte/store";
import { isMobile } from "@libs/utils";
import { browser } from "$app/environment";

let root_styles = {};
export const show_contact_form = writable(false);
export const hero_visible = writable(true);

export const layout_properties = {
    IS_MOBILE: false,
    VIEWPORT_WIDTH: 1920,
    VIEWPORT_HEIGHT: 1080,
};

export const defineLayout = () => {
    if (!browser) return;
    console.log("Running in the browser, defining layout properties");

    root_styles = getComputedStyle(document.documentElement);

    layout_properties.VIEWPORT_WIDTH = window.innerWidth;
    layout_properties.VIEWPORT_HEIGHT = window.innerHeight;

    layout_properties.IS_MOBILE = isMobile();
    
    layout_properties.SPACING = {
        VSPACING_1: root_styles.getPropertyValue("--vspacing-1"),
        VSPACING_2: root_styles.getPropertyValue("--vspacing-2"),
        VSPACING_3: root_styles.getPropertyValue("--vspacing-3"),
        VSPACING_4: root_styles.getPropertyValue("--vspacing-4"),
        VSPACING_5: root_styles.getPropertyValue("--vspacing-5"),
        VSPACING_6: root_styles.getPropertyValue("--vspacing-6"),
        VSPACING_7: root_styles.getPropertyValue("--vspacing-7"),
        VSPACING_8: root_styles.getPropertyValue("--vspacing-8"),
        VSPACING_9: root_styles.getPropertyValue("--vspacing-9")
    }
}