<script lang="ts">
  import type { ComponentType } from 'svelte';
  import { type Writable, writable } from "svelte/store";
  import { AuthStatuses } from "../lib/types";
    
  import Header from "../layout/Header.svelte";
  import Home from "./Home.svelte";
  import FriendsList from "../features/friends/FriendsList.svelte";
  
  export let authStatus: Writable<AuthStatuses>;

  const route: Writable<ComponentType> = writable(Home);
  const modals: Writable<ComponentType[]> = writable([]);
</script>

<div class="parent">
  <div style="height: 60px"><Header {route} /></div>
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
    overflow: scroll;
  }
  main {
    display: flex;
    flex-grow: 1; /* grow to the maximum */
    overflow: auto;
  }
  main > div {
    width: 100%;
  }
</style>
