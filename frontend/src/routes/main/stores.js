import { writable } from 'svelte/store';

export const displayContacts = writable(false);
export const contactsStore = writable([]);
export const LastMessage = writable({})
export const darkMode = writable(true)