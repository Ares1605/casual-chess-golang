import { writable, derived, type Readable } from 'svelte/store';

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

  return {
    subscribe,
    add: (notif: NotifType) => {
      update(notifs => [...notifs, notif]);
      setTimeout(() => {
        update(notifs => notifs.filter(temp => temp != notif))
      }, 5000);
    }
  };
}

export const notifs = createNotifStore();

export const readNotifs: Readable<NotifType[]> = derived(notifs, $notifs => $notifs);
