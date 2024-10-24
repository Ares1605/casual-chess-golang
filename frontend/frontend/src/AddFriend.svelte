<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import { notifs, TypesType } from "./lib/notifs";
  import HorizLine from "./common/HorizLine.svelte";

  const dispatch = createEventDispatcher();
  const closeModal = () => dispatch('close');

  let inviteCode: HTMLSpanElement;
  const copyInviteCode = () => {
    navigator.clipboard.writeText(inviteCode.textContent);
    notifs.add(TypesType.Success, "Copied to clipboard!");
  }
</script>
<div class="overlay">
  <div class="parent">
    <div class="nav-bar">
      <h2>Add Friend</h2>
      <span on:click={closeModal} class="close">✖</span>
    </div>
    <HorizLine />
    <div class="send-cont">
      <img src="/export.svg" alt="export">
      <div>
        <span bind:this={inviteCode}>1247189517358957</span>
        <button on:click={copyInviteCode} class="copy">
          <img src="/copy.svg" alt="copy">
        </button>
      </div>
    </div>
    <div class="invite-cont">
      <img src="/import.svg" alt="import">
      <div>
        <input>
      </div>
    </div>
  </div>
</div>
<style>
  .invite-cont div, .send-cont div {
    width: 100%;
    height: 80%;
    border: 2px solid #42210d;
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
  .copy img {
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
    border: 3px solid #42210d;
    border-radius: 5px;
  }
  .invite-cont > img, .send-cont > img {
    height: 65%;
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
    background-color: #d4bea1;
    border: 3px solid #42210d;
    border-radius: 5px;
  }
  .nav-bar {
    position: relative;
  }
  h2 {
    margin: 0px;
    text-align: center;
  }
  .close {
    position: absolute;
    font-size: 30px;
    cursor: pointer;
    top: 50%;
    transform: translateY(-50%);
    right: 20px;
  }
  .overlay {
    z-index: 3;
    position: absolute;
    top: 0px;
    left: 0px;
    width: 100%;
    height: 100%;
    background-color: rgba(0,0,0,.2);
  }
</style>
