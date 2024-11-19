<script lang="ts">
  import { modals } from '../../lib/modals';
  import Modal from "./Modal.svelte";
  
  function onClose() {
    modals.pop(); // pop top
  }
</script>

<div class:active={$modals.length > 0} class="interface">
  {#each $modals as modal, index}
  <div class:top={index === $modals.length - 1} style:z-index={index + 5} class="interface active">
    <Modal on:close={onClose} {modal}/>
  </div>
  {/each}
</div>

<style lang="scss">
  .interface {
    position: absolute;
    width: 100%;
    height: 100%;
    top: 0px;
    left: 0px;
    visibility: hidden;
    .active {
      visibility: visible;
    }
    .top {
      background-color: rgba(0, 0, 0, .7);
    }
  }
</style>
