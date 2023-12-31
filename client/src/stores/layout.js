import { writable } from "svelte/store";
import { isMobile } from "@libs/utils";

const root_styles = getComputedStyle(document.documentElement);
export const show_contact_form = writable(false);
export const hero_visible = writable(true);

export const layout_properties = {
    IS_MOBILE: isMobile(),
    VIEWPORT_WIDTH: document.documentElement.clientWidth,
    VIEWPORT_HEIGHT: document.documentElement.clientHeight,
};

export const defineLayout = () => {
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