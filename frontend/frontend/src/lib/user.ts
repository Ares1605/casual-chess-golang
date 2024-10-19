import type { user as userModel } from "../../wailsjs/go/models";
import { writable, type Writable } from 'svelte/store';

export const user: Writable<userModel.User| null> = writable(null);
