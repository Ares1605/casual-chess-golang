<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import { notifs, TypesType } from "../../lib/notifs";
  import HorizLine from "../../common/HorizLine.svelte";
  import { openModal } from "../../lib/modals";
  import AddFriendModal from './AddFriendModal.svelte';

  const dispatch = createEventDispatcher();
  const closeModal = () => dispatch('close');

  let inviteCode: HTMLSpanElement;
  const copyInviteCode = () => {
    navigator.clipboard.writeText(inviteCode.textContent);
    notifs.add(TypesType.Success, "Copied to clipboard!");
  }
  const open =() => {
    openModal({
      component: AddFriendModal,
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
  <div class="nav-bar">
    <h2>Add Friend</h2>
    <span on:click={closeModal} class="close">✖</span>
  </div>
  <HorizLine />
  <div class="send-cont">
    <svg viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path d="M8.71,7.71,11,5.41V15a1,1,0,0,0,2,0V5.41l2.29,2.3a1,1,0,0,0,1.42,0,1,1,0,0,0,0-1.42l-4-4a1,1,0,0,0-.33-.21,1,1,0,0,0-.76,0,1,1,0,0,0-.33.21l-4,4A1,1,0,1,0,8.71,7.71ZM21,14a1,1,0,0,0-1,1v4a1,1,0,0,1-1,1H5a1,1,0,0,1-1-1V15a1,1,0,0,0-2,0v4a3,3,0,0,0,3,3H19a3,3,0,0,0,3-3V15A1,1,0,0,0,21,14Z"/></svg>
    <div>
      <span bind:this={inviteCode}>1247189517358957</span>
      <button on:click={open} class="copy">
        <svg width="800px" height="800px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path fill-rule="evenodd" clip-rule="evenodd" d="M21 8C21 6.34315 19.6569 5 18 5H10C8.34315 5 7 6.34315 7 8V20C7 21.6569 8.34315 23 10 23H18C19.6569 23 21 21.6569 21 20V8ZM19 8C19 7.44772 18.5523 7 18 7H10C9.44772 7 9 7.44772 9 8V20C9 20.5523 9.44772 21 10 21H18C18.5523 21 19 20.5523 19 20V8Z" fill="#42210d"/>
          <path d="M6 3H16C16.5523 3 17 2.55228 17 2C17 1.44772 16.5523 1 16 1H6C4.34315 1 3 2.34315 3 4V18C3 18.5523 3.44772 19 4 19C4.55228 19 5 18.5523 5 18V4C5 3.44772 5.44772 3 6 3Z" fill="#42210d"/>
        </svg>
      </button>
    </div>
  </div>
  <div class="invite-cont">
    <svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path d="M4 13V19C4 20.1046 4.89543 21 6 21H18C19.1046 21 20 20.1046 20 19V13" stroke="#42210d" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/>
      <path d="M12 3L12 15M12 15L8.5 11.5M12 15L15.5 11.5">
    </svg>
    <div>
      <input>
    </div>
  </div>
</div>
<style>
  .invite-cont div, .send-cont div {
    width: 100%;
    height: 80%;
    border: 2px solid var(--color-primary-darker);
    border-radius: 3px;
    box-sizing: border-box;
  }
  .send-cont div {
    position: relative;
    display: flex;
    align-items: center;
    padding: 10px;
  }
  .copy {
    background-color: transparent;
    border: none;
    cursor: pointer;
    height: 80%;
    overflow: hidden;
    position: absolute;
    padding: 0px;
    top: 50%;
    transform: translateY(-50%);
    right: 15px;
  }
  .copy:hover {
    height: 90%;
  }
  .copy svg {
    height: 100%;
  }
  .invite-cont input {
    width: 100%;
    height: 100%;
    box-sizing: border-box;
    padding-left: 10px;
    background-color: transparent;
    border: none;
    outline: none;
  }
  .invite-cont, .send-cont {
    display: flex;
    gap: 10px;
    align-items: center;
    margin: auto;
    height: 50px;
    width: 85%;
    padding: 5px 10px;
    box-sizing: border-box;
    border: 3px solid var(--color-primary-darker);
    border-radius: 5px;
  }
  .invite-cont > svg, .send-cont > svg {
    height: 65%;
  }
  .send-cont > svg {
    fill: var(--color-accent);
  }
  .invite-cont > svg, .invite-cont path {
    stroke: var(--color-accent);
  }
  .parent {
    position: fixed;
    display: flex;
    flex-direction: column;
    padding: 10px 0px;
    gap: 10px;

    left: 50%;
    top: 60px;
    transform: translateX(-50%);
    width: 450px;
    overflow: auto;
    background-color: var(--color-primary);
    border: 3px solid var(--color-primary-darker);
    border-radius: 5px;
  }
  .nav-bar {
    position: relative;
  }
  h2 {
    margin: 0px;
    text-align: center;
    color: var(--color-neutral);
  }
  .close {
    position: absolute;
    color: var(--color-neutral);
    font-size: 30px;
    cursor: pointer;
    top: 50%;
    transform: translateY(-50%);
    right: 20px;
  }
</style>

