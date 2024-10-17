<script lang="ts">
  import { readNotifs } from "./lib/notifs";
  import { fade } from "svelte/transition";

</script>
<div class="tray">
  {#each $readNotifs as notif}
    <div class="notif show {notif.type}" in:fade out:fade>
      <img src="./public/error-ico.svg" alt={notif.type}>
      <div>
        <h2>{notif.title || notif.type}</h2>
        <span>{notif.body}</span>
      </div>
      <span class="close"></span>
    </div>
  {/each}
</div>

<style>
  .tray {
    z-index: 4;
    visibility: hidden;
    position: fixed;
    right: 15px;
    bottom: 15px;
    width: 425px;
    display: flex;
    justify-content: flex-end;
    flex-direction: column;
    gap: 7px;
  }
  .notif {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;

    position: relative;
    z-index: 4;
    overflow: hidden;
    visibility: visible;
    box-sizing: border-box;
    width: 100%;
    font-size: 14px;
    min-height: 50px;
    padding: 14px;
    opacity: 0;
    margin-bottom: -70px;
    white-space: break-spaces; /* so JS carriage returns work */
    color: black;
    cursor: pointer;
    border-left-width: 5px;
    border-left-style: solid;
    box-shadow: rgba(0, 0, 0, 0.24) 0px 3px 8px;
    transition: all .45s;
  }
  .show {
    opacity: 1;
    margin-bottom: 0px;
  }
  .notif div {
    display: flex;
    justify-content: space-between;
    flex-direction: column;
    gap: 5px;
    margin: 0px 30px 0px 20px;
  }
  img {
    margin-top: 3px;
    width: 20px;
  }
  h2 {
    font-size: 15px;
    margin: 0 3px 0 0;
  }
  .close {
    position: absolute;
    top: 7px;
    right: 10px;
    font-size: 15px;
    &::before {
      content: "✖";
    }
  }
  .Success {
    background-color: #E7F4E7;
    border-left-color: #0C7D07;
  }
  .Error {
    background-color: #FBE8E9;
    border-left-color: #DA1415;
  }
  .Warning {
    background-color: #FCF9E7;
    border-left-color:#F0C100;
  }
  .Informational {
    background-color: #EBF0FD;
    border-left-color: #2D60E5;
  }

</style>
