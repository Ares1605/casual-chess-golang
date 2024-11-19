<script lang="ts">
  import { onMount } from "svelte";
  import { userProfile } from "../../lib/userProfile";
    import UserProfile from "./UserProfile.svelte";

  export let user: any;
  export let minimizedWidth: string;
  const src = user.profile_url;

  let container: HTMLDivElement;

  let show = false;

  onMount(() => {
    container.addEventListener("mouseenter", () => {
      console.log({ user: user, reference: container });
      userProfile.set({ user: user, reference: container});
      show = true;
    });
    container.addEventListener("mouseleave", () => {
      userProfile.set(null);
      show = false;
    });
  });
</script>

<div bind:this={container} class="container">
  <div class="item">
    <img style="width: {minimizedWidth}"{src} alt="profile">
    <span>{user.username}</span>
  </div>
  {#if show}
    <UserProfile {user} />
  {/if}
</div>

<style>
  .container {
    position: relative;
    width: 100%;
  }
  .item {
    transition: .2s background-color;
    display: flex;
    flex-direction: row;
    justify-content: start;
    align-items: center;
    background-color: var(--color-primary-dark);
    color: var(--color-neutral);
    font-weight: bold;
    font-size: 15px;
    width: 100%;
  }
  .item:hover {
    background-color: var(--color-primary);
  }
  img {
    box-sizing: border-box;
    padding: 5px;
    opacity: .9;
  }
</style>
