<script lang="ts">
  import type { ComponentType } from "svelte";
  import type { Writable } from "svelte/store";
  import Profile from "./Profile.svelte";
  import FriendsButton from "./FriendsButton.svelte";
  import Home from "./Home.svelte";
  import BackButton from "./BackButton.svelte";

  export let route: Writable<ComponentType>;
</script>

<div class="parent">
  <div>
    {#if $route !== Home}
      <BackButton height="75%" {route} />
    {/if}
    <FriendsButton {route} />
  </div>
  <Profile height="75%" />
  <h1>Casual Chess</h1>
</div>
<style>
  .parent {
    position: relative;
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 60px;
  }
  .parent > div {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
  }
  h1 {
    position: absolute;
    /* throw it into the middle of the parent */
    left: 50%;
    font-size: 25px;
    transform: translateX(-50%);
    margin: 0px;
  }
</style>
