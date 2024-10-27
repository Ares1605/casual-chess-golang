<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import type { ModalOptions } from '../../lib/modals';
  import { onDestroy } from 'svelte';
  
  export let modal: ModalOptions;
  export let zIndex: number;
  
  const dispatch = createEventDispatcher();
  
  function handleClickOutside(event: MouseEvent) {
    if (
      modal.closeOnClickOutside && 
      event.target === event.currentTarget
    ) {
      dispatch('close');
    }
  }
  
  function handleKeydown(event: KeyboardEvent) {
    if (modal.closeOnEscape && event.key === 'Escape') {
      dispatch('close');
    }
  }

  // Cleanup function for keydown listener
  function cleanup() {
    window.removeEventListener('keydown', handleKeydown);
  }

  // Add keydown listener when component mounts
  window.addEventListener('keydown', handleKeydown);
  
  // Cleanup on component destroy
  onDestroy(cleanup);
</script>

<div
  class="modal-wrapper"
  style:z-index={zIndex}
  on:click={handleClickOutside}
  role="dialog"
  aria-modal="true"
>
</div>

<style>
  .modal-wrapper {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    /* Ensure click events pass through from stacked modals */
    pointer-events: auto;
  }
  
  .modal-content {
    background: white;
    border-radius: 4px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
    max-width: 90%;
    max-height: 90%;
    overflow: auto;
    /* Prevent click events from passing through the modal content */
    pointer-events: auto;
    /* Add some padding by default */
    padding: 1rem;
    /* Prevent text selection during modal dragging */
    user-select: none;
  }

  /* Ensure modal content scrolls properly on mobile */
  @media (max-width: 768px) {
    .modal-content {
      max-height: 85vh;
      width: 90vw;
      margin: auto;
    }
  }
</style>
