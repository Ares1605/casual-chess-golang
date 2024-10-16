import { writable, type Writable } from 'svelte/store'

export const addingFriend: Writable<boolean> = writable(false)
