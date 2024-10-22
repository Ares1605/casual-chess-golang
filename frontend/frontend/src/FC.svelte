<script lang="ts">
  import UserItem from "./UserItem.svelte";

  export let users;
  export let minimizedWidth: string;
  console.log(users);
  
</script>

<div class="spacer">
  <div style="width: {minimizedWidth}" class="minimized">{users.length}</div>

</div>
{#each users as user}
  <UserItem {minimizedWidth} {user} />
{/each}

<style>
  .minimized {
    text-align: center;
    font-weight: bold;
    padding: 5px;
  }
  .spacer {
    display: flex;
    align-items: center;
    text-align: center;
    max-height: 35px;
    min-height: 35px;
    width: 100%;
  }
</style>
