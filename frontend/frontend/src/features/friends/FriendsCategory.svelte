<script lang="ts">
  import UserItem from "./UserItem.svelte";

  export let users: any;
  export let minimizedWidth: string;
  export let name: string;
  
  let fold = false;
</script>

<div class="container">
  <button class="spacer" on:click={() => fold = !fold}>
    <span style="min-width: {minimizedWidth}; max-width: {minimizedWidth}" class="minimized">{users.length}</span>
    <span class="expanded">{name}</span>
    <button class:fold={fold} class="fold-btn">
      <svg class="arrow-icon" viewBox="90 0 76 256" xmlns="http://www.w3.org/2000/svg">
        <path d="M160,220a11.96287,11.96287,0,0,1-8.48535-3.51465l-80-80a12.00062,12.00062,0,0,1,0-16.9707l80-80a12.0001,12.0001,0,0,1,16.9707,16.9707L96.9707,128l71.51465,71.51465A12,12,0,0,1,160,220Z"/>
      </svg>
    </button>
  </button>
  <div class:fold={fold} class="users">
    {#each users as user}
      <UserItem {minimizedWidth} {user} />
    {/each}
  </div>
</div>

<style>
  .container {
    width: 100%;
  }
  .users {
    height: auto;
    width: 100%;
  }
  .users.fold {
    height: 0px;
    overflow-y: hidden;
  }
  .fold-btn {
    background-color: transparent;
    border: none;
    padding: 0px;
    position: absolute;
    top: 50%;
    right: 25px;
    height: 20px;
    cursor: pointer;
  }
  .fold-btn.fold .arrow-icon {
    transform: translateY(-50%) rotate(-90deg);
  }
  .arrow-icon {
    transition: transform .2s ease-out;
    transform: translateY(-50%) rotate(90deg);
    fill: var(--color-neutral);
    height: 100%;
    overflow: visible; /* otherwise part of it clips?? */
  }
  .minimized {
    text-align: center;
    padding: 5px;
    box-sizing: border-box;
  }
  .expanded {
    width: 100%;
    text-align: left;
    margin-left: 10px;
    letter-spacing: 2px;
  }
  .spacer {
    position: relative;
    color: var(--color-neutral);
    font-weight: bold;
    display: flex;
    align-items: center;
    text-align: center;
    max-height: 35px;
    min-height: 35px;
    width: 100%;
    background-color: transparent;
    border: none;
  }
</style>
