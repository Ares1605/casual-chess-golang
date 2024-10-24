<script lang="ts">
  import { onMount, onDestroy } from "svelte";
  import { GetFriends } from "../../../wailsjs/go/main/App";
  import { user } from "../../lib/user";
  import { notifs, TypesType } from "../../lib/notifs";
  import FriendsCategory from "./FriendsCategory.svelte";
  import UserItem from "./UserItem.svelte";

  let friends = [];
  GetFriends($user)
    .then(result => {
      if (!result.success)
        return notifs.addEndpointError(result);
      friends = result.data;
    })
    .catch(error => {
      notifs.add(TypesType.Error, String(error), "Fetching friends");
    });

  const minimizedWidth = "60px";

  let expanded = false;

  let listContainer: HTMLButtonElement;
  let closeListTimeout: number;
  onMount(() => {
    listContainer.addEventListener("mouseenter", () => {
      console.log("entered");
      clearTimeout(closeListTimeout);
    });
    listContainer.addEventListener("mouseleave", () => {
      console.log("left");
      clearTimeout(closeListTimeout);
      closeListTimeout = setTimeout(() => expanded = false, 200);
    });
  });
  onDestroy(() => {
    clearTimeout(closeListTimeout);
  });
  
</script>
<div class="container" class:expanded={expanded}>
  <button bind:this={listContainer} on:click={() => expanded = true} class="list-cont">
    <UserItem {minimizedWidth} user={$user} />
    <FriendsCategory users={friends} {minimizedWidth} />
  </button>
  <svg class="arrow-icon" viewBox="90 0 76 256" xmlns="http://www.w3.org/2000/svg">
    <path d="M160,220a11.96287,11.96287,0,0,1-8.48535-3.51465l-80-80a12.00062,12.00062,0,0,1,0-16.9707l80-80a12.0001,12.0001,0,0,1,16.9707,16.9707L96.9707,128l71.51465,71.51465A12,12,0,0,1,160,220Z"/>
  </svg>
</div>

<style>
  .container {
    display: flex;
    justify-content: right;
    flex-direction: row-reverse;
    align-items: center;
    height: 100%;
  }
  .arrow-icon {
    opacity: 0;
    transition: .2s opacity;
    width: 9px;
    padding: 5px;
    fill: #42210d;
  }
  .container:not(.expanded) .list-cont:hover + .arrow-icon {
    opacity: 1;
  }
  .container:not(.expanded) .list-cont:hover {
    opacity: .93;
  }
  .container:not(.expanded) .list-cont {
    margin-right: -190px;
  }
  .list-cont {
    transition: .2s opacity, .15s margin-right ease-in-out;
    display: flex;
    flex-direction: column;
    border: none;
    border-top: 5px solid #694f07;
    padding: 0px 3px 0px 0px;

    min-width: 250px;
    max-width: 250px;
    gap: 3px;
    height: 100%;
    box-sizing: border-box;
    background-color: #694f07;
    overflow-y: auto;
  }
  .list-cont::-webkit-scrollbar { 
    display: none;
  }
</style>

