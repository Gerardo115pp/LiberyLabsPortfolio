import { writable } from "svelte/store";

export const recaptcha_pk = RECAPTCHA_PK;
export let has_recaptcha_loaded = writable(false);