<script lang="ts">
  import { onDestroy } from "svelte";
  import Offline from "./Offline.svelte";
  import Pending from "./Pending.svelte";
  import { ServerOnline } from "../wailsjs/go/main/App.js";
  let pending = true;
  let serverOnline = false;
  let currentServerTimeout: number|undefined;
  
  let retryTimes = [5, 3];
  let extraRetryTime = 10;
  let retryTime = 0;
  const retryWithDelay = () => {
    retryTime = retryTimes.pop() || extraRetryTime;
    const countdown = () => {
      if (retryTime === 0) {
        ServerOnline()
          .then(result => {
            serverOnline = result;
            if (!result)
              retryWithDelay();
          })
      } else {
        currentServerTimeout = setTimeout(() => {
          retryTime--;
          countdown();
        }, 1000);
      }
    }
    countdown();
  }
  const updateServerStatus = () => {
    ServerOnline()
      .then(result => {
        pending = false;
        serverOnline = result;
        if (!result)
          retryWithDelay();
      })
  }
  updateServerStatus();

  onDestroy(() => currentServerTimeout && clearTimeout(currentServerTimeout));
</script>
{#if pending}
  <Pending />
{:else if serverOnline}
  <slot />
{:else}
  <Offline {retryTime} />
{/if}
