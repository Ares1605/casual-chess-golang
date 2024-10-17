<script lang="ts">
  import type { googleuser } from "../wailsjs/go/models"
  import { type Writable } from 'svelte/store';

  export let user: Writable<googleuser.GoogleUser | null>;

  export let width: string;
  export let height: string;
</script>
<div style="width: {width}; height: {height}" class="parent">
</div>
<style>
  .parent {
    background-color: yellow;
    border: 3px solid #42210d;
    border-radius: 5px;
    box-sizing: border-box;
  }
  .see-all {
    background-color: lightgrey;
  }
</style>
