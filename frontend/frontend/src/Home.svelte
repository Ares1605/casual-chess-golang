<script lang="ts">
  import { user } from "./lib/user";
  import OngoingGames from "./OngoingGames.svelte";
  import FriendsList from "./FriendsList.svelte";
</script>
<h1>Welcome, {$user.email}!</h1>
<p>You are now signed in.</p>
<OngoingGames width="250px" height="100%" />

