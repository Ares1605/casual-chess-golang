<script lang="ts">
  import { modals } from '../../lib/modals';
  import Modal from "./Modal.svelte";
  
  function onClose() {
    modals.pop(); // pop top
  }
</script>

<div class:active={$modals.length > 0} class="interface">
  {#each $modals as modal, index}
  <div style:z-index={index + 5} class="modal">
    <Modal on:close={onClose} {modal}/>
  </div>
  {/each}
</div>

<style>
  .interface {
    z-index: 4;
    transition: 0.2s opacity;
    position: absolute;
    visibility: hidden;
    opacity: 0;
    width: 100%;
    height: 100%;
    top: 0px;
    left: 0px;
    background-color: rgba(0, 0, 0, .4);
  }
  
  .interface.active {
    opacity: 1;
    visibility: visible;
  }
</style>
