import { writable } from "svelte/store";

export const is_available = writable(true);
export const token = writable("");
export const messages = writable([]);
export const is_user_human = writable(false);