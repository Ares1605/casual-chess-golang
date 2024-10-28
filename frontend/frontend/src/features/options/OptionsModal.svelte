<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import { notifs, TypesType } from "../../lib/notifs";
  import { openModal } from "../../lib/modals";
  import OptionsModal from './OptionsModal.svelte';
  import Button from "../../common/Button.svelte";

  const dispatch = createEventDispatcher();
  const closeModal = () => dispatch('close');

  let inviteCode: HTMLSpanElement;
  const copyInviteCode = () => {
    navigator.clipboard.writeText(inviteCode.textContent);
    notifs.add(TypesType.Success, "Copied to clipboard!");
  }
  const open =() => {
    openModal({
      component: OptionsModal,
      props: {
        title: 'Hello World'
      },
      closeOnClickOutside: true,
      closeOnEscape: true,
      onClose: () => console.log('Modal closed')
    });

  }
</script>
<div class="parent">
  <span on:click={closeModal} class="close">✖</span>
  <div class="buttons">
    <button>Options</button>
    <button>Sign Out</button>
    <button>Exit</button>
  </div>
</div>
<style>
  .buttons {
    display: flex;
    flex-direction: column;
    gap: 20px;
    width: 50%;
    margin: auto;
    padding: 20px 0px;
  }
  button {
    box-sizing: border-box;
    font-size: 20px;
    padding: 10px;
    letter-spacing: 5px;
    background-color: transparent;
    border: 1px solid var(--color-neutral);
    border-radius: 5px;
    color: var(--color-neutral);
    cursor: pointer;

    &:hover {
      background-color: var(--color-primary-light);
    }
  }
  .parent {
    position: fixed;

    left: 50%;
    top: 60px;
    transform: translateX(-50%);
    width: 450px;
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


