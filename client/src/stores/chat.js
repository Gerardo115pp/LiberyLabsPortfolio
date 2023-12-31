import { writable } from "svelte/store";

export const is_available = writable(true);
export const token = writable("");
export const messages = writable([]);
export const is_user_human = writable(false);
export const max_message_length = writable(0);
export const is_big_mode_enabled = writable(false);