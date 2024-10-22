<script lang="ts">
  import { onMount } from "svelte";
  import { GetFriends } from "wailsjs/go/main/App";
  import { user } from "./lib/user";
  import { notifs, TypesType } from "./lib/notifs";
  import FC from "./FC.svelte";
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

  let container: HTMLButtonElement;
  onMount(() => {
    container.addEventListener("mouseleave", () => {
      expanded = false;
    });
  });
  
</script>
<button bind:this={container} on:click={() => expanded = true} class:expanded={expanded} class="container">
  <UserItem {minimizedWidth} user={$user} />
  <FC users={friends} {minimizedWidth} />
</button>

<style>
  .container:not {
    margin-right: -360px;
  }
  .container {
    display: flex;
    flex-direction: column;
    border: none;
    border-top: 2px solid #694f07;

    min-width: 430px;
    max-width: 430px;
    gap: 2px;
    height: 100%;
    box-sizing: border-box;
    background-color: #694f07;
  }
</style>
