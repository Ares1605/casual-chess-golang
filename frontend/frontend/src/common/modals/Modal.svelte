<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import { type ModalOptions } from '../../lib/modals';
  export let modal: ModalOptions;

  const dispatch = createEventDispatcher();

  const closeModal = () => dispatch('close')
</script>

<div class="container">
  <div class="modal">
    <div class="content">
      <svelte:component
        this={modal.component}
          {...(modal.props || {})}
          on:close={() => dispatch("close")}
      />
    </div>
    <span on:click={closeModal} class="close">✖</span>
  </div>
</div>
<style>
  .container {
    position: absolute;
    top: 0px;
    left: 0px;
    width: 100%;
    height: 100%;
  }
  .content {
    position: relative;
  }
  .modal {
    position: fixed;

    left: 50%;
    top: 60px;
    transform: translateX(-50%);
    width: 450px;
    min-height: 40px;
    overflow: auto;
    background-color: var(--color-primary);
    border: 3px solid var(--color-primary-darker);
    border-radius: 5px;
  }
  .close {
    position: absolute;
    color: var(--color-neutral);
    font-size: 30px;
    cursor: pointer;
    top: 2px;
    right: 7px;
  }
</style>
