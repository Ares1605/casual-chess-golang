import { writable, derived, type Readable } from 'svelte/store';
import type { apiresps } from "../../wailsjs/go/models";

export enum TypesType {
  Error = "Error",
  Informational = "Informational",
  Warning = "Warning",
  Success = "Success"
}

type NotifType = {
  title?: string,
  body: string,
  type: TypesType
};

function createNotifStore() {
  const { subscribe, update } = writable<NotifType[]>([]);
  const add = (type: TypesType, body: string, title?: string) => {
    let notif: NotifType = {
      body: body,
      title: title,
      type: type
    };
    update(notifs => [...notifs, notif]);
    setTimeout(() => {
      update(notifs => notifs.filter(temp => temp != notif))
    }, 5000);
  }
  return {
    subscribe,
    addEndpointError: (result: apiresps.JunkResp) => add(TypesType.Error, result.error?.message, result.error?.type || undefined),
    add: add
  };
}

export const notifs = createNotifStore();

export const readNotifs: Readable<NotifType[]> = derived(notifs, $notifs => $notifs);
