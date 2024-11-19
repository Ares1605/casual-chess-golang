<script lang="ts">
    import { openModal } from "../../lib/modals";
    import SettingsModal from "../settings/SettingsModal.svelte";

    const onClick = () => {
      openModal({
        component: SettingsModal,
        props: {},
        closeOnClickOutside: true,
        closeOnEscape: true
      });
    }
</script>
<div class="buttons">
  <button on:click={onClick}>Settings</button>
  <button>Sign Out</button>
  <button>Exit</button>
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
</style>


