<script lang="ts">
  import type { ComponentType } from 'svelte';
  import { type Writable, writable } from "svelte/store";
  import { AuthStatuses } from "../lib/types";
    
  import Header from "../layout/Header.svelte";
  import Home from "./Home.svelte";
  import FriendsList from "../features/friends/FriendsList.svelte";
  
  export let authStatus: Writable<AuthStatuses>;

  const route: Writable<ComponentType> = writable(Home);
</script>

<div class="parent">
  <div><Header {route} /></div>
  <main>
    <div>
      <svelte:component this={$route} {route} />
    </div>
    <FriendsList />
  </main>
</div>

<style>
  .parent {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
  }
  main {
    display: flex;
    height: 100%;
  }
  main > div {
    width: 100%;
  }
</style>
