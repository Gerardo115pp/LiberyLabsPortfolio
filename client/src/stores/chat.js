import { writable } from "svelte/store";

export const is_available = writable(true);
export const token = writable("");
export const messages = writable([]);