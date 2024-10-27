import { writable } from 'svelte/store';

export interface ModalComponent {
  props?: Record<string, any>;
}

export interface ModalOptions {
  component: any;
  props?: Record<string, any>;
  onClose?: () => void;
  closeOnClickOutside?: boolean;
  closeOnEscape?: boolean;
}

function createModalStore() {
  const { subscribe, set, update } = writable<ModalOptions[]>([]);

  return {
    subscribe,
    push: (modal: ModalOptions) => {
      update(modals => [...modals, modal]);
    },
    pop: () => {
      update(modals => {
        const lastModal = modals[modals.length - 1];
        lastModal?.onClose?.();
        return modals.slice(0, -1);
      });
    },
    remove: (index: number) => {
      update(modals => {
        const newModals = [...modals];
        const removedModal = newModals.splice(index, 1)[0];
        removedModal?.onClose?.();
        return newModals;
      });
    },
    clear: () => {
      update(modals => {
        modals.forEach(modal => modal.onClose?.());
        return [];
      });
    }
  };
}

export const modals = createModalStore();

// Utility functions
export function openModal(options: ModalOptions) {
  modals.push({
    closeOnClickOutside: true,
    closeOnEscape: true,
    ...options
  });
}

export function closeModal() {
  modals.pop();
}

export function closeAllModals() {
  modals.clear();
}

// Helper function to close specific modal by index
export function closeModalAtIndex(index: number) {
  modals.remove(index);
}
