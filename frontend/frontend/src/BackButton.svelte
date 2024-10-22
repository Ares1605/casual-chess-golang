<script lang="ts">
  import type { ComponentType } from "svelte";
  import type { Writable } from "svelte/store";
  import Home from "./Home.svelte";
  export let route: Writable<ComponentType>

  export let height = "100%";
</script>
<svg on:click={() => $route = Home} style="height: {height}" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
  <path d="M4 4V20M8 12H20M8 12L12 8M8 12L12 16" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
</svg>

<style>
  svg {
    cursor: pointer;
  }
  path {
    stroke: #42210d;
  }
</style>
