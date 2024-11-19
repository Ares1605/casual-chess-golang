import { writable } from "svelte/store";

type userProfileType = {
  user: any,
  reference: HTMLElement
} | null;

const { subscribe, set } = writable<userProfileType>(null);

export const userProfile = {
  subscribe,
  set,
  clear: () => set(null)
};
