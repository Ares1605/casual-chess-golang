<script lang="ts">
  import type { apiresps } from "../wailsjs/go/models"
  import { GetFriends } from "../wailsjs/go/main/App"
  import AddFriend from './AddFriend.svelte';
  import { user } from "./lib/user";

  export let width: string;
  export let height: string;

  let friends: apiresps.Friends["data"] = [];
  GetFriends($user, $user.encoded_jwt).then(result => {
    if (result.success) {
      friends = result.data;
    }
  })
  let addFriend = false;
</script>
{#if addFriend}
  <AddFriend on:close={() => addFriend = false} />
{/if}
<div style="width: {width}; height: {height}" class="parent">
  <div class="nav-bar">
    <h2>Friends ({friends.length})</h2>
    <button on:click={() => addFriend = true} class="add-friend"></button>
  </div>
  <div class="friend-container">
    {#each friends as friend}
      <div class="friend">
        <img class="profile-picture" src={friend.profile_url} alt="Profile" />
        <span class="name">{friend.google_name}</span>
      </div>
    {/each}
  </div>
</div>
<style>
  .parent {
    display: flex;
    flex-direction: column;
    align-items: stretch;
    box-sizing: border-box;
    border: 3px solid #42210d;
    padding: 10px 5px;
    border-radius: 5px;
  }
  .nav-bar {
    width: 100%;
    height: 25px;
    display: flex;
    justify-content: space-between;
  }
  .nav-bar h2 {
    display: inline-block;
    margin: 0px;
  }
  .nav-bar .add-friend {
    padding: 0;
    height: 100%;
    aspect-ratio: 1;
    background-color: transparent;
    border: none;
    cursor: pointer;
    background-image: url('/add-friend.svg');
    background-size: cover;
    background-repeat: no-repeat;
  }
  .friend-container {
    position: relative;
    left: 8px;
    width: 100%;
    height: auto;
    overflow-y: scroll;
  }
  .friend {
    display: flex;
    justify-content: left;
    height: 45px;
    align-items: center;
    gap: 5px;
  }
  .profile-picture {
    border-radius: 3px;
    max-height: 100%;
    aspect-ratio: 1;
  }
</style>
