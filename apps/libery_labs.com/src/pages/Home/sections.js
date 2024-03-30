import { writable } from "svelte/store"


export const SECTIONS = {
    HERO: "hero-section-content",
    ABOUT: "about-section-content",
    ARSENAL: "arsenal-section-content",
    SERVICES: "services-section-content",
    PORTFOLIO: "portfolio-section-content",
    CONTACT: "contact-section-content",
    CHAT: "chat-section-content",
    FOOTER: "footer-section-content"
}



export const scrollToSection = section_id => {
    if (section_id === "") return;

    const target = document.querySelector(`[data-scroll-section="${section_id}"]`)

    if (target === null) return;

    target.scrollIntoView({
        behavior: "smooth",
        block: "center",
        inline: "end"
    });

}
