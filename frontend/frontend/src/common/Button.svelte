<script lang="ts">
  export let bubble: boolean = false;
</script>
<button
  on:click
  on:mouseover
  on:focus
  on:mouseenter
  on:mouseleave
  class:bubble={bubble}><slot /></button>
<style lang="scss">
  button {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    font-size: 25px;
    width: 175px;
    height: 50px;
    border-radius: 10px;
    cursor: pointer;

    background-color: var(--color-button);
    border: 4px solid var(--color-button-border);
    color: var(--color-button-text);

    &:hover {
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
      background-color: var(--color-button-hover);
      border-color: var(--color-button-hover);
      transform: translateY(-50%, calc(-50% - 1px));
    }
    &:active {
      background-color: var(--color-button-hover);
      border-color: var(--color-button-hover);
      /* Remove hover effects */
      transform: translate(-50%, calc(-50% + 1px));
      box-shadow: none;
    }
  }
  .bubble {
    border-radius: 25px;
  }
</style>
