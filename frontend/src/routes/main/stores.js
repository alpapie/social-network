import { writable } from 'svelte/store';

export const displayContacts = writable(false);
export const ContactsStore = writable([]);
export const LastMessage = writable({})
export const darkMode = writable(true)
export const OnlineStore = writable([])
export const OfflineStore = writable([])
