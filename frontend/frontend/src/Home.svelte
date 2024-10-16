<script lang="ts">
  import type { googleuser } from "../wailsjs/go/models"
  import type { Writable } from 'svelte/store';
  import OngoingGames from "./OngoingGames.svelte";
  import FriendsList from "./FriendsList.svelte";

  export let user: Writable<googleuser.GoogleUser | null>;
</script>

{#if $user}
  <h1>Welcome, {$user.email}!</h1>
  <p>You are now signed in.</p>
  <OngoingGames {user} width="250px" height="100%" />
  <FriendsList {user}  width="250px" height="200px" />
{:else}
  error
{/if}
