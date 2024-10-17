<script lang="ts">
  import { onDestroy } from "svelte";
  export let retryTime: number;

  let dots = 4;
  const intvl = setInterval(() => {
    dots = dots - 1 || 4;
  }, 350);
  onDestroy(() => {
    clearInterval(intvl);
  });
</script>
<div class="parent">
  <div class="container">
    <img src="/offline.svg" alt="Offline">
    <div class="text-content">
      <h2>Server appears to be offline{".".repeat(dots - 1)}</h2>
      <span>Retrying in {retryTime} seconds</span>
    </div>
  </div>
</div>

<style>
  .text-content {
    padding: 35px;
  }
  h2 {
    margin: 0px;
  }
  span {
    display: block;
    font-size: 14px;
    margin-left: 20px;
    margin-top: 3px;
  }

  .parent {
    position: absolute;
    width: 100%;
    height: 100%;
  }
  .container {
    width: 500px;
    height: 500px;

    position: absolute;
    top: 30%;
    left: 50%;
    transform: translate(-50%, -30%);
  }
  img {
    display: block;
    margin: auto;
    width: 175px;
  }
</style>
