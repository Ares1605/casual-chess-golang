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
    color: var(--color-button-text);
    border: var(--button-border-width) solid var(--color-button-border);

    transition: var(--button-transition);

    &:hover {
      background-color: var(--color-button-hover);
      color: var(--color-button-text-hover);
      border: var(--button-border-width-hover) solid var(--color-button-border-hover);
    }
    &:active {
      background-color: var(--color-button-active);
      color: var(--color-button-text-active);
      border: var(--button-border-width-active) solid var(--color-button-border-active);
      transform: translate(-50%, -50%) var(--button-transform-active);
    }
  }
  .bubble {
    border-radius: 20px;
  }
</style>
