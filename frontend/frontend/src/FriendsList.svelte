<script lang="ts">
  import { type Writable } from 'svelte/store';
  import type { apiresps, googleuser } from "../wailsjs/go/models"
  import { GetFriends } from "../wailsjs/go/main/App"

  export let user: Writable<googleuser.GoogleUser | null>;
  export let width: string;
  export let height: string;

  let friends: apiresps.Friends["data"] = [];
  GetFriends($user).then(result => {
    if (result.success) {
      friends = result.data;
    }
  })
</script>
<div style="width: {width}; height: {height}" class="parent">
  <h2>Friends ({friends.length})</h2>
  {#each friends as friend}
  <div class="friend">
    <img class="profile-picture" src={friend.profile_url} alt="Profile" />
    <span class="name">{friend.display_name}</span>
  </div>
  {/each}
</div>
<style>
  .parent {
    box-sizing: border-box;
    padding: 10px;
    background-color: yellow;
    border: 3px solid #42210d;
    border-radius: 5px;
  }
  .friend {
    background-color: grey;
    display: flex;
    justify-content: left;
    height: 65px;
    align-items: center;
    gap: 5px;
  }
  .profile-picture {
    border-radius: 3px;
    max-height: 100%;
    aspect-ratio: 1;
  }
</style>
