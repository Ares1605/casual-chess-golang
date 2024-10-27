<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import { type ModalOptions } from '../../lib/modals';
  export let modal: ModalOptions;

  const dispatch = createEventDispatcher();

</script>

<div>
  <svelte:component
    this={modal.component}
      {...(modal.props || {})}
      on:close={() => dispatch("close")}
  />
</div>
<style>
  div {
    width: 100%;
    height: 100%;
    position: absolute;
    top: 0px;
    left: 0px;
  }
</style>
