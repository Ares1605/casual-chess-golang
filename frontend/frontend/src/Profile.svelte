<script lang="ts">
   import { user } from "./lib/user";
   const src = $user.profile_url;
   export let height = "100%";
</script>

<div class="parent" style="height: {height}">
  <img {src} class="logo" alt="logo">
  <span class="username">@{$user.username}</span>
</div>
<style>
  .parent {
    display: flex;
    flex-direction: row;
    justify-content: start;
    align-items: center;
    gap: 4px;
  }
  .logo {
    height: 100%;
    border-radius: 5px;
  }
  .username {
    font-weight: bold;
    font-size: 16px;
  }
</style>
