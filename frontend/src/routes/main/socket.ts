import { writable } from 'svelte/store';

let ws // = new WebSocket("ws://localhost:8080/server/ws")
console.log("sennd the request" , ws)
let wss
let user
export const WS = writable(ws)
export const WSS = writable(wss)
